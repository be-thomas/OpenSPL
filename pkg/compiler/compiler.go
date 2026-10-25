// Package compiler wires front-ends and back-ends together.
package compiler

import (
	"fmt"
	"strings"

	"github.com/be-thomas/openspl/pkg/ast"
	"github.com/be-thomas/openspl/pkg/diag"
	"github.com/be-thomas/openspl/pkg/emitter/esql"
	"github.com/be-thomas/openspl/pkg/emitter/promql"
	"github.com/be-thomas/openspl/pkg/mapping"
	"github.com/be-thomas/openspl/pkg/parser/spl1"
)

// Target names a supported back-end.
type Target string

const (
	// ESQL targets Elastic ES|QL.
	ESQL Target = "esql"
	// PromQL targets Prometheus PromQL.
	PromQL Target = "promql"
)

var allTargets = []Target{ESQL, PromQL}

// ParseTarget resolves a user-facing target name.
func ParseTarget(s string) (Target, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "esql", "elastic", "elasticsearch":
		return ESQL, nil
	case "promql", "prometheus", "prom":
		return PromQL, nil
	default:
		return "", fmt.Errorf("unknown target %q; supported targets: esql, promql", s)
	}
}

// Targets returns the supported targets in display order.
func Targets() []Target { return append([]Target(nil), allTargets...) }

// Options configures a translation.
type Options struct {
	Index   string
	Metric  string
	Span    string
	Mapping *mapping.Config
}

// Result is a translation result.
type Result struct {
	Target      Target            `json:"target"`
	Output      string            `json:"query,omitempty"`
	Diagnostics []diag.Diagnostic `json:"diagnostics,omitempty"`
}

// Compile parses and translates a query in one call.
func Compile(src string, target Target, opts Options) (*Result, *ast.Pipeline, error) {
	p, err := spl1.Parse(src)
	if err != nil {
		return nil, nil, err
	}
	res, err := Transpile(p, target, opts)
	return res, p, err
}

// Transpile translates an already-parsed pipeline.
func Transpile(p *ast.Pipeline, target Target, opts Options) (*Result, error) {
	switch target {
	case ESQL:
		r, err := esql.Emit(p, esql.Options{Index: opts.Index, Span: opts.Span, Mapping: opts.Mapping})
		res := &Result{Target: target}
		if r != nil {
			res.Output = r.Query
			res.Diagnostics = r.Diagnostics
		}
		if err != nil {
			res.Diagnostics = diagnosticsFrom(res.Diagnostics, err)
			return res, err
		}
		return res, nil

	case PromQL:
		r, err := promql.Emit(p, promql.Options{Metric: opts.Metric, Span: opts.Span, Mapping: opts.Mapping})
		res := &Result{Target: target}
		if r != nil {
			res.Output = r.Query
			res.Diagnostics = r.Diagnostics
		}
		if err != nil {
			res.Diagnostics = diagnosticsFrom(res.Diagnostics, err)
			return res, err
		}
		return res, nil

	default:
		return nil, diag.NewError(ast.Pos{Line: 1, Col: 1}, "unknown-target", "unknown target %q", target)
	}
}

// diagnosticsFrom merges diagnostics already accumulated by earlier stages
// with the error that stopped a later one. Replacing the list would drop
// warnings that the user still needs to see.
func diagnosticsFrom(existing []diag.Diagnostic, err error) []diag.Diagnostic {
	if ds, ok := diag.FromError(err); ok {
		return diag.Merge(existing, ds)
	}
	if len(existing) > 0 {
		return existing
	}
	return []diag.Diagnostic{{
		Severity: diag.SeverityError,
		Code:     "internal",
		Message:  err.Error(),
		Pos:      ast.Pos{Line: 1, Col: 1},
	}}
}
