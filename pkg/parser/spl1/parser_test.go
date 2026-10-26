package spl1

import (
	"strings"
	"testing"

	"github.com/be-thomas/openspl/pkg/ast"
	"github.com/be-thomas/openspl/pkg/diag"
)

func mustParse(t *testing.T, src string) *ast.Pipeline {
	t.Helper()
	p, err := Parse(src)
	if err != nil {
		t.Fatalf("Parse(%q) failed: %v", src, err)
	}
	return p
}

func TestParseStageKinds(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want []string
	}{
		{"implicit search", `index=web status=500`, []string{"search"}},
		{"stats", `index=web | stats count by host`, []string{"search", "stats"}},
		{"timechart", `index=web | timechart span=5m count by host`, []string{"search", "timechart"}},
		{"eval", `index=web | eval kb=bytes/1024, who=lower(host)`, []string{"search", "eval"}},
		{"rex", `index=web | rex field=_raw "(?<user>\w+)"`, []string{"search", "rex"}},
		{"table", `index=web | table host, source`, []string{"search", "table"}},
		{"fields include", `index=web | fields host source`, []string{"search", "fields"}},
		{"fields exclude", `index=web | fields - _raw`, []string{"search", "fields"}},
		{"sort and head", `index=web | sort limit=10 -bytes | head 5`, []string{"search", "sort", "head"}},
		{"bin", `index=web | bin span=1h _time`, []string{"search", "bin"}},
		{"leading pipe", `| stats count`, []string{"stats"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := mustParse(t, tc.src)
			if len(p.Stages) != len(tc.want) {
				t.Fatalf("got %d stages, want %d", len(p.Stages), len(tc.want))
			}
			for i, want := range tc.want {
				if got := p.Stages[i].StageName(); got != want {
					t.Errorf("stage %d: got %q, want %q", i, got, want)
				}
			}
		})
	}
}

func TestSearchValueSemantics(t *testing.T) {
	t.Run("search value is a literal", func(t *testing.T) {
		p := mustParse(t, `search host=web`)
		b := p.Stages[0].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
		s, ok := b.Right.(*ast.StringLit)
		if !ok || s.Value != "web" {
			t.Fatalf("right side = %#v, want StringLit web", b.Right)
		}
	})

	t.Run("where value is a field reference", func(t *testing.T) {
		p := mustParse(t, `index=web | where host=web`)
		b := p.Stages[1].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
		f, ok := b.Right.(*ast.FieldRef)
		if !ok || f.Name != "web" {
			t.Fatalf("right side = %#v, want FieldRef web", b.Right)
		}
	})

	t.Run("quoted comparison keeps a string", func(t *testing.T) {
		p := mustParse(t, `index=web status="500"`)
		and := p.Stages[0].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
		b := and.Right.(*ast.BinaryExpr)
		if _, ok := b.Right.(*ast.StringLit); !ok {
			t.Fatalf("right side = %#v, want StringLit", b.Right)
		}
	})
}

func TestHyphenatedSearchValues(t *testing.T) {
	p := mustParse(t, `index=logs-2024-01 host=web-01 | stats count`)
	and := p.Stages[0].(*ast.SearchStage).Expr.(*ast.BinaryExpr)

	left := and.Left.(*ast.BinaryExpr)
	if s, ok := left.Right.(*ast.StringLit); !ok || s.Value != "logs-2024-01" {
		t.Fatalf("index value = %#v, want StringLit logs-2024-01", left.Right)
	}

	right := and.Right.(*ast.BinaryExpr)
	if s, ok := right.Right.(*ast.StringLit); !ok || s.Value != "web-01" {
		t.Fatalf("host value = %#v, want StringLit web-01", right.Right)
	}
}

func TestWildcards(t *testing.T) {
	tests := []struct {
		src  string
		want string
	}{
		{`index=web host=web*`, "web*"},
		{`index=web host=*web`, "*web"},
		{`index=web host=*`, "*"},
	}
	for _, tc := range tests {
		p := mustParse(t, tc.src)
		// The index filter is the left side of the implicit AND.
		and := p.Stages[0].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
		b := and.Right.(*ast.BinaryExpr)
		switch v := b.Right.(type) {
		case *ast.StringLit:
			if v.Value != tc.want {
				t.Errorf("%s: got %q, want %q", tc.src, v.Value, tc.want)
			}
		case *ast.WildcardLit:
			if tc.want != "*" {
				t.Errorf("%s: got wildcard, want %q", tc.src, tc.want)
			}
		default:
			t.Errorf("%s: right side = %#v", tc.src, b.Right)
		}
	}
}

func TestSearchTerms(t *testing.T) {
	p := mustParse(t, `error timeout`)
	b, ok := p.Stages[0].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
	if !ok || b.Op != "and" {
		t.Fatalf("expr = %#v, want implicit AND", p.Stages[0].(*ast.SearchStage).Expr)
	}
	left := b.Left.(*ast.SearchTerm)
	right := b.Right.(*ast.SearchTerm)
	if left.Value != "error" || right.Value != "timeout" {
		t.Fatalf("terms = %q, %q", left.Value, right.Value)
	}
}

func TestQuotedSearchPhrase(t *testing.T) {
	p := mustParse(t, `"failed password" host=web`)
	b := p.Stages[0].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
	term, ok := b.Left.(*ast.SearchTerm)
	if !ok || term.Value != "failed password" {
		t.Fatalf("left side = %#v, want SearchTerm \"failed password\"", b.Left)
	}
	if _, ok := b.Right.(*ast.BinaryExpr); !ok {
		t.Fatalf("right side = %#v, want the host filter", b.Right)
	}
}

func TestWildcardValues(t *testing.T) {
	tests := []struct {
		src  string
		want string
	}{
		{`index=web host=web-*`, "web-*"},
		{`index=web host=*error*`, "*error*"},
		{`index=web host=web*prod`, "web*prod"},
	}
	for _, tc := range tests {
		p := mustParse(t, tc.src)
		and := p.Stages[0].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
		b := and.Right.(*ast.BinaryExpr)
		s, ok := b.Right.(*ast.StringLit)
		if !ok || s.Value != tc.want {
			t.Errorf("%s: right side = %#v, want StringLit %q", tc.src, b.Right, tc.want)
		}
	}
}

func TestSearchTermWildcardRun(t *testing.T) {
	p := mustParse(t, `*error*`)
	term, ok := p.Stages[0].(*ast.SearchStage).Expr.(*ast.SearchTerm)
	if !ok || term.Value != "*error*" {
		t.Fatalf("expr = %#v, want SearchTerm *error*", p.Stages[0].(*ast.SearchStage).Expr)
	}
}

// TestWildcardTermBeforeStage covers a wildcard that ends a search term right
// before the pipe. It must stay part of the term, not become a multiplication
// whose right operand is missing.
func TestWildcardTermBeforeStage(t *testing.T) {
	tests := []struct {
		src   string
		terms []string
	}{
		{`a* | stats count`, []string{"a*"}},
		{`a** | stats count`, []string{"a**"}},
		{`a*1 | stats count`, []string{"a*1"}},
		{`a*b | stats count`, []string{"a*b"}},
		{`index=web error* | stats count`, []string{"error*"}},
		{`search a* | stats count`, []string{"a*"}},
		{`*error* | stats count`, []string{"*error*"}},
		{`NOT a* | stats count`, []string{"a*"}},
		{`a* OR b* | stats count`, []string{"a*", "b*"}},
		{`web*prod | stats count`, []string{"web*prod"}},
	}
	for _, tc := range tests {
		t.Run(tc.src, func(t *testing.T) {
			p := mustParse(t, tc.src)
			if len(p.Stages) != 2 || p.Stages[1].StageName() != "stats" {
				t.Fatalf("stages = %v, want search then stats", p.Stages)
			}
			expr := p.Stages[0].(*ast.SearchStage).Expr
			for _, want := range tc.terms {
				if !exprHasTerm(expr, want) {
					t.Errorf("expression %#v has no term %q", expr, want)
				}
			}
		})
	}
}

// TestAdjacentStarStaysMultiplication checks the other side of the boundary:
// a star that does not touch the preceding word is still the multiplication
// operator, as are stars after numbers and durations.
func TestAdjacentStarStaysMultiplication(t *testing.T) {
	tests := []string{`a * b | stats count`, `1*2 | stats count`, `5m*2 | stats count`}
	for _, src := range tests {
		t.Run(src, func(t *testing.T) {
			p := mustParse(t, src)
			b, ok := p.Stages[0].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
			if !ok || b.Op != "*" {
				t.Fatalf("expr = %#v, want multiplication", p.Stages[0].(*ast.SearchStage).Expr)
			}
		})
	}
}

// TestNotAfterOperatorIsAField covers the NOT_WORD retyping: after an
// arithmetic or comparison operator `not` is a field name, so it does not
// start the negation operator.
func TestNotAfterOperatorIsAField(t *testing.T) {
	t.Run("arithmetic right operand", func(t *testing.T) {
		p := mustParse(t, `index=web | eval x=1+not`)
		e := p.Stages[1].(*ast.EvalStage)
		b, ok := e.Assignments[0].Expr.(*ast.BinaryExpr)
		if !ok || b.Op != "+" {
			t.Fatalf("expr = %#v, want addition", e.Assignments[0].Expr)
		}
		f, ok := b.Right.(*ast.FieldRef)
		if !ok || f.Name != "not" {
			t.Fatalf("right side = %#v, want FieldRef not", b.Right)
		}
	})

	t.Run("comparison right operand", func(t *testing.T) {
		p := mustParse(t, `index=web | where a=not`)
		if !exprHasField(p.Stages[1].(*ast.SearchStage).Expr, "not") {
			t.Fatalf("expr lost the not field: %#v", p.Stages[1].(*ast.SearchStage).Expr)
		}
	})

	t.Run("group key", func(t *testing.T) {
		p := mustParse(t, `index=web | stats count by not`)
		s := p.Stages[1].(*ast.StatsStage)
		if len(s.By) != 1 {
			t.Fatalf("by = %#v, want one key", s.By)
		}
		if f, ok := s.By[0].(*ast.FieldRef); !ok || f.Name != "not" {
			t.Fatalf("by[0] = %#v, want FieldRef not", s.By[0])
		}
	})

	t.Run("assignment target", func(t *testing.T) {
		p := mustParse(t, `index=web | eval not=1`)
		e := p.Stages[1].(*ast.EvalStage)
		if e.Assignments[0].Name != "not" {
			t.Fatalf("assignment = %#v, want target not", e.Assignments[0])
		}
	})
}

// TestNotAtExpressionStartIsNegation is the other side of the retyping: at the
// start of an expression `not` is the negation operator, including after an
// `eval` assignment's `=`, and a call with several arguments is then rejected
// at the argument comma instead of becoming a function named `not`.
func TestNotAtExpressionStartIsNegation(t *testing.T) {
	tests := []struct {
		src  string
		line int
		col  int
	}{
		{`not%x`, 1, 4},
		{`index=web | eval y=not(x,y)`, 1, 25},
		{`eval y=not(x,y)`, 1, 13},
		{`index=web | eval x=not`, 1, 23},
	}
	for _, tc := range tests {
		t.Run(tc.src, func(t *testing.T) {
			_, err := Parse(tc.src)
			if err == nil {
				t.Fatal("expected an error")
			}
			ds, ok := diag.FromError(err)
			if !ok || len(ds) != 1 {
				t.Fatalf("err = %v, want one diagnostic", err)
			}
			d := ds[0]
			if d.Code != "syntax" || d.Pos.Line != tc.line || d.Pos.Col != tc.col {
				t.Fatalf("diagnostic = %#v, want syntax at %d:%d", d, tc.line, tc.col)
			}
		})
	}

	t.Run("a lone assignment value negates", func(t *testing.T) {
		p := mustParse(t, `eval x=not(2)`)
		e := p.Stages[0].(*ast.EvalStage)
		u, ok := e.Assignments[0].Expr.(*ast.UnaryExpr)
		if !ok || u.Op != "not" {
			t.Fatalf("expr = %#v, want NOT 2", e.Assignments[0].Expr)
		}
	})

	t.Run("assignment value negates a wildcard", func(t *testing.T) {
		p := mustParse(t, `eval y=not * b`)
		e := p.Stages[0].(*ast.EvalStage)
		and, ok := e.Assignments[0].Expr.(*ast.BinaryExpr)
		if !ok || and.Op != "and" {
			t.Fatalf("expr = %#v, want NOT TRUE AND b", e.Assignments[0].Expr)
		}
		u, ok := and.Left.(*ast.UnaryExpr)
		if !ok || u.Op != "not" {
			t.Fatalf("left = %#v, want a negation", and.Left)
		}
		if _, ok := u.Operand.(*ast.WildcardLit); !ok {
			t.Fatalf("operand = %#v, want the wildcard", u.Operand)
		}
	})

	t.Run("leading eval keeps a field after an operator", func(t *testing.T) {
		p := mustParse(t, `eval x=1+not`)
		e := p.Stages[0].(*ast.EvalStage)
		b, ok := e.Assignments[0].Expr.(*ast.BinaryExpr)
		if !ok || b.Op != "+" {
			t.Fatalf("expr = %#v, want addition", e.Assignments[0].Expr)
		}
		f, ok := b.Right.(*ast.FieldRef)
		if !ok || f.Name != "not" {
			t.Fatalf("right side = %#v, want FieldRef not", b.Right)
		}
	})
}

func TestCallRequiresAdjacency(t *testing.T) {
	t.Run("grouped term after a value is not a call", func(t *testing.T) {
		p := mustParse(t, `index=web (status=500)`)
		and := p.Stages[0].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
		index, ok := and.Left.(*ast.BinaryExpr)
		if !ok || index.Op != "=" {
			t.Fatalf("left side = %#v, want the index filter", and.Left)
		}
		right, ok := and.Right.(*ast.BinaryExpr)
		if !ok || right.Op != "=" {
			t.Fatalf("right side = %#v, want the grouped status filter", and.Right)
		}
	})

	t.Run("free-text term survives before a group", func(t *testing.T) {
		p := mustParse(t, `index=web error (status=500)`)
		expr := p.Stages[0].(*ast.SearchStage).Expr
		if !exprHasTerm(expr, "error") {
			t.Fatalf("expression lost the free-text term: %#v", expr)
		}
		if exprHasField(expr, "web") {
			t.Fatalf("index value was parsed as a function: %#v", expr)
		}
	})

	t.Run("time bound still lifts before a group", func(t *testing.T) {
		p := mustParse(t, `index=web earliest=now (status=500)`)
		s := p.Stages[0].(*ast.SearchStage)
		if s.Earliest != "now" {
			t.Fatalf("earliest = %q, want now", s.Earliest)
		}
		if exprHasField(s.Expr, "earliest") {
			t.Fatalf("earliest stayed in the row filter: %#v", s.Expr)
		}
	})

	t.Run("adjacent parenthesis is a call", func(t *testing.T) {
		p := mustParse(t, `index=web | eval x=lower(msg)`)
		e := p.Stages[1].(*ast.EvalStage)
		if _, ok := e.Assignments[0].Expr.(*ast.CallExpr); !ok {
			t.Fatalf("expr = %#v, want a call", e.Assignments[0].Expr)
		}
	})

	t.Run("field expressions keep the permissive call form", func(t *testing.T) {
		p := mustParse(t, `index=web | eval x=lower (msg)`)
		e := p.Stages[1].(*ast.EvalStage)
		if _, ok := e.Assignments[0].Expr.(*ast.CallExpr); !ok {
			t.Fatalf("expr = %#v, want a call", e.Assignments[0].Expr)
		}
	})

	t.Run("no-space parenthesis is not a search call", func(t *testing.T) {
		p := mustParse(t, `index=web(status=500)`)
		and := p.Stages[0].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
		if and.Op != "and" || !exprHasField(and.Left, "index") || !exprHasField(and.Right, "status") {
			t.Fatalf("expr = %#v, want index AND status", and)
		}
	})

	t.Run("no-space parenthesis keeps a free-text term", func(t *testing.T) {
		p := mustParse(t, `index=web error(status=500)`)
		expr := p.Stages[0].(*ast.SearchStage).Expr
		if !exprHasTerm(expr, "error") || !exprHasField(expr, "status") {
			t.Fatalf("expr = %#v, want error AND status", expr)
		}
		if exprHasField(expr, "web") {
			t.Fatalf("index value was parsed as a function: %#v", expr)
		}
	})
}

func exprHasTerm(e ast.Expr, value string) bool {
	switch v := e.(type) {
	case *ast.SearchTerm:
		return v.Value == value
	case *ast.BinaryExpr:
		return exprHasTerm(v.Left, value) || exprHasTerm(v.Right, value)
	case *ast.UnaryExpr:
		return exprHasTerm(v.Operand, value)
	}
	return false
}

func TestSearchValueRuns(t *testing.T) {
	tests := []struct {
		src  string
		want string
	}{
		{`index=web host=*-*`, "*-*"},
		{`index=web host=-a`, "-a"},
		{`index=web host=10-20`, "10-20"},
		{`index=web status=200-299`, "200-299"},
		{`index=web status=500*`, "500*"},
	}
	for _, tc := range tests {
		p := mustParse(t, tc.src)
		and := p.Stages[0].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
		b := and.Right.(*ast.BinaryExpr)
		s, ok := b.Right.(*ast.StringLit)
		if !ok || s.Value != tc.want {
			t.Errorf("%s: right side = %#v, want StringLit %q", tc.src, b.Right, tc.want)
		}
	}

	t.Run("a plain number stays numeric", func(t *testing.T) {
		p := mustParse(t, `index=web status=500`)
		and := p.Stages[0].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
		b := and.Right.(*ast.BinaryExpr)
		if _, ok := b.Right.(*ast.NumberLit); !ok {
			t.Fatalf("right side = %#v, want NumberLit", b.Right)
		}
	})

	t.Run("a trailing hyphen is a syntax error", func(t *testing.T) {
		_, err := Parse(`index=web host=a-`)
		if err == nil {
			t.Fatal("expected an error")
		}
		ds, _ := diag.FromError(err)
		if len(ds) == 0 || ds[0].Code != "syntax" || ds[0].Pos.Col != 17 {
			t.Fatalf("diagnostics = %#v, want syntax at 1:17", ds)
		}
	})
}

func TestTimeBoundsUnderOr(t *testing.T) {
	tests := []string{
		`earliest=-24h status=500 OR status=503`,
		`index=web earliest=-24h (status=500 OR status=503)`,
	}
	for _, src := range tests {
		p := mustParse(t, src)
		s := p.Stages[0].(*ast.SearchStage)
		if s.Earliest != "-24h" {
			t.Errorf("%s: earliest = %q, want -24h", src, s.Earliest)
		}
		if exprHasField(s.Expr, "earliest") {
			t.Errorf("%s: earliest stayed in the row filter: %#v", src, s.Expr)
		}
		if !exprHasOp(s.Expr, "or") {
			t.Errorf("%s: expr = %#v, want an OR branch", src, s.Expr)
		}
	}
}

func exprHasOp(e ast.Expr, op string) bool {
	switch v := e.(type) {
	case *ast.BinaryExpr:
		return v.Op == op || exprHasOp(v.Left, op) || exprHasOp(v.Right, op)
	case *ast.UnaryExpr:
		return exprHasOp(v.Operand, op)
	}
	return false
}

func TestTimeBoundsInNonConjunctivePositions(t *testing.T) {
	tests := []struct {
		src       string
		field     string
		wantField bool
	}{
		// A lifted bound is TRUE in the row expression, so TRUE OR X is TRUE.
		{`(earliest=-24h) OR status=500`, "status", false},
		{`status=500 OR earliest=-24h`, "status", false},
		{`status=500 AND (earliest=-24h OR status=503)`, "status", true},
		// NOT of a modifier is dropped; the rest of the term still applies.
		{`NOT earliest=-24h status=500`, "status", true},
		{`NOT (earliest=-24h AND status=500)`, "status", true},
	}
	for _, tc := range tests {
		p := mustParse(t, tc.src)
		s := p.Stages[0].(*ast.SearchStage)
		if s.Earliest != "-24h" {
			t.Errorf("%s: earliest = %q, want -24h", tc.src, s.Earliest)
		}
		if exprHasField(s.Expr, "earliest") {
			t.Errorf("%s: earliest stayed in the row filter: %#v", tc.src, s.Expr)
		}
		if got := exprHasField(s.Expr, tc.field); got != tc.wantField {
			t.Errorf("%s: has %s = %v, want %v (%#v)", tc.src, tc.field, got, tc.wantField, s.Expr)
		}
	}

	not := mustParse(t, `NOT (earliest=-24h AND status=500)`)
	expr := not.Stages[0].(*ast.SearchStage).Expr
	u, ok := expr.(*ast.UnaryExpr)
	if !ok || u.Op != "not" || !exprHasField(u.Operand, "status") {
		t.Fatalf("expr = %#v, want NOT status", expr)
	}
}

func TestConflictingTimeBounds(t *testing.T) {
	_, err := Parse(`(earliest=-1h status=500) OR (earliest=-2h status=503)`)
	if err == nil {
		t.Fatal("expected an error")
	}
	ds, ok := diag.FromError(err)
	if !ok || len(ds) != 1 {
		t.Fatalf("err = %v, want one diagnostic", err)
	}
	if ds[0].Code != "conflicting-time-bounds" || ds[0].Pos.Line != 1 || ds[0].Pos.Col != 39 {
		t.Fatalf("diagnostic = %#v, want conflicting-time-bounds at 1:39", ds[0])
	}

	same := mustParse(t, `(earliest=-1h status=500) OR (earliest=-1h status=503)`)
	if s := same.Stages[0].(*ast.SearchStage); s.Earliest != "-1h" {
		t.Fatalf("earliest = %q, want -1h", s.Earliest)
	}
}

func TestHeadCountSet(t *testing.T) {
	tests := []struct {
		src      string
		count    int
		countSet bool
	}{
		{`index=web | head`, 10, false},
		{`index=web | head 0`, 0, true},
		{`index=web | head 5`, 5, true},
		{`index=web | head limit=5`, 5, true},
	}
	for _, tc := range tests {
		p := mustParse(t, tc.src)
		h := p.Stages[1].(*ast.HeadStage)
		if h.Count != tc.count || h.CountSet != tc.countSet {
			t.Errorf("%s: count=%d countSet=%v, want %d/%v",
				tc.src, h.Count, h.CountSet, tc.count, tc.countSet)
		}
	}
}

func TestTimeBounds(t *testing.T) {
	p := mustParse(t, `index=web earliest=-24h latest=now status=500`)
	s := p.Stages[0].(*ast.SearchStage)
	if s.Earliest != "-24h" {
		t.Errorf("earliest = %q, want -24h", s.Earliest)
	}
	if s.Latest != "now" {
		t.Errorf("latest = %q, want now", s.Latest)
	}
	if s.Expr == nil {
		t.Fatal("status filter was dropped with the time bounds")
	}
	if !exprHasField(s.Expr, "status") {
		t.Fatalf("remaining expr lost the status filter: %#v", s.Expr)
	}
	if exprHasField(s.Expr, "earliest") || exprHasField(s.Expr, "latest") {
		t.Fatalf("time bounds were not lifted out of the expression: %#v", s.Expr)
	}
}

func exprHasField(e ast.Expr, name string) bool {
	switch v := e.(type) {
	case *ast.FieldRef:
		return v.Name == name
	case *ast.BinaryExpr:
		return exprHasField(v.Left, name) || exprHasField(v.Right, name)
	case *ast.UnaryExpr:
		return exprHasField(v.Operand, name)
	case *ast.CallExpr:
		for _, arg := range v.Args {
			if exprHasField(arg, name) {
				return true
			}
		}
	}
	return false
}

func TestPrecedence(t *testing.T) {
	p := mustParse(t, `index=web status=500 OR status=503 level=error`)
	root := p.Stages[0].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
	if root.Op != "or" {
		t.Fatalf("root op = %q, want or", root.Op)
	}
	if root.Right.(*ast.BinaryExpr).Op != "and" {
		t.Fatalf("OR should bind looser than implicit AND: %#v", root.Right)
	}
}

func TestStatsAggregates(t *testing.T) {
	p := mustParse(t, `index=web | stats count, dc(host) AS hosts, avg(bytes) BY sourcetype`)
	s := p.Stages[1].(*ast.StatsStage)
	if len(s.Aggs) != 3 {
		t.Fatalf("got %d aggregations, want 3", len(s.Aggs))
	}
	if s.Aggs[0].Func != "count" || s.Aggs[0].Arg != nil {
		t.Errorf("count agg = %#v", s.Aggs[0])
	}
	if s.Aggs[1].Func != "dc" || s.Aggs[1].Alias != "hosts" {
		t.Errorf("dc agg = %#v", s.Aggs[1])
	}
	if len(s.By) != 1 {
		t.Fatalf("by = %#v", s.By)
	}
}

func TestLexDurationsAndStrings(t *testing.T) {
	toks, err := Lex(`earliest=-24h latest=1mon msg="a \"quoted\" value" re="\d+"`)
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, tok := range toks {
		if tok.Kind == EOF {
			continue
		}
		got = append(got, tok.Text)
	}
	want := []string{"earliest", "=", "-", "24h", "latest", "=", "1mon", "msg", "=", `a "quoted" value`, "re", "=", `\d+`}
	if strings.Join(got, "|") != strings.Join(want, "|") {
		t.Fatalf("tokens = %q, want %q", got, want)
	}
}

func TestParseErrors(t *testing.T) {
	tests := []struct {
		src  string
		code string
	}{
		{"", "empty-query"},
		{"   ", "empty-query"},
		{"index=web |", "syntax"},
		{"index=web | join type=inner a b", "unsupported-command"},
		{"index=web [search index=other]", "unsupported-subsearch"},
		{"index=web | stats", "syntax"},
		{"index=web | eval x=", "syntax"},
		{"index=web | rex field=_raw", "syntax"},
		{"index=web | bin span=1h", "syntax"},
		{"index=web | sort", "syntax"},
	}
	for _, tc := range tests {
		t.Run(tc.src, func(t *testing.T) {
			_, err := Parse(tc.src)
			if err == nil {
				t.Fatal("expected an error")
			}
			ds, ok := diag.FromError(err)
			if !ok || len(ds) == 0 {
				t.Fatalf("error is not a diag.Error: %v", err)
			}
			if ds[0].Code != tc.code {
				t.Errorf("code = %q, want %q (%v)", ds[0].Code, tc.code, err)
			}
			if ds[0].Pos.Line < 1 || ds[0].Pos.Col < 1 {
				t.Errorf("diagnostic has no position: %#v", ds[0])
			}
		})
	}
}

// TestValuePositionStarMultiplies pins the star in value position. The driver
// retypes a star that touches the preceding word as STAR_TERM so `a*` stays one
// search term, but in an expression the same spelling is multiplication.
func TestValuePositionStarMultiplies(t *testing.T) {
	t.Run("eval assignment", func(t *testing.T) {
		p := mustParse(t, `eval x=a*b`)
		e := p.Stages[0].(*ast.EvalStage)
		b, ok := e.Assignments[0].Expr.(*ast.BinaryExpr)
		if !ok || b.Op != "*" {
			t.Fatalf("expr = %#v, want multiplication", e.Assignments[0].Expr)
		}
		if f, ok := b.Left.(*ast.FieldRef); !ok || f.Name != "a" {
			t.Fatalf("left = %#v, want FieldRef a", b.Left)
		}
		if f, ok := b.Right.(*ast.FieldRef); !ok || f.Name != "b" {
			t.Fatalf("right = %#v, want FieldRef b", b.Right)
		}
	})

	t.Run("comparison left operand", func(t *testing.T) {
		p := mustParse(t, `where a*b>2`)
		cmp, ok := p.Stages[0].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
		if !ok || cmp.Op != ">" {
			t.Fatalf("expr = %#v, want a comparison", p.Stages[0].(*ast.SearchStage).Expr)
		}
		if b, ok := cmp.Left.(*ast.BinaryExpr); !ok || b.Op != "*" {
			t.Fatalf("left = %#v, want a product", cmp.Left)
		}
	})

	t.Run("aggregate argument", func(t *testing.T) {
		p := mustParse(t, `stats sum(a*b)`)
		s := p.Stages[0].(*ast.StatsStage)
		if b, ok := s.Aggs[0].Arg.(*ast.BinaryExpr); !ok || b.Op != "*" {
			t.Fatalf("arg = %#v, want a product", s.Aggs[0].Arg)
		}
	})

	t.Run("left associativity", func(t *testing.T) {
		p := mustParse(t, `eval pct=count/total*100`)
		e := p.Stages[0].(*ast.EvalStage)
		b, ok := e.Assignments[0].Expr.(*ast.BinaryExpr)
		if !ok || b.Op != "*" {
			t.Fatalf("expr = %#v, want (count/total)*100", e.Assignments[0].Expr)
		}
		if div, ok := b.Left.(*ast.BinaryExpr); !ok || div.Op != "/" {
			t.Fatalf("left = %#v, want a division", b.Left)
		}
	})
}

// TestStarAfterOperandIsMultiplication covers the operator role on the other
// side: a star with an operand to its left is a multiplication operator, so the
// star after it is that product's right operand (a wildcard), not the start of
// an implicit-AND term.
func TestStarAfterOperandIsMultiplication(t *testing.T) {
	t.Run("wildcard operand", func(t *testing.T) {
		p := mustParse(t, `a * * | stats count`)
		expr := p.Stages[0].(*ast.SearchStage).Expr
		b, ok := expr.(*ast.BinaryExpr)
		if !ok || b.Op != "*" {
			t.Fatalf("expr = %#v, want multiplication", expr)
		}
		if _, ok := b.Left.(*ast.SearchTerm); !ok {
			t.Fatalf("left = %#v, want the search term a", b.Left)
		}
		if _, ok := b.Right.(*ast.WildcardLit); !ok {
			t.Fatalf("right = %#v, want the wildcard", b.Right)
		}
	})

	t.Run("field operand", func(t *testing.T) {
		p := mustParse(t, `where a * not`)
		b, ok := p.Stages[0].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
		if !ok || b.Op != "*" {
			t.Fatalf("expr = %#v, want multiplication", p.Stages[0].(*ast.SearchStage).Expr)
		}
		if f, ok := b.Right.(*ast.FieldRef); !ok || f.Name != "not" {
			t.Fatalf("right = %#v, want FieldRef not", b.Right)
		}
	})
}

// TestStarNeedsAValueAfterTheOperator checks where the parser stops when a
// multiplication star has no operand: at the token that cannot begin one, in
// both term position and value position.
func TestStarNeedsAValueAfterTheOperator(t *testing.T) {
	tests := []struct {
		src  string
		line int
		col  int
	}{
		{`a * * *`, 1, 8},
		{`a * * * | stats count`, 1, 9},
		{`where * *`, 1, 10},
		{`where (a*)`, 1, 10},
		{`where (a * )`, 1, 12},
		{`stats sum(a*)`, 1, 13},
		{`stats sum(a * )`, 1, 15},
		{`where * not`, 1, 12},
	}
	for _, tc := range tests {
		t.Run(tc.src, func(t *testing.T) {
			_, err := Parse(tc.src)
			if err == nil {
				t.Fatal("expected an error")
			}
			ds, ok := diag.FromError(err)
			if !ok || len(ds) != 1 {
				t.Fatalf("err = %v, want one diagnostic", err)
			}
			d := ds[0]
			if d.Code != "syntax" || d.Pos.Line != tc.line || d.Pos.Col != tc.col {
				t.Fatalf("diagnostic = %#v, want syntax at %d:%d", d, tc.line, tc.col)
			}
		})
	}
}

// TestSortSignOnAKeywordField covers `sort` with a sign before a word that is
// otherwise the limit keyword or the negation keyword: the sign rules the
// keyword reading out, so the word is a field name.
func TestSortSignOnAKeywordField(t *testing.T) {
	tests := []struct {
		src  string
		name string
		desc bool
	}{
		{`head 5 | sort -limit`, "limit", true},
		{`head 5 | sort +limit`, "limit", false},
		{`head 5 | sort -not`, "not", true},
		{`head 5 | sort -limit, +host`, "limit", true},
	}
	for _, tc := range tests {
		t.Run(tc.src, func(t *testing.T) {
			p := mustParse(t, tc.src)
			s, ok := p.Stages[len(p.Stages)-1].(*ast.SortStage)
			if !ok || len(s.Fields) == 0 {
				t.Fatalf("last stage = %#v, want a sort field", p.Stages[len(p.Stages)-1])
			}
			if s.Fields[0].Name != tc.name || s.Fields[0].Desc != tc.desc {
				t.Fatalf("field = %#v, want %s desc=%v", s.Fields[0], tc.name, tc.desc)
			}
		})
	}
}

// TestBinRejectsSecondField covers the single-field rule: the second field is
// refused at its own token, before any later argument error can be reported.
func TestBinRejectsSecondField(t *testing.T) {
	const want = `bin accepts exactly one field; found "b" after "a"`
	tests := []string{`bin a b(`, `bin a b +`, `bin a b span=1h x(`, `bin a b where x ==`}
	for _, src := range tests {
		t.Run(src, func(t *testing.T) {
			_, err := Parse(src)
			if err == nil {
				t.Fatal("expected an error")
			}
			ds, ok := diag.FromError(err)
			if !ok || len(ds) != 1 {
				t.Fatalf("err = %v, want one diagnostic", err)
			}
			d := ds[0]
			if d.Code != "syntax" || d.Pos.Line != 1 || d.Pos.Col != 7 {
				t.Fatalf("diagnostic = %#v, want syntax at 1:7", d)
			}
			if d.Message != want {
				t.Fatalf("message = %q, want %q", d.Message, want)
			}
		})
	}
}
