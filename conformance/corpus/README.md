# Documented-query corpus

The example queries in the Splunk documentation, gathered once and checked in,
so the parser is exercised against far more real syntax than the hand-written
conformance cases can cover.

This is a breadth and robustness corpus, not the compatibility contract.
[`../cases/`](../cases/) is what OpenSPL promises to emit; the corpus is what
the parser must survive: every query either parses, or is refused with a
positioned diagnostic, and either way the outcome is recorded so a change shows
up in review.

## Layout

```text
corpus/
├── spl1/
│   ├── metadata.json    dialect, documentation versions, counts, provenance
│   └── queries.jsonl    one documented query per line
└── spl2/
    ├── metadata.json
    └── queries.jsonl
```

Each `queries.jsonl` line is one query:

```json
{"query":"index=* | stats count(eval(status=\"404\")) AS count_status BY sourcetype","pages":["search-commands/stats"],"versions":["9.0","9.1","9.2","9.3","9.4","10.0"],"expect":"accepted"}
```

| Key | Meaning |
|---|---|
| `query` | The example query, whitespace-normalized to one line. |
| `pages` | Documentation pages the query appears on, relative to `base_url`. |
| `versions` | Documentation versions that contain the query. |
| `expect` | `accepted`, or `error:<code>` for the diagnostic the parser returns. |

A query that appears in several versions is stored once and lists every version
it appears in. A query that only one version documents keeps that narrower
list, which is how version-specific syntax stays visible.

## The version is metadata, never inferred

A query file does not say which Splunk release it belongs to, and the parser
does not try to guess: `from` is the only textual hint, and classic SPL and
SPL2 both use it. The version lives in `metadata.json` (`versions`,
`queries_per_version`) and on each line (`versions`). The documentation for
SPL2 is published unversioned, so its collection records the single entry
`current` plus the retrieval date.

`metadata.json` also records `source`, `base_url`, `retrieved`, the query
counts, and the gatherer that produced the files. Required keys are validated
by `TestCorpus`, so a truncated or hand-edited file fails instead of silently
shrinking the corpus.

## Refreshing

The gatherer needs network access and is never run by the tests or by CI:

```sh
scripts/gather-corpus.py                      # both dialects, all versions
scripts/gather-corpus.py --dialect spl2       # one collection
scripts/gather-corpus.py --versions 9.4,10.0  # a subset of SPL versions
```

It caches every page under `~/.cache/openspl/docs`, so a re-run after a partial
failure does not refetch what it already has. Existing `expect` values are
carried over by query text, so re-gathering alone never rewrites a baseline.

## Outcomes

`expect` is a snapshot of what the parser currently does with the query, not a
claim that it is right. A documented query the parser refuses stays refused
until the parser is fixed; fixing it changes exactly that line.

```sh
make corpus-update        # go test ./conformance -run TestCorpus -update-corpus
```

Read the diff before committing it: every changed outcome is a behaviour
change, and the corpus is how an unintended one gets noticed.

`TestCorpus` also fails when an implemented command never appears at the start
of a stage anywhere in its collection, because a command with no corpus query
is a command nothing tests.

## Provenance

`query` values are the documentation's own example blocks, reproduced as short
syntax specimens for parser coverage; the prose around them is not included.
They remain the property of Splunk and are attributed here with the page and
version they came from. [`../../THIRD_PARTY_NOTICES.md`](../../THIRD_PARTY_NOTICES.md)
records the documentation as a source.
