# Query corpus

Two kinds of query collections, gathered once and checked in, so the parser is
exercised against far more real syntax than the hand-written conformance cases
can cover.

- **Documented.** The example queries of the Splunk SPL and SPL2 search
  references, gathered across documentation versions.
- **Derived.** The searches of the Splunk `security_content` detections, with
  the repository's own macro definitions expanded. They are not published as
  example queries, so they are labelled derived and their transformation is
  recorded in `metadata.json`.

This is a breadth and robustness corpus, not the compatibility contract.
[`../cases/`](../cases/) is what OpenSPL promises to emit; the corpus is what
the parser must survive: every query either parses, or is refused with a
positioned diagnostic, and either way the outcome is recorded so a change shows
up in review.

## Layout

```text
corpus/
├── spl1/                documented, Splunk Enterprise 9.0-10.0
│   ├── metadata.json    dialect, versions, counts, provenance
│   └── queries.jsonl    one query per line
├── spl1-detections/     derived from security_content detection searches
│   ├── LICENSE          Apache-2.0, the source's licence
│   ├── metadata.json
│   └── queries.jsonl
└── spl2/                documented, SPL2 (published unversioned)
    ├── metadata.json
    └── queries.jsonl
```

Every collection has the same two data files; `LICENSE` appears only where the
source carries a licence of its own.

Each `queries.jsonl` line is one query:

```json
{"query":"index=* | stats count(eval(status=\"404\")) AS count_status BY sourcetype","pages":["search-commands/stats"],"versions":["9.0","9.1","9.2","9.3","9.4","10.0"],"expect":"accepted"}
```

| Key | Meaning |
|---|---|
| `query` | The query, whitespace-normalized to one line. |
| `pages` | Source pages the query appears on, relative to `base_url`. |
| `versions` | Versions of the source that contain the query. |
| `expect` | `accepted`, or `error:<code>` for the diagnostic the parser returns. |

A query that appears in several versions is stored once and lists every version
it appears in. A query that only one version documents keeps that narrower
list, which is how version-specific syntax stays visible.

`metadata.json` records `source`, `base_url`, `retrieved`, `versions`, the query
counts, and the gatherer that produced the files, plus these keys:

| Key | Meaning |
|---|---|
| `kind` | `documented` for the source's own examples, `derived` for a transformation. |
| `transform` | For derived collections: exactly what was done to the source text. |
| `license` | For derived collections: the source's licence. |

Required keys are validated by `TestCorpus`, so a truncated or hand-edited file
fails instead of silently shrinking the corpus.

## The version is metadata, never inferred

A query file does not say which Splunk release it belongs to, and the parser
does not try to guess: `from` is the only textual hint, and classic SPL and
SPL2 both use it. The version lives in `metadata.json` (`versions`,
`queries_per_version`) and on each line (`versions`). The documentation for
SPL2 is published unversioned, so its collection records the single entry
`current` plus the retrieval date. The derived collection records the
`security_content` commit it was taken from.

## The derived collection

`detections/**/*.yml` holds one detection per file, and each `search:` is a real
classic-SPL query. Two things stand between that text and a corpus line:

1. **Macros.** Most searches call macros, and the call alone is not a query. The
   definitions in the repository's own `macros/` are expanded first, with
   `$argument$` substitution, up to eight rounds of nesting.
2. **Undefined macros.** A detection usually ends with a filter macro that
   `security_content` does not define, because it is meant to come from an app.
   Those cannot be expanded, so the pipeline is cut at the first stage that
   still calls one. What remains is the detection's own SPL up to that point,
   with nothing invented and nothing reordered.

The queries are deduplicated by text; `pages` lists every detection that shares
a query. Duplicated detection searches are common, and storing the same text
twice would only slow the test down.

## Refreshing

The gatherer needs network access and is never run by the tests or by CI:

```sh
scripts/gather-corpus.py                      # every collection
scripts/gather-corpus.py --dialect spl2       # one dialect
scripts/gather-corpus.py --versions 9.4,10.0  # a subset of SPL versions
scripts/gather-corpus.py --skip-detections    # documentation only
```

It caches every page and archive under `~/.cache/openspl/docs`, so a re-run
after a partial failure does not refetch what it already has. Existing `expect`
values are carried over by query text, so re-gathering alone never rewrites a
baseline. A version that yields no queries aborts the run rather than recording
empty coverage, which is what a documentation version that redirects to the
login page looks like.

## Outcomes

`expect` is a snapshot of what the parser currently does with the query, not a
claim that it is right. A query the parser refuses stays refused until the
parser is fixed; fixing it changes exactly that line.

```sh
make corpus-update        # go test ./conformance -run TestCorpus -update-corpus
```

Read the diff before committing it: every changed outcome is a behaviour
change, and the corpus is how an unintended one gets noticed.

`TestCorpus` also fails when an implemented command never appears at the start
of a stage anywhere in the collections of its dialect, because a command with no
corpus query is a command nothing tests. The rule is per dialect rather than per
collection so that a derived collection, which is heavy on searches over raw
events, does not have to reproduce the documentation's command coverage.

## Provenance

Documented collections: `query` values are the documentation's own example
blocks, reproduced as short syntax specimens for parser coverage; the prose
around them is not included. They remain the property of Splunk and are
attributed here with the page and version they came from.

Derived collection: the detection searches come from
[splunk/security_content](https://github.com/splunk/security_content), pinned at
commit `a3f6523eabf696d66e20fdd80766bfae95575577`, under the Apache License 2.0;
the licence text is in [`spl1-detections/LICENSE`](spl1-detections/LICENSE).
Changes to that material, as the licence requires stating: the `search` value of
each detection was extracted, its macros were expanded from the same
repository's `macros/` definitions, the pipeline was cut at the first stage that
calls a macro the repository does not define, whitespace was normalized, and
duplicate texts were stored once. The detection title, description, and every
other field were dropped.

[`../../THIRD_PARTY_NOTICES.md`](../../THIRD_PARTY_NOTICES.md) records both
sources.
