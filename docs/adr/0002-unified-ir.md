# ADR 0002 — One unified IR for every language and target

| Field | Value |
|---|---|
| Status | Accepted |
| Date | 2026-10-26 |
| Deciders | OpenSPL maintainers |

## Context

The venture starts with classic SPL and two targets (ES|QL and PromQL). SPL2 and
more targets are planned. A direct SPL → ES|QL translator would be simple for
one pair, then become an N×M matrix of special cases.

## Decision

All front-ends lower into `pkg/ast`. All back-ends read only `pkg/ast`.

```text
SPL v1 ─┐
        ├─► IR ─┬─► ES|QL
SPL v2 ─┘       └─► PromQL
```

## Consequences

Positive:

- Adding a target is additive: one package, one registration.
- Adding a language does not touch emitters.
- Tests can target the IR directly.
- Semantic checks live in one place per target, not per source language.

Negative:

- The IR must be expressive enough for every language, or it will accumulate
  lowest-common-denominator nodes.
- Target-specific constructs need either IR nodes or a documented refusal.

## Guardrails

| Rule | Enforcement |
|---|---|
| Parsers never import emitters | CI import check / code review |
| Emitters never parse source | Package boundaries and review |
| IR nodes carry positions | `Position()` on every node |
| Unsupported constructs are refused | Stable diagnostic codes |

## Alternatives considered

| Option | Rejected because |
|---|---|
| Direct source-to-target transpilation | N×M growth; duplicated semantics |
| Target-specific ASTs | Requires re-parsing; no shared validation |
| Text templating | Cannot reason about precedence, positions, or semantics |
