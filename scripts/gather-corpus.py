#!/usr/bin/env python3
"""Gather the documented and real-world SPL/SPL2 example queries.

The checked-in corpus under conformance/corpus is what the tests read; this
script only refreshes it. It needs network access, so it is never run by tests
or CI.

    scripts/gather-corpus.py --cache ~/.cache/openspl/docs
    scripts/gather-corpus.py --versions 9.0,9.2,10.0
    scripts/gather-corpus.py --dialect spl2
    scripts/gather-corpus.py --skip-detections      # documentation only

Two kinds of collection are gathered.

Documented (`kind`: documented). The Splunk documentation renders its pages
client-side and exposes each page's body at a fragment URL. Command example
queries appear in `<textarea class="code-raw-content">` elements, which is
what this script reads: one query per element, no prose.

Derived (`kind`: derived). The Splunk security_content repository carries
detection searches in the same SPL an analyst writes. Those searches call
macros, so the repository's own macro definitions are expanded first, and the
pipeline is cut at the first stage that calls a macro the repository does not
define. The result is the detection's SPL up to that point, with nothing
invented; the collection is labelled derived because it is not the text as
published.

Each source collection gets a metadata.json that records the dialect and the
source version the queries were taken from, because a query file does not say
which Splunk release it belongs to. A query that appears in several versions is
stored once, with the versions listed on its line; a query that appears in only
some versions keeps that narrower list, so version-specific syntax stays
visible.

The expected outcome on each line (accepted, or the diagnostic code of a
refusal) is not written here. It is a snapshot of the current parser, recorded
by `go test ./conformance -run TestCorpus -update-corpus`, so a behaviour
change shows up as a reviewable corpus diff.

A fetch that fails is not treated as a page without queries: the run exits
without writing, because the corpus would otherwise lose whatever that page
held and the baseline diff would look like a documentation change. Pass
`--allow-partial` to write anyway.
"""

import argparse
import concurrent.futures
import html
import json
import os
import re
import shutil
import sys
import tarfile
import tempfile
import time
import urllib.error
import urllib.request

UA = ("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) "
      "AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0 Safari/537.36")

# SPL1 documentation: one search reference per Splunk Enterprise version.
#
# Only versions whose pages are public. A version number that redirects to the
# login page yields no pages, so it is not listed here; check `--versions`
# output before adding one.
SPL1_BASE = "https://help.splunk.com/en/splunk-enterprise/search/spl-search-reference"
SPL1_VERSIONS = ["9.0", "9.1", "9.2", "9.3", "9.4", "10.0"]
SPL1_GROUPS = [
    "search-commands",
    "quick-reference",
    "search-in-the-cli",
    "evaluation-functions",
    "statistical-and-charting-functions",
]

# SPL2 documentation is published for Splunk Cloud Platform at one URL; the
# navigation lists the groups, and each group page lists its own pages.
SPL2_BASE = "https://help.splunk.com/en/splunk-cloud-platform/search/spl2-search-reference"
SPL2_GROUPS = [
    "spl2-compatibility-profiles-and-quick-references",
    "quick-reference-for-spl2-commands",
    "evaluation-functions",
    "statistical-and-charting-functions",
    "dataset-functions",
]
SPL2_COMMAND_GROUPS = [
    "addinfo-command", "append-command", "appendcols-command",
    "appendpipe-command", "bin-command", "branch-command", "convert-command",
    "decrypt-command", "dedup-command", "eval-command", "eventstats-command",
    "expand-command", "fields-command", "fieldsummary-command",
    "fillnull-command", "flatten-command", "from-command", "head-command",
    "if-command", "into-command", "iplocation-command", "join-command",
    "loadjob-command", "lookup-command", "makemv-command",
    "makeresults-command", "mstats-command", "mvcombine-command",
    "mvexpand-command", "nomv-command", "ocsf-command", "rename-command",
    "replace-command", "reverse-command", "rex-command", "route-command",
    "search-command", "sort-command", "spath-command", "spl1-command",
    "stats-command", "streamstats-command", "table-command", "tags-command",
    "thru-command", "timechart-command", "timewrap-command", "tstats-command",
    "typer-command", "union-command", "untable-command", "where-command",
]

# Real-world SPL: the Splunk security_content detections, Apache-2.0. Pinned to
# a commit because the corpus is checked in and a moving branch would rewrite
# it on every gather.
SECURITY_CONTENT_URL = "https://github.com/splunk/security_content"
SECURITY_CONTENT_REF = "a3f6523eabf696d66e20fdd80766bfae95575577"
SECURITY_CONTENT_TRANSFORM = (
    "Each detection's search with the repository's own macro definitions "
    "expanded, cut at the first stage that calls a macro the repository does "
    "not define, and whitespace-normalized."
)

CODE_BLOCK = re.compile(
    r'<textarea class="code-raw-content"[^>]*>(.*?)</textarea>', re.S)
HREF = re.compile(r'(?:hx-get|href)="(/en/[^"]+)"')
TAG = re.compile(r"<[^>]+>")

# Placeholder syntax in a documentation template, not a query.
#
# Brackets are deliberately absent: `[` starts a subsearch and an array
# literal, so filtering every bracketed span would remove a whole class of real
# queries. A template that uses `[...]` is caught by the words inside it, which
# are the same metavariables and ellipses covered here.
PLACEHOLDER = re.compile(
    r"<[A-Za-z][^<>]*>|\{[^}]*\}|\.\.\.|…|\$[A-Za-z_]\w*\$")

# A stand-in word the documentation writes where a command or field name
# belongs, to show syntax rather than an example: `| commandname`. A block
# containing one documents a grammar, not a query.
METAVARIABLE = re.compile(r"\bcommandname\b")

# A block is a query rather than a syntax template when it looks like the start
# of a search or a pipeline stage. `search` must be the command, not the start
# of a parameter fragment such as `search="disabled=0"`.
SEARCH_LEAD = re.compile(
    r"^(?:from\s|index\s*=|source\s*=|sourcetype\s*=|host\s*=|eventtype\s*=|"
    r"tag\s*=|search(?:\s|$))", re.I)

# A documentation page often shows the other dialect beside its own: the SPL2
# reference compares SPL2 syntax with the classic SPL equivalent, and the
# function pages show both. Each collection keeps only queries that begin the
# way its dialect begins a pipeline, so an SPL1 example never lands in the SPL2
# collection or the reverse.
SPL1_START = re.compile(r"^(?:from|select)\b", re.I)
SPL2_START = re.compile(r"^(\|?\s*)([A-Za-z_]\w*)\b")

# A shell transcript, not a query: `./splunk search ...`, `myhost $ splunk ...`.
SHELL = re.compile(r"^\./|^\$|\s\$\s")

MAX_QUERY = 2000

# A macro call: `name` or `name(arguments)`.
MACRO_CALL = re.compile(r"`([A-Za-z0-9_.\-]+)(?:\(([^`]*)\))?`")


def fetch_bytes(url, cache_dir, retries=3):
    """Return the body as bytes, using the cache when it has one."""
    key = re.sub(r"[^A-Za-z0-9]+", "_", url)
    path = os.path.join(cache_dir, key) if cache_dir else None
    if path and os.path.exists(path):
        with open(path, "rb") as f:
            return f.read()
    last = None
    for attempt in range(retries):
        try:
            req = urllib.request.Request(url, headers={"User-Agent": UA})
            with urllib.request.urlopen(req, timeout=30) as r:
                body = r.read()
            if path:
                os.makedirs(cache_dir, exist_ok=True)
                with open(path, "wb") as f:
                    f.write(body)
            return body
        except (urllib.error.URLError, TimeoutError, OSError) as err:
            last = err
            time.sleep(1 + attempt)
    print("fetch failed: %s: %s" % (url, last), file=sys.stderr)
    return None


def fetch(url, cache_dir, retries=3):
    """Return the page body as text, using the cache when it has one."""
    body = fetch_bytes(url, cache_dir, retries)
    return None if body is None else body.decode("utf-8", "replace")


def fragment_url(base, path):
    """The fragment URL for a page path below the documentation base."""
    return "https://help.splunk.com/en/fragments/%s/%s" % (
        base.split("/en/", 1)[1], path)


def nav_url(base, group):
    """The fragment URL of a navigation group: one level below the base."""
    return "https://help.splunk.com/en/fragments/nav/%s/%s" % (
        base.split("/en/", 1)[1], group)


def page_paths(body, base):
    """Page paths named by a navigation fragment, relative to the base."""
    out = []
    prefix = "/en/" + base.split("/en/", 1)[1] + "/"
    for m in HREF.finditer(body):
        href = m.group(1)
        if "#" in href:
            href = href.split("#", 1)[0]
        if not href.startswith(prefix):
            continue
        rel = href[len(prefix):]
        if rel not in out:
            out.append(rel)
    return out


def code_blocks(body):
    """The example code blocks of a page, as plain text."""
    for m in CODE_BLOCK.finditer(body):
        text = TAG.sub("", m.group(1))
        text = html.unescape(text)
        text = re.sub(r"\s+", " ", text).strip()
        if text:
            yield text


def is_query(text, dialect):
    if len(text) > MAX_QUERY:
        return False
    if PLACEHOLDER.search(text) or SHELL.search(text):
        return False
    if METAVARIABLE.search(text):
        return False
    # A block that starts with a bracket is a configuration stanza such as
    # `[my-access-extractions] REGEX=...`, never a pipeline.
    if text.startswith("["):
        return False
    if dialect == "spl1":
        # SPL2's `from ...` and `select ...` pipelines are refused by the SPL1
        # front-end as a known command, so they add a refusal, not coverage.
        if SPL1_START.match(text):
            return False
        return "|" in text or bool(SEARCH_LEAD.match(text))
    # An SPL2 pipeline starts with `from` or `select`, or with a leading pipe
    # that continues one. Anything else on an SPL2 page is the SPL1 example
    # shown beside it.
    m = SPL2_START.match(text)
    return bool(m) and m.group(2).lower() in ("from", "select")


def gather_group(base, group, cache_dir, pool, dialect):
    """Map page path -> queries for one navigation group, and failed fetches.

    A page that could not be fetched is counted, not skipped silently: the
    caller refuses to write a corpus that is short because the network was.
    """
    nav = fetch(nav_url(base, group), cache_dir)
    if nav is None:
        return {}, 1
    pages = page_paths(nav, base)
    bodies = list(pool.map(lambda page: (page, fetch(fragment_url(base, page),
                                                     cache_dir)), pages))
    out = {}
    failed = 0
    for page, body in bodies:
        if body is None:
            failed += 1
            continue
        found = [q for q in code_blocks(body) if is_query(q, dialect)]
        if found:
            out[page] = found
    return out, failed


def yaml_scalar(value):
    """One YAML scalar as written: quoted, or plain up to a comment."""
    value = value.strip()
    if len(value) >= 2 and value[0] == "'" and value[-1] == "'":
        return value[1:-1].replace("''", "'")
    if len(value) >= 2 and value[0] == '"' and value[-1] == '"':
        return re.sub(r"\\(.)", lambda m: {"n": "\n", "t": "\t"}.get(
            m.group(1), m.group(1)), value[1:-1])
    if " #" in value:
        value = value.split(" #", 1)[0]
    return value


def yaml_block(lines, index, value):
    """The value of `key: <value>` at lines[index], and the next line index."""
    if value not in ("|", "|-", "|+", ">", ">-", ">+"):
        return yaml_scalar(value), index
    body = []
    indent = None
    i = index + 1
    while i < len(lines):
        line = lines[i]
        if line.strip():
            cur = len(line) - len(line.lstrip())
            if indent is None:
                indent = cur
            if cur < indent:
                break
            body.append(line[indent:])
        else:
            body.append("")
        i += 1
    if value.startswith(">"):
        return " ".join(x for x in body if x), i
    return "\n".join(body).strip(), i


def macro_definitions(root):
    """name -> (definition, arguments) for the repository's macros."""
    macros = {}
    for dirpath, _, names in os.walk(os.path.join(root, "macros")):
        for name in sorted(names):
            if not name.endswith(".yml"):
                continue
            lines = open(os.path.join(dirpath, name), encoding="utf-8",
                         errors="replace").read().split("\n")
            name_value = definition = None
            arguments = []
            in_arguments = False
            for i, line in enumerate(lines):
                if line.startswith("arguments:"):
                    in_arguments = True
                    continue
                if in_arguments and re.match(r"^\s+-\s+", line):
                    arguments.append(line.split("-", 1)[1].strip())
                    continue
                in_arguments = False
                for field in ("name", "definition"):
                    if line.startswith(field + ":"):
                        value, _ = yaml_block(
                            lines, i, line[len(field) + 1:].strip())
                        if field == "name":
                            name_value = value.strip()
                        else:
                            definition = value.strip()
                        break
            if name_value and definition is not None:
                macros[name_value] = (definition, arguments)
    return macros


def detection_search(path):
    """The `search` value of one detection file."""
    lines = open(path, encoding="utf-8", errors="replace").read().split("\n")
    for i, line in enumerate(lines):
        if line.startswith("search:"):
            value, _ = yaml_block(lines, i, line[len("search:"):].strip())
            return " ".join(value.split())
    return ""


def split_macro_args(text):
    """Split macro arguments on commas outside quotes and brackets."""
    out, cur, quote, depth = [], "", "", 0
    i, n = 0, len(text)
    while i < n:
        ch = text[i]
        if quote:
            cur += ch
            if ch == "\\" and i + 1 < n:
                i += 1
                cur += text[i]
            elif ch == quote:
                quote = ""
        elif ch in "\"'":
            quote = ch
            cur += ch
        elif ch == "(":
            depth += 1
            cur += ch
        elif ch == ")":
            depth -= 1
            cur += ch
        elif ch == "," and depth == 0:
            out.append(cur)
            cur = ""
        else:
            cur += ch
        i += 1
    if cur.strip():
        out.append(cur)
    return out


def expand_macros(query, macros, depth=0):
    """Expand every macro the repository defines, as far as it goes."""
    if depth > 8:
        return query

    def expand(match):
        name, raw = match.group(1), match.group(2)
        if name not in macros:
            return match.group(0)
        definition, arguments = macros[name]
        if not arguments:
            return definition
        if raw is None:
            return match.group(0)
        values = [x.strip() for x in split_macro_args(raw)]
        if len(values) != len(arguments):
            return match.group(0)
        out = definition
        for argument, value in zip(arguments, values):
            out = out.replace("$%s$" % argument, value)
        return out

    out = MACRO_CALL.sub(expand, query)
    return out if out == query else expand_macros(out, macros, depth + 1)


def pipeline_stages(query):
    """Split a query on the pipes that separate stages."""
    out, cur, quote, depth = [], "", "", 0
    i, n = 0, len(query)
    while i < n:
        ch = query[i]
        if quote:
            cur += ch
            if ch == "\\" and i + 1 < n:
                # Only the quote character is escaped in SPL, but the splitter
                # cannot tell `\"` from `\\` without decoding, and skipping the
                # pair is right for both: it keeps a rex pattern's \| and \"
                # from being read as a stage boundary or a string end.
                i += 1
                cur += query[i]
            elif ch == quote:
                quote = ""
            i += 1
            continue
        if ch in "\"'":
            quote = ch
        elif ch in "[({":
            depth += 1
        elif ch in "])}":
            depth = max(0, depth - 1)
        elif ch == "|" and depth == 0:
            out.append(cur)
            cur = ""
            i += 1
            continue
        cur += ch
        i += 1
    out.append(cur)
    return out


def truncate_at_unresolved_macro(query):
    """The pipeline up to the first stage that still calls a macro.

    A detection ends with a filter macro the security_content repository does
    not ship an expansion for. Cutting there keeps the part of the pipeline the
    parser can be held to, instead of recording the same refusal at a backtick
    for every detection.
    """
    stages = [s.strip() for s in pipeline_stages(query)]
    keep = []
    for stage in stages:
        if "`" in stage:
            break
        if stage:
            keep.append(stage)
    return " | ".join(keep)


def security_content_tree(cache_dir):
    """Extract the pinned security_content tarball; return root and temp dir."""
    url = "https://codeload.github.com/splunk/security_content/tar.gz/%s" % (
        SECURITY_CONTENT_REF)
    body = fetch_bytes(url, cache_dir)
    if body is None:
        sys.exit("could not download %s" % url)
    tmp = tempfile.mkdtemp(prefix="openspl-security-content-")
    archive = os.path.join(tmp, "source.tar.gz")
    with open(archive, "wb") as f:
        f.write(body)
    with tarfile.open(archive) as tar:
        if hasattr(tarfile, "data_filter"):
            tar.extractall(tmp, filter="data")
        else:
            tar.extractall(tmp)
    os.remove(archive)
    roots = [os.path.join(tmp, n) for n in os.listdir(tmp)]
    return roots[0], tmp


def gather_security_content(root):
    """Map detection path -> queries, with macros expanded."""
    macros = macro_definitions(root)
    out = {}
    for dirpath, _, names in os.walk(os.path.join(root, "detections")):
        for name in sorted(names):
            if not name.endswith(".yml"):
                continue
            path = os.path.join(dirpath, name)
            rel = os.path.relpath(path, root).replace(os.sep, "/")
            query = detection_search(path)
            if not query:
                continue
            query = truncate_at_unresolved_macro(
                expand_macros(" ".join(query.split()), macros))
            if query:
                out.setdefault(rel, []).append(query)
    return out


def collect(args):
    """Return the sources, {(dialect, query): entry}, and per-version counts."""
    seen = {}
    sources = []

    if args.dialect in ("spl1", "all"):
        versions = args.versions.split(",") if args.versions else SPL1_VERSIONS
        sources.append({
            "key": "spl1",
            "dialect": "spl1",
            "kind": "documented",
            "source": "Splunk Enterprise search reference, documented examples",
            "base": SPL1_BASE,
            "versions": versions,
            "groups": SPL1_GROUPS,
        })
    if args.dialect in ("spl1", "all") and not args.skip_detections:
        sources.append({
            "key": "spl1-detections",
            "dialect": "spl1",
            "kind": "derived",
            "source": "Splunk security_content detections, macro-expanded",
            "base": SECURITY_CONTENT_URL,
            "versions": [SECURITY_CONTENT_REF],
            "transform": SECURITY_CONTENT_TRANSFORM,
            "license": "Apache-2.0",
            "gather": gather_security_content,
        })
    if args.dialect in ("spl2", "all"):
        sources.append({
            "key": "spl2",
            "dialect": "spl2",
            "kind": "documented",
            "source": "Splunk Cloud Platform SPL2 search reference, documented examples",
            "base": SPL2_BASE,
            "versions": ["current"],
            "groups": SPL2_GROUPS + SPL2_COMMAND_GROUPS,
        })

    counts = {}
    empty = []
    failed = 0
    with concurrent.futures.ThreadPoolExecutor(
            max_workers=args.workers) as pool:
        for src in sources:
            per_version = {}
            root = temp = None
            if src.get("gather"):
                root, temp = security_content_tree(args.cache)
            try:
                for version in src["versions"]:
                    if src.get("gather"):
                        pages = src["gather"](root)
                    else:
                        base = src["base"] + (
                            "/" + version if version != "current" else "")
                        pages = {}
                        for group in src["groups"]:
                            got_pages, group_failed = gather_group(
                                base, group, args.cache, pool, src["dialect"])
                            failed += group_failed
                            pages.update(got_pages)
                    got = {}
                    for page, queries in pages.items():
                        for q in queries:
                            entry = seen.setdefault(
                                (src["key"], q),
                                {"pages": set(), "versions": set()})
                            entry["pages"].add(page)
                            entry["versions"].add(version)
                            got[q] = True
                    per_version[version] = len(got)
                    print("%s %s: %d queries" % (src["key"], version,
                                                 len(got)), file=sys.stderr)
                    if not got:
                        empty.append((src["key"], version))
            finally:
                if temp:
                    shutil.rmtree(temp, ignore_errors=True)
            counts[src["key"]] = per_version
    # A page that failed while its neighbours succeeded would drop its queries
    # from the corpus without changing anything else, and TestCorpus cannot tell
    # that from the documentation having removed them. Refuse to write instead;
    # the cache makes the retry cheap.
    if failed and not args.allow_partial:
        sys.exit("%d page fetches failed, so the corpus would lose their "
                 "queries; re-run to retry them, or pass --allow-partial to "
                 "write the corpus anyway" % failed)
    if empty:
        # A version with no pages is almost always one whose documentation
        # redirects to the login page, and recording it would claim coverage
        # the corpus does not have.
        for key, version in empty:
            print("%s %s yielded no queries: is that version public?"
                  % (key, version), file=sys.stderr)
        sys.exit("refusing to write a corpus with an empty version")
    return sources, seen, counts


def existing_outcomes(path):
    """The outcome baseline of a previous run, keyed by query text.

    Re-gathering must not throw away the recorded outcomes: they are a
    snapshot of the parser's behaviour, and losing them would turn the next
    run into a full baseline diff with no behaviour change behind it.
    """
    outcomes = {}
    if not os.path.exists(path):
        return outcomes
    with open(path, encoding="utf-8") as f:
        for line in f:
            line = line.strip()
            if not line:
                continue
            try:
                entry = json.loads(line)
            except ValueError:
                continue
            if entry.get("expect"):
                outcomes[entry["query"]] = entry["expect"]
    return outcomes


def write(source, seen, per_version, root):
    """Write one collection's metadata.json and queries.jsonl."""
    outdir = os.path.join(root, source["key"])
    os.makedirs(outdir, exist_ok=True)
    query_path = os.path.join(outdir, "queries.jsonl")
    outcomes = existing_outcomes(query_path)

    lines = []
    for (key, query), entry in seen.items():
        if key != source["key"]:
            continue
        line = {
            "query": query,
            "pages": sorted(entry["pages"]),
            "versions": sorted(entry["versions"]),
        }
        if query in outcomes:
            line["expect"] = outcomes[query]
        lines.append(line)
    lines.sort(key=lambda e: (e["query"]))

    with open(query_path, "w", encoding="utf-8") as f:
        for entry in lines:
            f.write(json.dumps(entry, ensure_ascii=False,
                               separators=(",", ":")) + "\n")

    meta = {
        "corpus_format": 1,
        "dialect": source["dialect"],
        "kind": source.get("kind", "documented"),
        "source": source["source"],
        "base_url": source["base"],
        "versions": source["versions"],
        "retrieved": time.strftime("%Y-%m-%d"),
        "queries": len(lines),
        "queries_per_version": per_version,
        "gatherer": "scripts/gather-corpus.py",
    }
    for key in ("transform", "license"):
        if source.get(key):
            meta[key] = source[key]
    if source.get("kind") == "derived":
        meta["note"] = (
            "Queries are the repository's own detection searches, transformed "
            "as recorded in `transform`, stored as syntax coverage. The "
            "version is the repository commit, never inferred from the query "
            "text."
        )
    else:
        meta["note"] = (
            "Queries are the documentation's own example blocks, stored as "
            "syntax coverage. The version is metadata, never inferred from "
            "the query text."
        )
    with open(os.path.join(outdir, "metadata.json"), "w",
              encoding="utf-8") as f:
        json.dump(meta, f, indent=2, ensure_ascii=False)
        f.write("\n")
    print("%s: %d queries" % (source["key"], len(lines)), file=sys.stderr)


def main():
    p = argparse.ArgumentParser(description=__doc__.split("\n", 1)[0])
    p.add_argument("--root", default="conformance/corpus",
                   help="corpus directory (default: conformance/corpus)")
    p.add_argument("--cache", default=os.path.expanduser(
        "~/.cache/openspl/docs"), help="page cache directory")
    p.add_argument("--dialect", default="all", choices=["all", "spl1", "spl2"])
    p.add_argument("--versions", default="",
                   help="comma-separated Splunk Enterprise versions")
    p.add_argument("--skip-detections", action="store_true",
                   help="skip the security_content detection collection")
    p.add_argument("--no-cache", action="store_true",
                   help="ignore the page cache")
    p.add_argument("--allow-partial", action="store_true",
                   help="write the corpus even if page fetches failed")
    p.add_argument("--workers", type=int, default=8,
                   help="concurrent page fetches (default: 8)")
    args = p.parse_args()
    if args.no_cache:
        args.cache = ""

    sources, seen, counts = collect(args)
    for src in sources:
        write(src, seen, counts[src["key"]], args.root)
    if not seen:
        sys.exit("no queries gathered")


if __name__ == "__main__":
    main()
