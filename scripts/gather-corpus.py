#!/usr/bin/env python3
"""Gather the documented SPL/SPL2 example queries into the corpus.

The checked-in corpus under conformance/corpus is what the tests read; this
script only refreshes it. It needs network access, so it is never run by tests
or CI.

    scripts/gather-corpus.py --cache ~/.cache/openspl/docs
    scripts/gather-corpus.py --versions 9.0,9.2,10.0
    scripts/gather-corpus.py --dialect spl2

The Splunk documentation renders its pages client-side and exposes each page's
body at a fragment URL. Command example queries appear in
`<textarea class="code-raw-content">` elements, which is what this script
reads: one query per element, no prose.

Each source collection gets a metadata.json that records the dialect and the
documentation versions the queries were taken from, because a query file does
not say which Splunk version it belongs to. A query that appears in several
versions is stored once, with the versions listed on its line; a query that
appears in only some versions keeps that narrower list, so version-specific
syntax stays visible.

The expected outcome on each line (accepted, or the diagnostic code of a
refusal) is not written here. It is a snapshot of the current parser, recorded
by `go test ./conformance -run TestCorpus -update`, so a behaviour change shows
up as a reviewable corpus diff.
"""

import argparse
import concurrent.futures
import html
import json
import os
import re
import sys
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

CODE_BLOCK = re.compile(
    r'<textarea class="code-raw-content"[^>]*>(.*?)</textarea>', re.S)
HREF = re.compile(r'(?:hx-get|href)="(/en/[^"]+)"')
TAG = re.compile(r"<[^>]+>")

# Placeholder syntax in a documentation template, not a query.
PLACEHOLDER = re.compile(
    r"<[A-Za-z][^<>]*>|\[[^\]]*\]|\{[^}]*\}|\.\.\.|…|\$[A-Za-z_]\w*\$")

# A block is a query rather than a syntax template when it looks like the start
# of a search or a pipeline stage.
SEARCH_LEAD = re.compile(
    r"^(?:from\s|index\s*=|source\s*=|sourcetype\s*=|host\s*=|eventtype\s*=|"
    r"tag\s*=|search\b)", re.I)

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


def fetch(url, cache_dir, retries=3):
    """Return the page body, using the cache when it has one."""
    key = re.sub(r"[^A-Za-z0-9]+", "_", url)
    path = os.path.join(cache_dir, key) if cache_dir else None
    if path and os.path.exists(path):
        with open(path, "rb") as f:
            return f.read().decode("utf-8", "replace")
    last = None
    for attempt in range(retries):
        try:
            req = urllib.request.Request(url, headers={"User-Agent": UA})
            with urllib.request.urlopen(req, timeout=30) as r:
                body = r.read().decode("utf-8", "replace")
            if path:
                os.makedirs(cache_dir, exist_ok=True)
                with open(path, "w", encoding="utf-8") as f:
                    f.write(body)
            return body
        except (urllib.error.URLError, TimeoutError, OSError) as err:
            last = err
            time.sleep(1 + attempt)
    print("fetch failed: %s: %s" % (url, last), file=sys.stderr)
    return None


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
    """Map page path -> queries for one navigation group."""
    nav = fetch(nav_url(base, group), cache_dir)
    if nav is None:
        return {}
    pages = page_paths(nav, base)
    bodies = list(pool.map(lambda page: (page, fetch(fragment_url(base, page),
                                                     cache_dir)), pages))
    out = {}
    for page, body in bodies:
        if body is None:
            continue
        found = [q for q in code_blocks(body) if is_query(q, dialect)]
        if found:
            out[page] = found
    return out


def collect(args):
    """Return {(dialect, query): {"pages": set, "versions": set}}."""
    seen = {}
    sources = []

    if args.dialect in ("spl1", "all"):
        versions = args.versions.split(",") if args.versions else SPL1_VERSIONS
        sources.append({
            "key": "spl1",
            "dialect": "spl1",
            "source": "Splunk Enterprise search reference, documented examples",
            "base": SPL1_BASE,
            "versions": versions,
            "groups": SPL1_GROUPS,
        })
    if args.dialect in ("spl2", "all"):
        sources.append({
            "key": "spl2",
            "dialect": "spl2",
            "source": "Splunk Cloud Platform SPL2 search reference, documented examples",
            "base": SPL2_BASE,
            "versions": ["current"],
            "groups": SPL2_GROUPS + SPL2_COMMAND_GROUPS,
        })

    counts = {}
    empty = []
    with concurrent.futures.ThreadPoolExecutor(
            max_workers=args.workers) as pool:
        for src in sources:
            per_version = {}
            for version in src["versions"]:
                base = src["base"] + ("/" + version if version != "current"
                                      else "")
                got = {}
                for group in src["groups"]:
                    pages = gather_group(base, group, args.cache, pool,
                                         src["dialect"])
                    for page, queries in pages.items():
                        for q in queries:
                            entry = seen.setdefault(
                                (src["dialect"], q),
                                {"pages": set(), "versions": set()})
                            entry["pages"].add(page)
                            entry["versions"].add(version)
                            got[q] = True
                per_version[version] = len(got)
                print("%s %s: %d queries" % (src["dialect"], version,
                                             len(got)), file=sys.stderr)
                if not got:
                    empty.append((src["key"], version))
            counts[src["key"]] = per_version
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
    for (dialect, query), entry in seen.items():
        if dialect != source["dialect"]:
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
        "source": source["source"],
        "base_url": source["base"],
        "versions": source["versions"],
        "retrieved": time.strftime("%Y-%m-%d"),
        "queries": len(lines),
        "queries_per_version": per_version,
        "gatherer": "scripts/gather-corpus.py",
        "note": (
            "Queries are the documentation's own example blocks, stored as "
            "syntax coverage. The version is metadata, never inferred from "
            "the query text."
        ),
    }
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
    p.add_argument("--no-cache", action="store_true",
                   help="ignore the page cache")
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
