package diag

import (
	"errors"
	"reflect"
	"testing"

	"github.com/be-thomas/openspl/pkg/ast"
)

func at(line, col int, severity Severity, code, message string) Diagnostic {
	return Diagnostic{
		Severity: severity,
		Code:     code,
		Message:  message,
		Pos:      ast.Pos{Line: line, Col: col, Offset: line*100 + col},
	}
}

func TestMergeDeduplicatesExactDiagnostics(t *testing.T) {
	warning := at(1, 1, SeverityWarning, "fulltext-approximation", "approximate")
	sameWarning := warning
	sameWarning.Pos.Offset = 999 // position fields not in the dedupe key are ignored
	otherCode := at(1, 1, SeverityWarning, "late-index-filter", "late")
	otherSeverity := at(1, 1, SeverityError, "fulltext-approximation", "approximate")

	got := Merge([]Diagnostic{warning}, []Diagnostic{sameWarning, otherCode}, []Diagnostic{otherSeverity})

	if len(got) != 3 {
		t.Fatalf("Merge returned %d diagnostics, want 3: %#v", len(got), got)
	}
	// Sort puts the error before the warnings at the same position.
	want := []Diagnostic{otherSeverity, warning, otherCode}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Merge = %#v, want %#v", got, want)
	}
	// The surviving copy is the first one seen, so its offset is untouched.
	if got[1].Pos.Offset != warning.Pos.Offset {
		t.Fatalf("dedupe kept offset %d, want first copy's %d", got[1].Pos.Offset, warning.Pos.Offset)
	}
}

func TestMergeWithNoLists(t *testing.T) {
	if got := Merge(); got != nil {
		t.Fatalf("Merge() = %#v, want nil", got)
	}
}

func TestSortOrdersByPositionThenSeverityThenCode(t *testing.T) {
	ds := []Diagnostic{
		at(2, 1, SeverityWarning, "a", "line two"),
		at(1, 5, SeverityWarning, "a", "later column"),
		at(1, 2, SeverityWarning, "b", "same position, code b"),
		at(1, 2, SeverityWarning, "a", "same position, code a"),
		at(1, 2, SeverityError, "z", "same position, error"),
		at(1, 2, SeverityInfo, "z", "same position, info"),
	}

	Sort(ds)

	want := []string{
		"same position, error",
		"same position, info",
		"same position, code a",
		"same position, code b",
		"later column",
		"line two",
	}
	got := make([]string, len(ds))
	for i, d := range ds {
		got[i] = d.Message
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Sort order = %q, want %q", got, want)
	}
}

func TestRenderSortsAndJoins(t *testing.T) {
	ds := []Diagnostic{
		at(2, 3, SeverityError, "syntax", "second"),
		at(1, 1, SeverityWarning, "unsupported-agg", "first"),
	}

	got := Render(ds)
	want := "1:1: warning[unsupported-agg]: first\n2:3: error[syntax]: second"
	if got != want {
		t.Fatalf("Render = %q, want %q", got, want)
	}
	// Render sorts its input in place, and callers rely on that for JSON output.
	if ds[0].Message != "first" {
		t.Fatalf("Render left the slice unsorted: %#v", ds)
	}
}

func TestErrorPrefersTheFirstErrorSeverity(t *testing.T) {
	tests := []struct {
		name string
		err  *Error
		want string
	}{
		{
			name: "skips leading warning",
			err: &Error{Diagnostics: []Diagnostic{
				at(1, 1, SeverityWarning, "late-index-filter", "warned"),
				at(2, 2, SeverityError, "syntax", "failed"),
			}},
			want: "2:2: error[syntax]: failed",
		},
		{
			name: "warning only falls back to the first diagnostic",
			err: &Error{Diagnostics: []Diagnostic{
				at(1, 1, SeverityWarning, "late-index-filter", "warned"),
			}},
			want: "1:1: warning[late-index-filter]: warned",
		},
		{
			name: "empty",
			err:  &Error{},
			want: "compilation failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Fatalf("Error() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestHasErrors(t *testing.T) {
	if HasErrors(nil) {
		t.Fatal("HasErrors(nil) = true, want false")
	}
	if HasErrors([]Diagnostic{at(1, 1, SeverityWarning, "w", "w")}) {
		t.Fatal("HasErrors(warnings) = true, want false")
	}
	if !HasErrors([]Diagnostic{at(1, 1, SeverityError, "e", "e")}) {
		t.Fatal("HasErrors(error) = false, want true")
	}
}

func TestFromError(t *testing.T) {
	if ds, ok := FromError(nil); ok || ds != nil {
		t.Fatalf("FromError(nil) = %#v, %v, want nil, false", ds, ok)
	}
	if ds, ok := FromError(errors.New("plain")); ok || ds != nil {
		t.Fatalf("FromError(plain error) = %#v, %v, want nil, false", ds, ok)
	}

	inner := at(3, 4, SeverityError, "syntax", "bad")
	ds, ok := FromError(&Error{Diagnostics: []Diagnostic{inner}})
	if !ok {
		t.Fatal("FromError(*Error) did not report ok")
	}
	if !reflect.DeepEqual(ds, []Diagnostic{inner}) {
		t.Fatalf("FromError(*Error) = %#v, want %#v", ds, []Diagnostic{inner})
	}
}
