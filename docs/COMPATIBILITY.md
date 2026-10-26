# Compatibility

This page is the support contract for v0.1.x. A construct is either **exact**,
**approximate** (emitted with a warning), or **refused** (positioned error).
Nothing is silently rewritten.

## Input dialects

Select the input dialect with `--spl-version 1|2` (default `1`) or
`compiler.Options.SPLVersion`. In the API, `0` and `1` both select classic SPL;
any other value is refused with `unsupported-spl-version` at 1:1. The CLI
accepts only `1` and `2` and exits with a usage error for anything else. The
compiler never falls back to classic SPL.

| Version | Dialect | Status |
|---|---|---|
| `1` (and `0` in the API) | Classic Splunk SPL | ✅ Supported, see below |
| `2` | SPL2 subset | ✅ Supported: `from`, `where`, `stats` |

The SPL2 subset differs from classic SPL in surface syntax and semantics:

| Aspect | SPL2 subset |
|---|---|
| Pipeline start | Must be `from <dataset>`; a leading pipe is not accepted |
| Dataset names | Plain index names, including `logs-*` and `logs-2024-01` |
| Namespaced datasets | ⛔ `unsupported-dataset` (`datamodel:…`, `lookup:…`) |
| Bare words | Field references on both sides of a comparison |
| Conditions | Explicit `AND` is required; there is no implicit AND |
| `stats` | Parenthesized aggregate calls; `distinct_count` normalizes to `dc` |
| Wildcards | Only in a dataset name; `where status=*` is `syntax` |
| Other commands | ⛔ `unsupported-command`, or `unknown-command` for a non-command word |

`from main` lowers to the same IR as `index=main`, so index mapping, `FROM`
emission, and PromQL `job` matchers behave exactly as for classic SPL.

## Commands

| SPL command | Parser | ES\|QL | PromQL |
|---|---|---|---|
| `search` / implicit search | ✅ | ✅ exact filters; ⚠️ free text | ⛔ unless it is a field filter before an aggregation |
| `where` | ✅ | ✅ | ⚠️ `=` / `!=` only |
| `stats` | ✅ | ✅ | ⚠️ single `count` |
| `timechart` | ✅ | ✅ | ✅ |
| `eval` | ✅ | ✅ known functions; ⚠️ unknown passthrough | ⛔ |
| `rex` | ✅ | ⚠️ regex → Grok | ⛔ |
| `table` | ✅ | ✅ | ⛔ |
| `fields` | ✅ | ✅ | ⛔ |
| `sort` | ✅ | ✅ | ⚠️ ordering only after an aggregation |
| `head` | ✅ | ✅ | ⚠️ only with `sort` |
| `bin` | ✅ | ⚠️ derived column | ⛔ |
| `rename` | ✅ | ✅ one `RENAME`; ⛔ wildcards and interacting pairs | ⛔ |

## Expressions

| Construct | Syntax | Notes |
|---|---|---|
| Comparisons | `=`, `!=`, `<`, `<=`, `>`, `>=` | `==` accepted |
| Boolean | `AND`, `OR`, `NOT`, `!` | Implicit AND between terms |
| Arithmetic | `+`, `-`, `*`, `/`, `%` | Standard precedence |
| Parentheses | `( ... )` | |
| Wildcards | `field=web*`, `field=*web` | ES\|QL `LIKE`; PromQL `=~` |
| Quoted strings | `"text"`, `'text'` | Backslashes preserved for regex |
| Numbers | `42`, `3.14`, `1e3` | |
| Durations | `5m`, `24h`, `1mon` | Used by `earliest`, `span`, PromQL ranges |
| Function calls | `lower(x)`, `if(a,b,c)` | ES\|QL mapping for a known subset |
| Subsearches | `[search ...]` | ⛔ `unsupported-subsearch` |
| Assignment in eval | `eval x = expr` | Multiple comma-separated assignments |

### ES|QL function mapping

| SPL | ES\|QL |
|---|---|---|
| `if(cond, a, b)` | `CASE(cond, a, b)` |
| `isnull(x)` | `x IS NULL` |
| `isnotnull(x)` / `notnull(x)` | `x IS NOT NULL` |
| `now()` | `NOW()` |
| `tonumber(x)` | `TO_DOUBLE(x)` |
| `tostring(x)` | `TO_STRING(x)` |
| `lower`, `upper`, `length`, `abs`, `round`, `coalesce`, `concat`, `trim`, `replace`, `substring`, `ceil`, `floor`, `sqrt`, `pow` | Uppercased pass-through |

Unknown functions are passed through uppercased with an `unmapped-function`
warning. Verify them against your Elasticsearch version.

## Aggregations

| SPL | ES\|QL | PromQL |
|---|---|---|
| `count` | `COUNT(*)` | ✅ (as counter increase) |
| `count(field)` | `COUNT(field)` | ⚠️ counts increments, not non-null values |
| `dc(field)` / `distinct_count(field)` | `COUNT_DISTINCT(field)` | ⛔ |
| `sum(field)` | `SUM(field)` | ⛔ |
| `avg(field)` | `AVG(field)` | ⛔ |
| `min(field)` / `max(field)` | `MIN` / `MAX` | ⛔ |
| everything else | ⛔ `unsupported-agg` | ⛔ |

## Time bounds

| SPL | ES\|QL |
|---|---|---|
| `earliest=-24h` | `@timestamp >= NOW() - 24 hours` |
| `latest=now` | omitted (open-ended) |
| `earliest="2026-01-01T00:00:00Z"` | `@timestamp >= TO_DATETIME("2026-01-01T00:00:00Z")` |
| `latest=-1h` | `@timestamp <= NOW() - 1 hour` |

`earliest` and `latest` are global search-time modifiers, so they are lifted
out of both `AND` and `OR` branches. Conflicting values for the same bound are
refused with `conflicting-time-bounds`.

PromQL does not embed search windows. `earliest`/`latest` produce a
`time-bounds-external` info diagnostic; set the evaluation range on the query
API instead.

## Index selection

| SPL | ES\|QL |
|---|---|---|
| `index=web` | `FROM web` |
| `index=a OR index=b` | `FROM a, b` |
| `index=web-*` | `FROM web-*` |
| no index | mapping `default_index` (`logs-*`) |
| index filter after stage 1 | ⚠️ `late-index-filter`; ignored |
| `index=web OR host=api` | ⛔ `mixed-index-filter`; `FROM` cannot mix index and row filters |

Index names pass through the optional `indexes` map in the mapping file.
PromQL has no `FROM` clause: index filters become `job` matchers, so an `OR`
over index filters is refused with `promql-unsupported-filter`.

## Refused SPL commands

Recognized but not implemented in v0.1 (classic SPL; in SPL2, `from` is the start of a pipeline):

`append`, `appendcols`, `addtotals`, `analyzefields`, `anomalydetection`,
`apply`, `associate`, `bucket`, `chart`, `cluster`, `collect`, `contingency`,
`correlate`, `datamodel`, `dbinspect`, `dedup`, `delete`, `diff`, `eventcount`,
`eventstats`, `extract`, `fieldformat`, `fieldsummary`, `filldown`, `fillnull`,
`findtypes`, `fit`, `folderize`, `foreach`, `format`, `from`, `geostats`,
`highlight`, `inputcsv`, `inputlookup`, `iplocation`, `join`, `kmeans`, `kvform`,
`loadjob`, `localize`, `lookup`, `makecontinuous`, `makeresults`, `map`,
`mcollect`, `metadata`, `metasearch`, `ml`, `mstats`, `multikmeans`,
`multisearch`, `mvcombine`, `mvexpand`, `nomv`, `outlier`, `outputcsv`,
`outputlookup`, `outputtext`, `pivot`, `predict`, `rangemap`, `rare`, `regex`,
`reltime`, `replace`, `rest`, `return`, `reverse`, `script`, `scrub`,
`select`, `selfjoin`, `sendemail`, `set`, `setfields`, `spath`, `strcat`,
`streamstats`, `tags`, `tail`, `top`, `transaction`, `transpose`, `trendline`,
`tscollect`, `tstats`, `typeahead`, `typer`, `union`, `uniq`, `untable`, `x11`,
`xmlkv`, `xmlunescape`, `xpath`.

They return `unsupported-command` with the command's source position.

## Diagnostic codes

### Parse

| Code | Meaning |
|---|---|
| `empty-query` | Whitespace-only or missing query |
| `syntax` | Unexpected token, missing operand, malformed stage |
| `unsupported-command` | Known SPL command that is not implemented |
| `unknown-command` | Identifier that is not a known command |
| `unsupported-subsearch` | `[ ... ]` subsearch |
| `conflicting-time-bounds` | Same `earliest`/`latest` bound with different values |
| `unsupported-dataset` | SPL2 dataset is namespaced and cannot become an index |
| `unsupported-spl-version` | Dialect version is not 1 or 2 |

### ES|QL emitter

| Code | Severity | Meaning |
|---|---|---|
| `fulltext-approximation` | warning | Free text → `MATCH`; analyzer may differ |
| `rex-lossy` | warning | Regex → Grok; verify captures |
| `unmapped-function` | warning | Function passed through |
| `default-span` | warning | No span; mapping default used |
| `late-index-filter` | warning | Index filter after `FROM` was ignored |
| `mixed-index-filter` | error | `OR` mixes an index filter with a row filter |
| `bin-bucket` | info | `bin` created a derived column |
| `invalid-time-bound` | error | Unparseable time bound |
| `invalid-span` | error | Unparseable duration |
| `invalid-aggregation` | error | Aggregation missing an argument |
| `unsupported-agg` | error | No ES\|QL mapping |
| `invalid-group-by` | error | Group key is not a field |
| `rex-no-captures` | error | Pattern has no named capture |
| `rename-wildcard` | error | `rename` wildcard has no ES\|QL form |
| `rename-duplicate` | error | One field is renamed twice in a command |
| `rename-conflict` | error | Pairs would interact: a target is another source, or two pairs share a target |
| `rename-source-dropped` | error | A later stage references a name that `rename` moved to a new name |
| `unsupported-stage` | error | No emitter for the IR node |

### PromQL emitter

| Code | Severity | Meaning |
|---|---|---|
| `promql-window` | warning | Range window replaces the search range |
| `promql-count-field` | warning | `count(field)` counts increments |
| `promql-sort-ignored` | warning | Sort before aggregation is dropped |
| `promql-head-ignored` | warning | `head` without sort, or `head 0`, is dropped |
| `promql-sort-fields` | warning | Only one sort key is expressible |
| `promql-sort-limit-ignored` | warning | A sort limit without a sort field is dropped |
| `default-span` | warning | No span; mapping default used |
| `time-bounds-external` | info | Set the range on the query API |
| `promql-no-aggregation` | error | PromQL needs an aggregation |
| `promql-multiple-aggs` | error | Only one aggregation is expressible |
| `promql-unsupported-agg` | error | Only `count` is supported |
| `promql-search-after-agg` | error | Filters after aggregation |
| `promql-unsupported-stage` | error | Row-shaping stage |
| `promql-unsupported-filter` | error | Filter is not a label matcher |
| `promql-fulltext` | error | Free text in a PromQL filter |
| `promql-conflicting-filters` | error | Same label with incompatible or distinct matchers |
| `promql-head-before-agg` | error | `head` changes the aggregation input |
| `invalid-span` | error | Not a valid PromQL range |
| `invalid-mapping` | error | Bad Prometheus function in mapping |
| `unknown-target` | error | Unknown compiler target |

## Versioning

Diagnostic codes and CLI flags are semver-stable after v1.0. Before v1.0 they
may change; changes are recorded in [CHANGELOG.md](../CHANGELOG.md).
