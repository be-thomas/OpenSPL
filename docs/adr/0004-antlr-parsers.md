# ADR 0004 — Generate the SPL grammars with ANTLR

| Field | Value |
|---|---|
| Status | Accepted |
| Date | 2026-10-26 |
| Deciders | OpenSPL maintainers |
| Amends | [ADR 0001](0001-go-single-binary.md) follow-up (empty module graph) |

## Context

OpenSPL began with a hand-written lexer and a hand-written recursive-descent
parser for classic SPL, about 1,370 lines together. The language is
context-sensitive in ways that are easy to get subtly wrong: implicit `AND`
between search terms, literal values on the right of a search comparison versus
field references in `where`, adjacency-sensitive wildcard and hyphen runs, and
case-insensitive keywords that remain usable as identifiers. Every new command
meant editing allocation-heavy parsing code by hand.

SPL2 needs a second front-end. Duplicating the hand-written approach would
double the surface that can drift, and neither front-end would have a
declarative description a reviewer can check against the language rules.

The maintainer requires all parsing to be expressed in ANTLR.

## Decision

Describe both dialects as ANTLR 4 grammars and generate Go parsers from them.

| Item | Pinned value |
|---|---|
| Tool | ANTLR `4.13.2`, complete jar, SHA-256 `eae2dfa119a64327444672aff63e9ec35a20180dc5b8090b7a6ab85125df4d76` |
| Go runtime | `github.com/antlr4-go/antlr/v4 v4.13.1` |
| Indirect module | `golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842` (required by the runtime) |
| Grammars | `pkg/parser/spl1/grammar/SPL1.g4`, `pkg/parser/spl2/grammar/SPL2.g4` |
| Generated code | `pkg/parser/spl1/gen`, `pkg/parser/spl2/gen`, committed |
| Regeneration | `make generate` (`scripts/generate-parsers.sh`); CI runs `make generate-check` |

The generated Go is committed. Building, installing, containerising, or
releasing OpenSPL needs neither Java nor the ANTLR jar; regeneration is a
maintainer task that requires Java 17 or newer.

The grammars are the source of truth. Generated files carry a "DO NOT EDIT"
header and are never hand-edited; a change to parsing behaviour changes the
`.g4` file and regenerates. `.gitattributes` marks `pkg/parser/*/gen/**` as
generated so reviews focus on the grammar.

## Consequences

Positive:

- Parsing rules are declared in one place per dialect, including the
  context-sensitive boundaries, rather than spread across hand-written
  recursive descent.
- Adding a command is a grammar rule plus a visitor method, not new
  token-shuffling code.
- Stronger position guarantees: a source map converts ANTLR rune indexes to
  OpenSPL's byte-based `ast.Pos`.
- ANTLR's error listener and bail strategy can preserve the existing
  fail-fast contract without accepting ANTLR's own error text.

Negative:

- The module graph is no longer empty. ADR 0001's follow-up said to keep it
  empty; this ADR supersedes that for exactly these two modules, and any
  further dependency still needs its own ADR.
- A first build from source downloads the two modules into the Go module
  cache. The shipped binary still makes no network calls, and generated code
  is committed so ordinary builds never touch the network once the cache is
  warm.
- Generated Go is large and repetitive. It is excluded from `go vet`'s
  `unreachable` analyzer (upstream codegen emits a defensive `goto` after a
  `return`); every other analyzer still runs on it.
- Regenerating requires Java 17, which contributors changing a grammar must
  install. CI enforces that the committed output matches.

## Alternatives considered

| Option | Rejected because |
|---|---|
| Keep the hand-written parsers | The requirement is ANTLR, and a second dialect would double hand-written parsing code. |
| ANTLR with the grammar compiled at runtime | Would ship or require the jar and Java, contradicting the single-binary promise. |
| A parser generator with a pure-Go front-end (e.g. goyacc) | Does not meet the requirement, and lacks ANTLR's ALL(*) prediction for context-sensitive rules. |
| Vendoring the runtime | Adds a large tree to review; the module cache and `go.sum` give reproducible builds without it. |

## Upgrade procedure

1. Download the new complete jar from Maven Central and record its SHA-256 in
   `scripts/antlr-<version>.sha256`.
2. Bump the version in `scripts/generate-parsers.sh`.
3. Run `make generate`, then `make check`, `go test -race ./...`, and the
   conformance suite. Goldens must not change.
4. Bump the runtime only to the matching v4 release; if the generated code
   needs a symbol the pinned runtime lacks, pin both back together.
5. Commit the tool version, checksum, generated diff, and `go.mod`/`go.sum` in
   one change.
