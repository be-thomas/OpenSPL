# ADR 0003 — Refuse translations rather than guess

| Field | Value |
|---|---|
| Status | Accepted |
| Date | 2026-10-26 |
| Deciders | OpenSPL maintainers |

## Context

Observability queries drive paging, security response, and capacity decisions.
A translator that silently changes meaning is worse than one that fails: the
broken query may keep running and only reveal itself during an incident.

SPL and PromQL are not isomorphic. SPL returns arbitrary rows, including raw
text; PromQL returns time series. Likewise, free-text SPL search depends on the
index analyzer, and `rex` captures depend on regex dialect.

## Decision

Classify every construct into one of three tiers:

| Tier | Behavior |
|---|---|
| Exact | Emit directly. |
| Approximate | Emit with a warning code that names the risk. |
| Refused | Stop with a positioned error. |

`--strict` promotes warnings to errors for CI.

## Consequences

Positive:

- Users can review every approximation.
- CI can enforce a no-warnings policy.
- The tool never fabricates a semantic guarantee it cannot keep.
- Error codes are stable and testable.

Negative:

- Some users will be told "not supported" when a rough translation would have
  been convenient.
- PromQL coverage is intentionally narrow.

## Examples

| Input | Target | Tier | Reason |
|---|---|---|---|
| `stats count by host` | ES\|QL | Exact | Direct aggregate |
| `search timeout` | ES\|QL | Approximate | Analyzer/tokenization may differ |
| `rex "(?<u>\S+)"` | ES\|QL | Approximate | Regex → Grok type inference |
| `stats avg(bytes)` | PromQL | Refused | A row average is not a counter rate |
| `eval x = ...` | PromQL | Refused | No row-shaping in PromQL |

## Alternatives considered

| Option | Rejected because |
|---|---|
| Best-effort translation | Silent semantic drift |
| Warn but emit invalid target syntax | Fails later, far from the cause |
| Refuse all approximations | Would block useful, reviewable migrations |
