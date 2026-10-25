// Package spl2 is the reserved front-end for Splunk's SPL2 dialect.
//
// SPL2 differs from classic SPL in both surface syntax (`from ... | where ...`
// pipelines, typed expressions, `eval` without pipes in some contexts) and
// semantics. Keeping it a separate package prevents classic-SPL quirks from
// leaking into the newer grammar while both lower into pkg/ast.
//
// The front-end is intentionally not implemented in v0.1.0. Returning a
// positioned, explicit error is better than silently accepting SPL2 with the
// classic grammar and translating it incorrectly.
package spl2

import (
	"github.com/be-thomas/openspl/pkg/ast"
	"github.com/be-thomas/openspl/pkg/diag"
)

// Parse parses an SPL2 query into the unified IR.
func Parse(src string) (*ast.Pipeline, error) {
	return nil, diag.NewError(ast.Pos{Line: 1, Col: 1}, "spl2-not-implemented",
		"the SPL2 front-end is not implemented yet; see docs/ROADMAP.md")
}
