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
│   └── errors/                parser and unsupported-construct refusals
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
    "index": "logs-*"
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
| `options` | `metric`, `span`, and `index` passed to the compiler. |
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
```

Cases and their per-target checks run in parallel, bounded by `GOMAXPROCS`, so
the suite uses every core by default. Use `-parallel 1` when debugging an
interleaved failure, and `-parallel N` to raise or lower the bound.

```sh
go test ./conformance -parallel 1 -v
```

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
