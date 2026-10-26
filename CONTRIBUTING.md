# Contributing to OpenSPL

Thanks for helping make observability queries portable. This project values
small, reviewable changes with explicit semantics over broad, best-effort
conversion.

## Setup

```bash
git clone https://github.com/be-thomas/openspl.git
cd openspl
make hooks     # install the commit-msg hook
make check     # format check, vet, tests, build
```

Requirements: Go 1.24+ and `make`. Docker is optional, for the engine harness.
Java 17 is needed only to regenerate the parser sources, not to build or test.

## Generated parsers

Both dialects are ANTLR grammars. Generated Go is committed, so a normal build
never needs Java or the ANTLR jar.

| Path | Contents |
|---|---|
| `pkg/parser/spl1/grammar/SPL1.g4` | Classic SPL grammar (source of truth) |
| `pkg/parser/spl2/grammar/SPL2.g4` | SPL2 subset grammar (source of truth) |
| `pkg/parser/spl1/gen`, `pkg/parser/spl2/gen` | Generated lexer, parser, visitor; never hand-edited |

```bash
make generate        # rewrite the gen/ trees (needs Java 17)
make generate-check  # fail if gen/ is stale; runs in CI
```

Edit the `.g4`, run `make generate`, and commit both the grammar and the
regenerated output. A change that edits only `gen/` is rejected by review and
by `make generate-check`.

## Repository layout

| Path | Contents |
|---|---|
| `cmd/openspl` | CLI entry point |
| `pkg/ast` | Unified IR |
| `pkg/diag` | Positioned diagnostics |
| `pkg/parser/spl1` | Classic SPL front-end: stage driver, visitor, `Lex` adapter |
| `pkg/parser/spl2` | SPL2 subset front-end |
| `pkg/parser/commands` | Implemented and known command-name tables, per dialect |
| `pkg/parser/internal/antlrpos` | Byte-accurate `ast.Pos` source map |
| `pkg/parser/internal/antlrsupport` | Error listener and bail/recover |
| `pkg/parser/internal/lexerparity` | Test-only SPL1/SPL2 lexical parity gate |
| `pkg/mapping` | Field, index, and label mapping |
| `pkg/emitter/esql` | ES\|QL back-end |
| `pkg/emitter/promql` | PromQL back-end |
| `pkg/compiler` | Driver |
| `conformance/cases` | SPL queries with golden outputs and recorded diagnostics |
| `conformance/engines` | Docker harness for live-engine cases |
| `docs/adr` | Architecture decision records |

## Commit messages

**Hard rule:** every commit message must be a plain description of the change.
It must not contain AI or tool attribution, generated-by signatures, or
co-author trailers.

The rule applies to all commits, including squash merges and release commits.
It is enforced by:

1. `.githooks/commit-msg` (enable with `make hooks`)
2. `scripts/check-commit-message.sh`
3. The `commit-message` CI job

If a commit is rejected, rewrite the message:

```bash
git commit --amend
# or, for several commits:
git rebase -i HEAD~N
```

Write messages in the imperative mood, for example:

```text
Add dedup lowering for ES|QL
Fix wildcard adjacency in search terms
Reject mixed label matchers in PromQL
```

## Adding a command

1. **IR:** add a stage or expression node in `pkg/ast` if no existing node fits.
   Every node implements `StageName()`/`Position()` or `exprNode()`/`Position()`.
2. **Grammar:** add the command's rule to the relevant `.g4`, run
   `make generate`, and lower the new rule context in that front-end's
   `build.go`. The Go driver in `parse.go` only sequences stages and classifies
   command names; it is not a second parser. If the command is recognized but
   not implemented, add it to `pkg/parser/commands` so users get
   `unsupported-command` rather than a misleading syntax error.
3. **Emitters:** implement the ES|QL mapping and decide the PromQL tier
   (supported or refused). An unsupported stage must return a diagnostic with a
   position, never a partially valid query.
4. **Tests:** add parser tests, emitter golden tests, and a conformance case.
5. **Docs:** update `docs/COMPATIBILITY.md` and, if behavior is surprising,
   `docs/ARCHITECTURE.md`.

### Semantic rules

| Rule | Why |
|---|---|
| Search values are literals; `eval`/`where` identifiers are fields | Classic SPL context sensitivity |
| Unsupported targets are errors, not comments | The output must be executable or absent |
| Diagnostics are sorted by position, severity, then code | Deterministic output |
| Generated identifiers are sanitized deterministically | Stable golden files |
| Map iteration never determines output order | Reproducible builds and diffs |

## Adding a conformance case

The suite layout and metadata schema are documented in
[CONFORMANCE.md](CONFORMANCE.md).

1. Add `conformance/cases/<area>/<name>.spl`.
2. Add `<name>.esql` and/or `<name>.promql` with the exact expected output.
3. If the case warns or is refused, add `<name>.json` with the expected
   diagnostics or error.
4. Run `go test ./conformance`.
5. If a golden was hand-written and is wrong, regenerate with
   `make conformance-update` and review the diff.

The conformance suite compares exact output and exact diagnostics, including
positions. Do not loosen a golden to make a change pass.

## Adding a target

1. Create `pkg/emitter/<target>` exposing:

   ```go
   func Emit(p *ast.Pipeline, opts Options) (*Result, error)
   ```

2. Return `Result{Query, Diagnostics}` and an error when any diagnostic is
   error-severity.
3. Register the target in `pkg/compiler` and add a badge/row to the README.
4. Add golden tests and a compatibility section.

Do not parse SPL in the emitter. If the IR lacks something, add an IR node.

## Code style

| Area | Convention |
|---|---|
| Formatting | `gofmt`; `make fmt-check` must pass |
| Errors | `diag.NewError` with a stable code and source position |
| Warnings | `warnf` with a code that appears in `docs/COMPATIBILITY.md` |
| Naming | Short, package-local; exported API documented |
| Dependencies | Standard library plus the pinned ANTLR runtime; any further exception needs an ADR |
| Tests | Table-driven; assert codes and exact output |

## Pull request checklist

- [ ] `make check` and `make generate-check` pass
- [ ] New behavior has tests and, when relevant, a conformance case
- [ ] Diagnostic codes are documented
- [ ] No new third-party dependencies beyond the ANTLR runtime ([ADR 0004](docs/adr/0004-antlr-parsers.md))
- [ ] Commit messages follow the hard rule
- [ ] The PR description explains the semantic mapping and its risks

## Reporting bugs

Include:

```text
OpenSPL version:
Target and engine version:
SPL query (redacted):
Mapping (redacted):
Expected output:
Actual output / diagnostics:
```

Sensitive queries can be reduced to the smallest failing shape. Do not post
production credentials, customer identifiers, or private hostnames.

## Security

See [SECURITY.md](SECURITY.md). Do not open a public issue for a vulnerability.
