# Security policy

## Supported versions

| Version | Supported |
|---|---|
| `0.1.x` | ✅ Security fixes |
| `< 0.1` | ❌ |

## Reporting a vulnerability

Please use GitHub's private vulnerability reporting:

1. Go to the repository's **Security** tab.
2. Choose **Report a vulnerability**.
3. Include a description, reproduction steps, affected version, and impact.

If private reporting is unavailable, open a minimal issue asking for a private
channel without disclosing details.

Do not open a public issue containing an exploit, credentials, customer data, or
production query text.

## Response expectations

| Step | Target |
|---|---|
| Acknowledgement | 3 business days |
| Triage and severity | 7 business days |
| Fix or mitigation | Depends on severity; coordinated disclosure |

## Threat model

OpenSPL is an offline compiler. It parses text and prints text.

| Property | Status |
|---|---|
| Network access | None in the translation path |
| Telemetry | None |
| Accounts or credentials | None |
| Query persistence | None; input is processed in memory |
| Third-party Go dependencies | One: the ANTLR Go parser runtime and its transitive `golang.org/x/exp` helper, both pinned in `go.mod` and recorded in [ADR 0004](docs/adr/0004-antlr-parsers.md) |
| Build-time tools | ANTLR 4.13.2 and Java 17 are needed only to regenerate the committed parser code, never to build or run the binary |
| Shell execution | None |
| File writes | Only the output the caller redirects |

The primary risks are therefore:

1. **Parsing denial of service** — pathological input causing excessive CPU or
   memory. Reports are welcome; lexing is single-pass and bounded by input size,
   but nested expressions and ANTLR's adaptive prediction are recursive. The
   generated parsers have fuzz targets (`FuzzParseSPL1`, `FuzzParseSPL2`); their
   seed corpora run with the normal test suite, and every accepted failure must
   be a positioned diagnostic, never a panic.
2. **Diagnostic injection** — untrusted query text containing control characters
   that corrupt terminal output. Messages preserve source values; avoid piping
   diagnostics from untrusted sources into privileged logs without sanitizing.
3. **Misleading translation** — a warning that is ignored. Use `--strict` in CI
   and review `rex-lossy`, `fulltext-approximation`, and `default-span` warnings.
4. **Supply chain** — the runtime graph is two pinned modules. Releases embed
   `go.sum`; verify releases against their published checksums when available.
   Generated parser sources are committed, so a release build compiles the same
   Go files a reviewer can see, with no code generation step.

## Secrets in queries

The tool does not transmit queries, but shell history and CI logs can leak
them. Prefer `--file` or stdin over inline arguments for sensitive searches:

```bash
openspl transpile --file savedsearch.spl
```

Redact index names and hostnames before sharing a reproduction.

## Safe deployment guidance

| Practice | Reason |
|---|---|
| Pin a release and verify its checksum | Reproducible, auditable binary |
| Run in CI with `--strict` | Approximations fail the build |
| Do not feed untrusted queries to a privileged shell | Standard shell hygiene |
| Review emitted ES\|QL/PromQL before production | The compiler is alpha |
| Keep mappings in version control | Field drift becomes reviewable |
