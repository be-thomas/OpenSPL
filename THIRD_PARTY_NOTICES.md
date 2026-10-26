# Third-party notices

OpenSPL is distributed under the license in [LICENSE](LICENSE). It links one
third-party runtime library and uses one third-party build tool. Both are
listed here with their versions and licenses.

## ANTLR 4 Go runtime

| Field | Value |
|---|---|
| Module | `github.com/antlr4-go/antlr/v4` |
| Version | `v4.13.1` |
| License | BSD 3-Clause |
| Upstream | <https://github.com/antlr/antlr4-go> |
| Role | Runtime support for the generated SPL1 and SPL2 parsers. Linked into every OpenSPL binary. |

## golang.org/x/exp

| Field | Value |
|---|---|
| Module | `golang.org/x/exp` |
| Version | `v0.0.0-20240506185415-9bf2ced13842` |
| License | BSD 3-Clause |
| Upstream | <https://pkg.go.dev/golang.org/x/exp> |
| Role | Indirect dependency of the ANTLR Go runtime. |

## ANTLR tool

| Field | Value |
|---|---|
| Artifact | `org.antlr:antlr4:4.13.2` complete jar |
| SHA-256 | `eae2dfa119a64327444672aff63e9ec35a20180dc5b8090b7a6ab85125df4d76` |
| License | BSD 3-Clause |
| Upstream | <https://www.antlr.org/> |
| Role | Build-time only. Generates the Go parser sources committed under `pkg/parser/*/gen`. Never shipped, and not needed to build OpenSPL. |

## Splunk documentation examples

| Field | Value |
|---|---|
| Source | Splunk Enterprise search reference and Splunk Cloud Platform SPL2 search reference |
| Upstream | <https://help.splunk.com/en/splunk-enterprise/search/spl-search-reference>, <https://help.splunk.com/en/splunk-cloud-platform/search/spl2-search-reference> |
| Role | Test data only. `conformance/corpus` reproduces the documentation's short example queries as parser-coverage inputs. No prose is included, nothing is linked into the binary, and the gatherer records each query's page and documentation versions. |
| Terms | The example queries remain Splunk's material and are not covered by OpenSPL's license. |

## Generated code

The Go files under `pkg/parser/spl1/gen` and `pkg/parser/spl2/gen` are produced
from the grammars in this repository by the ANTLR tool. They are covered by the
BSD 3-Clause license above because they embed portions of ANTLR's generated
runtime scaffolding; the grammar itself is part of OpenSPL.

See [ADR 0004](docs/adr/0004-antlr-parsers.md) for the version pairing and the
upgrade procedure.
