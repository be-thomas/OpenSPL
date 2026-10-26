# Changelog

All notable changes to this project are documented here. The format follows
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and the project uses
[Semantic Versioning](https://semver.org/).

## [Unreleased]

### Added

- Parsing for both dialects is generated from ANTLR grammars
  (`pkg/parser/spl1/grammar/SPL1.g4`, `pkg/parser/spl2/grammar/SPL2.g4`). The
  generated Go is committed, so builds and releases need neither Java nor the
  ANTLR jar; `make generate-check` fails on drift.
- A working SPL2 subset: `from <dataset>`, `where`, and `stats`, selected with
  `--spl-version 2` (or `compiler.Options.SPLVersion`). `from main` lowers to
  the same IR as `index=main`, so all emitters and mappings apply unchanged.
- `unsupported-dataset` for namespaced SPL2 datasets and
  `unsupported-spl-version` for unknown dialect numbers.
- Conformance cases under `cases/spl2/` and an `options.spl_version` metadata
  key; unknown versions fail the case instead of silently running as classic
  SPL.
- Fuzz targets `FuzzParseSPL1` and `FuzzParseSPL2`, and a lexer-parity test
  that tokenizes one corpus with both generated lexers.

### Changed

- The module graph is no longer empty: the binary links the ANTLR Go runtime
  and its transitive `golang.org/x/exp` helper. [ADR 0004](docs/adr/0004-antlr-parsers.md)
  records the exception and the pinned version pair.
- SPL2 no longer returns `spl2-not-implemented`; constructs outside the subset
  are refused with a positioned `unsupported-command`, `unknown-command`, or
  `unsupported-dataset` diagnostic.

### Fixed

- ES|QL aggregations now use the valid `alias = COUNT(*)` form instead of the
  invalid `COUNT(*) AS alias`.
- Quoted phrases in a search are lowered to search terms rather than string
  literals, so `"failed password"` keeps its full-text meaning.
- Wildcard and hyphenated search values (`web-*`, `*error*`, `web*prod`,
  `logs-2024-01`) stay single values instead of fracturing into arithmetic.
- An `OR` that mixes an index filter with a row filter is refused with
  `mixed-index-filter`; `index=a OR index=b OR index=c` flattens into `FROM`.
- Warnings recorded before a later-stage error are preserved instead of being
  replaced by the error.
- `--span` now supplies the ES|QL timechart fallback; `late-index-filter`
  names the `FROM` source that was actually emitted.
- PromQL distinguishes `head 0` from an omitted head and warns when a sort
  limit has no sort field.
- Internal diagnostics use a valid `1:1` position instead of `0:0`.
- A parenthesized group after a value (`index=web (status=500)`,
  `index=web(status=500)`) is no longer parsed as a function call in search
  context, so the index filter and free-text terms survive. Function calls
  remain available in `where`, `eval`, and aggregate arguments.
- Search values with leading wildcards or signs, numeric ranges, and trailing
  wildcards (`*-*`, `-a`, `10-20`, `200-299`, `500*`) stay single literals
  instead of fracturing into arithmetic.
- `earliest` and `latest` are lifted out of `OR` branches as well as `AND`,
  and conflicting values are refused with `conflicting-time-bounds`. A lifted
  bound is TRUE in the row expression, so a bound that is a direct `OR` operand
  leaves no residual filter, and `NOT` around a bound is dropped.
- ES|QL parenthesizes a top-level `OR` next to time bounds, so the time window
  cannot be silently widened by operator precedence.
- PromQL refuses distinct positive filters for one label (`host=a host=b`)
  instead of turning their conjunction into an alternation; negative filters
  still merge.
- Columns are documented and tested as byte columns, and invalid multibyte
  input is reported as the decoded character.

### Planned

- Docker parity harness results
- `dedup`, `top`, `rare`, `rename`, `fillnull`
- Batch migration reports

## [0.1.0] — 2026-10-26

### Added

- Unified IR for pipelines, stages, expressions, and source positions.
- Classic SPL lexer and parser:
  - implicit and explicit `search`
  - `where`, `stats`, `timechart`, `eval`, `rex`, `table`, `fields`, `sort`,
    `head`, `bin`
  - implicit AND, boolean and arithmetic precedence, parentheses
  - wildcard values and search terms
  - `earliest` / `latest` extraction
  - recognized-but-unsupported command diagnostics
- Elastic ES|QL emitter:
  - `FROM`, `WHERE`, `STATS`, `BUCKET`, `EVAL`, `GROK`, `KEEP`, `DROP`, `SORT`,
    `LIMIT`
  - index extraction and `index=a OR index=b` support
  - ECS identifier quoting
  - relative and ISO 8601 time bounds
  - aggregation aliases and `COUNT_DISTINCT`
- Prometheus PromQL emitter:
  - label matchers, wildcard regex, matcher merging
  - `sum by (...) (increase(...[span]))`
  - `topk`, `bottomk`, `sort`, `sort_desc`
  - semantic refusal for row-shaped pipelines and non-count aggregations
- ECS-oriented mapping with strict JSON loading and Prometheus label mapping.
- CLI:
  - `transpile` / `translate`, `ast`, `targets`, `version`, `help`
  - stdin and file input
  - `--target`, `--mapping`, `--index`, `--metric`, `--span`, `--json`,
    `--strict`
  - stable exit codes
- Table-driven parser, mapping, emitter, and conformance tests.
- Checked-in conformance suite with golden ES|QL and PromQL files and recorded diagnostics.
- CI for formatting, vet, tests, build, and commit-message policy.
- Documentation: architecture, compatibility, mapping, roadmap, ADRs,
  contributing, security.
- Apache-2.0 license.

### Known limitations

- PromQL supports only a single `count` aggregation.
- `rex` maps to Grok with lossy type inference and always warns.
- Free-text search maps to `MATCH` and may not match the Splunk analyzer.
- Subsearches, joins, lookups, and most SPL commands are not implemented.
- SPL2 returns `spl2-not-implemented`.
- Engine validation runs against Elasticsearch and Prometheus in the `engines (docker)` CI job; Splunk validation is optional.

[Unreleased]: https://github.com/be-thomas/openspl/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/be-thomas/openspl/releases/tag/v0.1.0
