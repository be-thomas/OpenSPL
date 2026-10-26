# Conformance

Everything that proves OpenSPL behaves as documented lives in the
[`conformance/`](conformance/) directory. There is no second test corpus
elsewhere in the repository.

```text
conformance/
├── conformance_test.go        runner: go test ./conformance
├── cases/                     offline, deterministic, no network
│   ├── search/                search expressions, wildcards, time bounds
│   ├── stats/                 aggregations and grouping
│   ├── timechart/             bucketed aggregation
│   ├── eval/                  expression lowering
│   ├── rex/                   extraction lowering
│   ├── shaping/               table, fields, sort, head
│   ├── mapping/               custom field and index mappings
│   ├── promql/                PromQL-specific output and refusals
│   ├── spl2/                  SPL2 subset (`spl_version: 2`)
│   └── errors/                parser and unsupported-construct refusals
├── corpus/                    query collections, one directory per source
│   ├── spl1/                  documented, with the Splunk versions as metadata
│   ├── spl1-detections/       derived from security_content detections
│   └── spl2/                  documented, published unversioned
├── fixtures/
│   ├── mappings/              mapping files referenced by case metadata
│   └── data/                  NDJSON seed data for the engines
└── engines/                   live-engine cases, run with Docker
    ├── cases.tsv
    ├── docker-compose.yml
    ├── prometheus.yml
    ├── run.sh
    └── README.md
```

The split is deliberate: `cases/` is the compatibility contract and runs on
every `go test ./...`; `engines/` proves that generated queries execute and
needs Docker, so it runs separately.

## Offline cases

Each case is one query plus its expected result per target. A case is named by
the SPL file and may have any of these siblings:

| File | Meaning |
|---|---|
| `<name>.spl` | The input query. One query per file, no comments. |
| `<name>.esql` | Exact expected ES\|QL. Its presence makes that target applicable. |
| `<name>.promql` | Exact expected PromQL. Its presence makes that target applicable. |
| `<name>.json` | Optional metadata and expectations. |

A target that has neither a golden file nor an entry in `<name>.json` is not
exercised by that case. A case must be exercised by at least one target.

### Case metadata

```json
{
  "description": "why this case exists",
  "mapping": "../../fixtures/mappings/custom-ecs.json",
  "options": {
    "metric": "http_requests_total",
    "span": "1m",
    "index": "logs-*",
    "spl_version": 2
  },
  "targets": {
    "esql": {
      "diagnostics": [
        {"code": "rex-lossy", "severity": "warning", "line": 1, "col": 14}
      ]
    },
    "promql": {
      "error": {"code": "promql-unsupported-stage", "line": 1, "col": 14}
    }
  }
}
```

| Key | Meaning |
|---|---|
| `description` | Free text. Never asserted. |
| `mapping` | Mapping file, relative to the case file. |
| `options` | `metric`, `span`, `index`, and `spl_version` passed to the compiler. |
| `options.spl_version` | Input dialect. `1` (or absent) is classic SPL, `2` is the SPL2 subset. Any other value fails the case rather than silently falling back. |
| `targets.<target>.diagnostics` | The complete warning/info list with exact positions. Order is line, then column, then severity (`error`, `info`, `warning`), then code. |
| `targets.<target>.error` | The target must refuse the query with this error code and position. No golden file may exist. |

Unknown metadata keys are rejected, so a typo fails the suite instead of
silently weakening it.

### Rules

- Output is compared exactly after trailing-newline normalization.
- Diagnostics are compared exactly, including code, severity, line, and column.
- A refusal is a result, not a failure: the query must produce the recorded
  error and no output.
- Golden files are the contract. A behavior change updates the golden in the
  same change, and the diff is reviewed.

### Adding a case

1. Pick the area directory and add `<name>.spl`.
2. Add `<name>.esql` and/or `<name>.promql` with the exact output you expect.
3. If the case warns or is refused, add `<name>.json` with the expectation.
4. Run `go test ./conformance` and fix the implementation or the golden.
5. If the golden was authored by hand and is wrong, regenerate with `make
   conformance-update`, then read the diff before committing.

`-update` refreshes only targets a case already declares; it never silently
adds a target to a case.

### Running a subset

Directories are the suite grouping, so Go's `-run` selects an area:

```sh
go test ./conformance -run 'TestConformance/promql' -v
go test ./conformance -run 'TestConformance/errors' -v
go test ./conformance -run 'TestConformance/spl2' -v
```

Cases and their per-target checks run in parallel, bounded by `GOMAXPROCS`, so
the suite uses every core by default. Use `-parallel 1` when debugging an
interleaved failure, and `-parallel N` to raise or lower the bound.

```sh
go test ./conformance -parallel 1 -v
```

## Query corpus

`corpus/` holds the queries the parser is tested against at breadth. It is not
a second compatibility contract: `cases/` is what OpenSPL promises to emit,
while the corpus is the broad input the parser has to survive. A collection is
one of two kinds:

- **Documented.** The example queries of the Splunk SPL and SPL2 search
  references, gathered across documentation versions.
- **Derived.** The searches of the Splunk `security_content` detections, with
  the repository's own macro definitions expanded. The source does not publish
  them as example queries, so the collection records its transformation and its
  licence in `metadata.json` and ships the licence text beside the queries.

`TestCorpus` runs with the offline suite and asserts three things:

- **Robustness.** Every query parses, or is refused with a positioned
  diagnostic. A panic, an error that is not a `*diag.Error`, or a position
  outside the input fails, whatever the query says.
- **Stability.** The outcome (`accepted`, or `error:<code>`) matches the
  recorded baseline in `queries.jsonl`, so a behaviour change on gathered
  syntax appears as a reviewable diff. Re-record it with `make corpus-update`
  and read the diff before committing.
- **Coverage.** Every implemented command appears at the start of at least one
  stage in the collections of its dialect. A command with no corpus query is a
  command nothing tests.

The dialect, the kind, and the source versions live in each collection's
`metadata.json`, and each query lists the versions that contain it. Nothing is
inferred from the query text: `from` is the only textual hint, and classic SPL
and SPL2 both use it.

`scripts/gather-corpus.py` refreshes the corpus from the documentation and from
the pinned `security_content` commit. It needs network access, so neither the
tests nor CI run it. See [conformance/corpus/README.md](conformance/corpus/README.md)
for the format, the cache, and the provenance.

## Engine cases

`engines/cases.tsv` lists queries that must be accepted and executed by a real
engine. Columns are tab-separated:

```text
target  name  flags  expected-rows  query
```

`-` means no flags or no row assertion. The PromQL cases use `--metric up`
against a Prometheus that scrapes itself. Seed data for Elasticsearch lives in
`fixtures/data/logs-e2e.ndjson`.

```sh
make e2e
```

See [conformance/engines/README.md](conformance/engines/README.md) for
endpoints, profiles, and the optional Splunk reference check.

## CI

The offline suite runs as part of `go test ./...` on every push and pull
request. The engine harness runs in the `e2e` CI job, which starts
Elasticsearch and Prometheus with Docker Compose and fails on any engine error
or row-count mismatch.

## Why one folder

- One place to look for evidence, instead of `tests/` growing a third runner.
- Inputs live once; targets are declared per case, so supported-in-one-target
  and refused-in-another cases stay side by side.
- Goldens and diagnostics are data, not code, so adding a case never needs Go.
- The offline suite and the engine suite check different properties and have
  different requirements, so they are separate directories rather than one
  mixed harness.
