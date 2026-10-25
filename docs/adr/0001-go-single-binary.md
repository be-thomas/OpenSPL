# ADR 0001 — Build in Go as a single static binary

| Field | Value |
|---|---|
| Status | Accepted |
| Date | 2026-10-26 |
| Deciders | OpenSPL maintainers |

## Context

The transpiler must be easy to run in CI, in containers, on laptops, and inside
air-gapped networks. It must not require a runtime, a package manager, or a
service. Query text is sensitive, so the tool must be auditable and offline.

The realistic choices were Go and Rust.

## Decision

Implement OpenSPL in Go with no third-party modules.

## Consequences

Positive:

- One `go build` produces a static binary for Linux, macOS, and Windows.
- The standard library covers JSON, regex, flags, and testing.
- A zero-dependency module graph is easy to audit and has no supply-chain risk.
- Cross-compilation is trivial (`GOOS`, `GOARCH`).
- Go's table-driven testing fits a compiler corpus.

Negative:

- Rust would offer stronger sum types for the IR. Go's interfaces plus type
  switches are adequate but require discipline.
- Go's flag package is minimal; small wrappers are needed for interspersed
  flags.

## Alternatives considered

| Option | Rejected because |
|---|---|
| Rust | Strong fit technically, but slower iteration and a larger toolchain for contributors; the IR does not need zero-cost abstractions. |
| TypeScript/Node | Requires a runtime and a dependency tree; not a single binary. |
| Python | Packaging and startup overhead; deployment friction in containers. |

## Follow-up

Keep the module graph empty. Any proposed dependency must justify why the
standard library is insufficient.
