# Architecture

OpenSPL is a compiler with a hard boundary between source languages and target
engines. Front-ends lower syntax into one IR; back-ends read only that IR.

```
 SPL v1 ──► spl1 parser ─┐
                         ├──► unified IR ──► mapping ──► ES|QL emitter
 SPL v2 ──► spl2 parser ─┘                          └──► PromQL emitter
```

## Design goals

| Goal | Consequence |
|---|---|
| Deterministic output | No timestamps, randomness, map iteration, or network calls in the translation path. |
| Explainable failure | Every error and warning carries `line`, `column`, `offset`, a stable code, and a message. |
| Target independence | Parsers never import emitters. Emitters never parse source syntax. |
| Additive growth | A new language is a front-end; a new engine is a back-end; neither rewrites the other. |
| Zero supply chain | The CLI has no third-party Go modules. |
| Offline by default | Queries can contain sensitive identifiers; the process never sends them anywhere. |

## Package map

| Package | Depends on | Must never know about |
|---|---|---|
| `pkg/ast` | stdlib only | Parsers, emitters, mapping |
| `pkg/diag` | `pkg/ast` | Targets |
| `pkg/parser/spl1` | `pkg/ast`, `pkg/diag` | ES\|QL, PromQL, mapping |
| `pkg/parser/spl2` | `pkg/ast`, `pkg/diag` | ES\|QL, PromQL, mapping |
| `pkg/mapping` | stdlib only | Queries |
| `pkg/emitter/esql` | `pkg/ast`, `pkg/diag`, `pkg/mapping` | SPL syntax |
| `pkg/emitter/promql` | `pkg/ast`, `pkg/diag`, `pkg/mapping` | SPL syntax |
| `pkg/compiler` | front-ends, back-ends | CLI concerns |
| `cmd/openspl` | `pkg/compiler`, `pkg/mapping` | Emitter internals |

## The IR

`pkg/ast` models a pipeline as an ordered slice of typed stages:

```go
type Pipeline struct { Stages []Stage }

type Stage interface {
    StageName() string
    Position() Pos
}
```

Supported nodes:

| Node | Purpose |
|---|---|
| `SearchStage` | Search/where filters plus extracted `earliest`/`latest` bounds |
| `StatsStage` | Aggregations with optional grouping |
| `TimechartStage` | Aggregations bucketed by time |
| `EvalStage` | Field assignments |
| `RexStage` | Named-capture regex extraction |
| `TableStage`, `FieldsStage` | Column selection and removal |
| `SortStage`, `HeadStage` | Ordering and truncation |
| `BinStage` | Fixed-width bucketing |
| `FieldRef`, `StringLit`, `NumberLit`, `DurationLit` | Scalars |
| `SearchTerm`, `WildcardLit` | Full-text and wildcard atoms |
| `BinaryExpr`, `UnaryExpr`, `CallExpr` | Expressions |

Expressions are deliberately small. A construct that cannot be represented
faithfully becomes a positioned `unsupported-*` diagnostic rather than an
untyped escape hatch.

## Semantic boundary: `search` vs `where`

Classic SPL changes the meaning of the right-hand side depending on context:

| Query | Right-hand side |
|---|---|
| `search host=web` | string literal `"web"` |
| `where host=web` | field reference `web` |

The parser tracks this with a context flag (`searchLiterals`) and emits
`StringLit` or `FieldRef` accordingly. Getting this wrong would silently turn a
literal into a column reference, so it is covered by table-driven tests.

Other classic-SPL behaviors encoded in the front-end:

| Behavior | Example | IR result |
|---|---|---|
| Implicit AND | `error timeout` | `BinaryExpr{Op: "and"}` |
| Wildcard values | `host=web*` | `StringLit{"web*"}` |
| Adjacent wildcard tokens | `web*`, `*web` | `SearchTerm` with the full atom |
| Time-bound lifting | `earliest=-24h latest=now` | `SearchStage.Earliest/Latest` |
| Space-separated field lists | `stats count by host source` | two `FieldRef` group keys |
| Leading pipe | `\| stats count` | `StatsStage` after an implicit match-all |

## ES|QL emitter decisions

| SPL | ES\|QL | Why |
|---|---|---|
| `index=web`, `index=a OR index=b` | `FROM web`, `FROM a, b` | Index selection is a source, not a row filter |
| `status=500` | `status == 500` | ES\|QL equality is `==` |
| `host=web*` | `` `host.name` LIKE "web*" `` | Preserves wildcard semantics |
| `stats count by host` | ``STATS count = COUNT(*) BY `host.name` `` | ES\|QL needs `alias = expression` |
| `timechart span=5m` | ``BUCKET(`@timestamp`, 5 minutes)`` | Native time bucketing |
| `rex "(?<u>\S+)"` | `GROK "... %{NOTSPACE:u}"` | Grok is ES\|QL's extraction primitive |
| `table a b` | `KEEP a, b` | Column projection |
| `fields - _raw` | `DROP message` | Column removal |
| `sort -bytes` | `SORT bytes DESC` | Direction is explicit |
| `head 5` | `LIMIT 5` | Truncation |

Identifiers are emitted unquoted when simple and backtick-quoted otherwise
(`host.name`, `@timestamp`), because ECS field names contain dots and `@`.

Index filters after the first stage cannot move the `FROM` clause. They are
removed and reported as `late-index-filter` warnings instead of silently
changing which rows are scanned.

## PromQL guardrails

SPL returns rows; PromQL returns time series. The emitter therefore validates
the shape of the pipeline before producing anything:

| Check | Failure code |
|---|---|
| Pipeline contains an aggregation | `promql-no-aggregation` |
| Exactly one aggregation | `promql-multiple-aggs` |
| Aggregation is `count` | `promql-unsupported-agg` |
| Filters appear before the aggregation | `promql-search-after-agg` |
| No row-shaping stages (`eval`, `rex`, `table`, `fields`, `bin`) | `promql-unsupported-stage` |
| Filters map to `=` / `!=` label matchers | `promql-unsupported-filter` |
| Free text is not used as a filter | `promql-fulltext` |
| No conflicting matchers for one label | `promql-conflicting-filters` |
| Range is a valid PromQL duration | `invalid-span` |

The result is `sum [by (labels)] (increase(metric{matchers}[span]))`, with
`topk`, `bottomk`, `sort`, or `sort_desc` applied when the SPL pipeline orders
results.

## Diagnostics and determinism

`pkg/diag` sorts by position, then severity, then code before rendering. Map
iteration is never used to produce output order. The same source and mapping
therefore yield identical stdout and stderr on every machine.

Exit codes are part of the contract:

| Code | Meaning |
|---|---|
| `0` | Success |
| `1` | Parse/semantic failure, or warning with `--strict` |
| `2` | Usage, flag, mapping, or input error |

## Testing strategy

| Layer | Test type |
|---|---|
| Lexer | Token and position tests |
| Parser | Stage shape, semantics, precedence, error codes |
| Mapping | Defaults, strict JSON, label sanitization |
| Emitters | Golden strings and diagnostic codes |
| Conformance (offline) | `conformance/cases/**/*.spl` vs `.esql` / `.promql` goldens, plus diagnostics |
| Conformance (engines) | `conformance/engines/cases.tsv` executed against Splunk, Elasticsearch, and Prometheus |

## Extending the compiler

1. **New SPL command:** add IR nodes if needed, lex/parse it, then update every
   emitter or return an explicit unsupported diagnostic.
2. **New target:** create `pkg/emitter/<target>`, accept `*ast.Pipeline` and a
   mapping, return a query plus diagnostics. Register it in `pkg/compiler`.
3. **New source dialect:** create `pkg/parser/<dialect>` and lower into the same
   IR. Do not share parser internals unless the grammar is actually shared.

See [ADRs](adr/) for the decisions behind Go, the unified IR, and the
"refuse rather than guess" policy.
