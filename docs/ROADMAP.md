# Roadmap

The goal is a trustworthy translation layer between observability query
languages — not a Splunk clone, not an Elasticsearch client, and not a model.

## Principles

| Principle | What it means in practice |
|---|---|
| Determinism first | Same input + mapping = same bytes, forever within a major version. |
| Refuse rather than guess | Missing semantics are errors with positions, not approximations. |
| Evidence over claims | Every mapping gets a golden test and, where possible, an engine-level check. |
| Offline | No telemetry, no accounts, no network in the translation path. |
| Boring dependencies | One parser runtime and its transitive helper, both pinned and recorded in [ADR 0004](adr/0004-antlr-parsers.md); nothing else third-party. |

## Status

| Phase | Scope | State |
|---|---|---|
| **0.1 — Foundation** | IR, SPL v1 parser, ES\|QL emitter, PromQL emitter, mapping, CLI, corpus, CI | ✅ Shipped |
| **0.2 — Validation** | Docker harness for Splunk, Elasticsearch, Prometheus; parity fixtures | 🧪 In progress |
| **0.3 — SPL v2** | Separate grammar, `from` pipelines, `where`, `stats` | 🚧 Subset shipped |
| **0.4 — Command coverage** | `dedup`, `top`, `rare`, `eventstats`, `streamstats`, `rename`, `fillnull` | 🗺️ Planned |
| **0.5 — Migration tooling** | Batch mode, saved-search and dashboard extraction, diff reports | 💡 Proposed |
| **1.0 — Stable contract** | Frozen diagnostic codes and flags, semantic versioning, signed releases | 🎯 Target |

## 0.1 — Foundation (done)

- [x] Unified IR (`pkg/ast`)
- [x] Positioned diagnostics (`pkg/diag`)
- [x] Classic SPL lexer and parser (ANTLR grammar, generated Go)
- [x] `search`, `where`, `stats`, `timechart`, `eval`, `rex`, `table`, `fields`, `sort`, `head`, `bin`, `rename`
- [x] ES|QL emitter
- [x] PromQL emitter with semantic refusal
- [x] ECS mapping and strict JSON loading
- [x] CLI with stdin, JSON, exit codes, `--strict`
- [x] Unit tests, golden conformance suite, CI
- [x] Commit-message policy enforcement

## 0.2 — Engine validation

- [ ] Elasticsearch container with an ECS sample dataset
- [ ] Execute every engine conformance case through `POST /_query`
- [ ] Compare row counts and aggregates with Splunk for a shared fixture
- [ ] Prometheus container with synthetic counters
- [ ] Compare `increase` results for the same window
- [ ] Record engine versions in results
- [ ] Publish a compatibility report per release

## 0.3 — SPL v2

Shipped subset: `from`, `where`, `stats`, selected with `--spl-version 2`.

- [x] SPL2 lexer and grammar (separate ANTLR grammar, generated Go)
- [x] `from` sources, `where`, `stats`
- [x] Shared IR nodes with classic SPL where semantics match
- [x] Distinct diagnostics for dialect differences (`unsupported-dataset`,
  `unsupported-command`, `unknown-command`)
- [x] `--spl-version 1|2` (default: 1)
- [ ] `select`, `eval`, `sort`, `head`, `fields`, `table`, `timechart`
- [ ] Typed expressions and function signatures
- [ ] Namespaced datasets (data models, lookups)

## 0.4 — Broader command coverage

| Command | Why it matters | Target |
|---|---|---|
| `dedup` | Very common in alerts | ES\|QL `DEDUP` / stats trick |
| `top`, `rare` | Dashboards and triage | `STATS` + `SORT` + `LIMIT` |
| `rename` | Field normalization | `RENAME` |
| `fillnull` | Report hygiene | `EVAL` + `COALESCE` |
| `eventstats` | Window aggregates | `STATS` + join or `INLINE STATS` |
| `streamstats` | Running totals | Limited; likely refuse |
| `lookup`, `join` | Enrichment | Planned after schema support |
| `transaction` | Sessionization | Likely refuse |

## 0.5 — Migration tooling

- [ ] Directory and saved-search batch mode
- [ ] Dashboard JSON extraction
- [ ] Machine-readable migration report
- [ ] Field-coverage warnings: SPL fields absent from the mapping
- [ ] CI action for repositories of saved searches

## Explicit non-goals

| Non-goal | Reason |
|---|---|
| Executing queries itself | OpenSPL is a compiler; use Elasticsearch or Prometheus to run the result. |
| Replacing Splunk | The point is portability, not a clone. |
| LLM-based translation | Non-deterministic output cannot be reviewed or diffed as a migration artifact. |
| Silent best-effort conversion | A wrong alert is worse than a failed compile. |
| Hosted query processing | Queries can contain sensitive identifiers and must stay local. |

## How to influence the roadmap

Open an issue with:

1. The SPL query (redact sensitive values).
2. The target engine and version.
3. The expected result, if you know it.
4. What a wrong translation would cost.

Queries with real-world shape are more useful than synthetic one-liners.
