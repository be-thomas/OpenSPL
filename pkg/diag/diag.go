// Package diag carries compiler diagnostics with source positions.
//
// Deterministic, positioned errors are the product: a transpiler that cannot
// say where and why it refused a query is not trustworthy enough to migrate
// production alerts.
package diag

import (
	"fmt"
	"sort"
	"strings"

	"github.com/be-thomas/openspl/pkg/ast"
)

// Severity is how seriously a diagnostic affects the result.
type Severity string

const (
	SeverityError   Severity = "error"
	SeverityWarning Severity = "warning"
	SeverityInfo    Severity = "info"
)

// Diagnostic is a single positioned message.
type Diagnostic struct {
	Severity Severity `json:"severity"`
	Code     string   `json:"code"`
	Message  string   `json:"message"`
	Pos      ast.Pos  `json:"pos"`
}

// String renders a diagnostic in the CLI form: line:col: severity[code]: message.
func (d Diagnostic) String() string {
	return fmt.Sprintf("%d:%d: %s[%s]: %s", d.Pos.Line, d.Pos.Col, d.Severity, d.Code, d.Message)
}

// Error is an error that carries one or more diagnostics.
type Error struct {
	Diagnostics []Diagnostic
}

func (e *Error) Error() string {
	for _, d := range e.Diagnostics {
		if d.Severity == SeverityError {
			return d.String()
		}
	}
	if len(e.Diagnostics) > 0 {
		return e.Diagnostics[0].String()
	}
	return "compilation failed"
}

// NewError builds an Error with a single error-severity diagnostic.
func NewError(pos ast.Pos, code, format string, args ...any) *Error {
	return &Error{Diagnostics: []Diagnostic{{
		Severity: SeverityError,
		Code:     code,
		Message:  fmt.Sprintf(format, args...),
		Pos:      pos,
	}}}
}

// HasErrors reports whether any diagnostic is error-severity.
func HasErrors(ds []Diagnostic) bool {
	for _, d := range ds {
		if d.Severity == SeverityError {
			return true
		}
	}
	return false
}

// Sort orders diagnostics by position, then severity, then code.
func Sort(ds []Diagnostic) {
	sort.SliceStable(ds, func(i, j int) bool {
		a, b := ds[i], ds[j]
		if a.Pos.Line != b.Pos.Line {
			return a.Pos.Line < b.Pos.Line
		}
		if a.Pos.Col != b.Pos.Col {
			return a.Pos.Col < b.Pos.Col
		}
		if a.Severity != b.Severity {
			return a.Severity < b.Severity
		}
		return a.Code < b.Code
	})
}

// Merge combines diagnostic lists, removes exact duplicates, and sorts the
// result. It is used when a later stage fails after earlier stages have
// already recorded warnings: the warning must survive the error.
func Merge(lists ...[]Diagnostic) []Diagnostic {
	type key struct {
		severity Severity
		code     string
		line     int
		col      int
		message  string
	}
	seen := map[key]bool{}
	var out []Diagnostic
	for _, ds := range lists {
		for _, d := range ds {
			k := key{d.Severity, d.Code, d.Pos.Line, d.Pos.Col, d.Message}
			if seen[k] {
				continue
			}
			seen[k] = true
			out = append(out, d)
		}
	}
	Sort(out)
	return out
}

// FromError extracts diagnostics from err when it is a *Error.
func FromError(err error) ([]Diagnostic, bool) {
	if err == nil {
		return nil, false
	}
	if de, ok := err.(*Error); ok {
		return de.Diagnostics, true
	}
	return nil, false
}

// Render produces a stable, human-readable block for stderr.
func Render(ds []Diagnostic) string {
	Sort(ds)
	lines := make([]string, 0, len(ds))
	for _, d := range ds {
		lines = append(lines, d.String())
	}
	return strings.Join(lines, "\n")
}
