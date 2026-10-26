package spl2

import (
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

func TestParseStages(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want []string
	}{
		{"from only", `from main`, []string{"search"}},
		{"where", `from main | where status=500`, []string{"search", "search"}},
		{"stats", `from main | stats count() AS count BY host`, []string{"search", "stats"}},
		{"full subset", `from main | where status=500 | stats count() AS count BY host`,
			[]string{"search", "search", "stats"}},
		{"no spaces around pipe", `from main|where status=500|stats count()`, []string{"search", "search", "stats"}},
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

// TestFromShape pins the lowering: `from <dataset>` becomes the same IR shape
// classic SPL produces for `index=<dataset>`, so the emitters resolve the
// index through one path.
func TestFromShape(t *testing.T) {
	tests := []struct {
		name    string
		src     string
		dataset string
	}{
		{"plain", `from main`, "main"},
		{"hyphenated", `from logs-2024-01`, "logs-2024-01"},
		{"trailing wildcard", `from logs*`, "logs*"},
		{"hyphen and wildcard", `from logs-*`, "logs-*"},
		{"keyword word", `from stats`, "stats"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := mustParse(t, tc.src)
			stage := p.Stages[0].(*ast.SearchStage)
			if stage.Pos != (ast.Pos{Line: 1, Col: 1}) {
				t.Errorf("stage position = %#v, want 1:1", stage.Pos)
			}
			b := stage.Expr.(*ast.BinaryExpr)
			if b.Op != "=" {
				t.Fatalf("operator = %q, want", b.Op)
			}
			f, ok := b.Left.(*ast.FieldRef)
			if !ok || f.Name != "index" || f.Pos != (ast.Pos{Line: 1, Col: 1}) {
				t.Fatalf("left side = %#v, want FieldRef index at 1:1", b.Left)
			}
			s, ok := b.Right.(*ast.StringLit)
			if !ok || s.Value != tc.dataset {
				t.Fatalf("right side = %#v, want StringLit %q", b.Right, tc.dataset)
			}
			if s.Pos != (ast.Pos{Line: 1, Col: 6, Offset: 5}) {
				t.Errorf("dataset position = %#v, want 1:6", s.Pos)
			}
		})
	}
}

// TestWhereSemantics pins the SPL2 rule that a bare word is a field reference
// everywhere, including on the right of a comparison, and that conditions need
// an explicit AND.
func TestWhereSemantics(t *testing.T) {
	t.Run("word on the right is a field reference", func(t *testing.T) {
		p := mustParse(t, `from main | where host=web`)
		b := p.Stages[1].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
		f, ok := b.Right.(*ast.FieldRef)
		if !ok || f.Name != "web" {
			t.Fatalf("right side = %#v, want FieldRef web", b.Right)
		}
	})

	t.Run("number stays a number", func(t *testing.T) {
		p := mustParse(t, `from main | where status=500`)
		b := p.Stages[1].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
		n, ok := b.Right.(*ast.NumberLit)
		if !ok || n.Value != 500 {
			t.Fatalf("right side = %#v, want NumberLit 500", b.Right)
		}
	})

	t.Run("quoted string stays a string", func(t *testing.T) {
		p := mustParse(t, `from main | where host="web"`)
		b := p.Stages[1].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
		if _, ok := b.Right.(*ast.StringLit); !ok {
			t.Fatalf("right side = %#v, want StringLit", b.Right)
		}
	})

	t.Run("function call", func(t *testing.T) {
		p := mustParse(t, `from main | where lower(host)="web"`)
		b := p.Stages[1].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
		c, ok := b.Left.(*ast.CallExpr)
		if !ok || c.Name != "lower" || len(c.Args) != 1 {
			t.Fatalf("left side = %#v, want lower(host)", b.Left)
		}
		if f, ok := c.Args[0].(*ast.FieldRef); !ok || f.Name != "host" {
			t.Fatalf("call argument = %#v, want FieldRef host", c.Args[0])
		}
	})

	t.Run("boolean and arithmetic", func(t *testing.T) {
		p := mustParse(t, `from main | where (a=1 OR b=2) AND NOT c>-1`)
		top := p.Stages[1].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
		if top.Op != "and" {
			t.Fatalf("top operator = %q, want and", top.Op)
		}
		if or, ok := top.Left.(*ast.BinaryExpr); !ok || or.Op != "or" {
			t.Fatalf("left side = %#v, want or", top.Left)
		}
		not, ok := top.Right.(*ast.UnaryExpr)
		if !ok || not.Op != "not" {
			t.Fatalf("right side = %#v, want not", top.Right)
		}
		cmp, ok := not.Operand.(*ast.BinaryExpr)
		if !ok || cmp.Op != ">" {
			t.Fatalf("negated operand = %#v, want >", not.Operand)
		}
		if u, ok := cmp.Right.(*ast.UnaryExpr); !ok || u.Op != "-" {
			t.Fatalf("right side = %#v, want unary minus", cmp.Right)
		}
	})

	t.Run("bang is not", func(t *testing.T) {
		p := mustParse(t, `from main | where !a=1`)
		u, ok := p.Stages[1].(*ast.SearchStage).Expr.(*ast.UnaryExpr)
		if !ok || u.Op != "not" {
			t.Fatalf("expression = %#v, want not", p.Stages[1].(*ast.SearchStage).Expr)
		}
	})
}

func TestStatsAggregates(t *testing.T) {
	t.Run("count without an argument", func(t *testing.T) {
		p := mustParse(t, `from main | stats count()`)
		s := p.Stages[1].(*ast.StatsStage)
		if len(s.Aggs) != 1 || s.Aggs[0].Func != "count" || s.Aggs[0].Arg != nil {
			t.Fatalf("aggs = %#v, want count with no argument", s.Aggs)
		}
	})

	t.Run("argument, alias, and by fields", func(t *testing.T) {
		p := mustParse(t, `from main | stats sum(bytes) AS total, dc(host) BY host, user`)
		s := p.Stages[1].(*ast.StatsStage)
		if len(s.Aggs) != 2 {
			t.Fatalf("got %d aggregates, want 2", len(s.Aggs))
		}
		if s.Aggs[0].Func != "sum" || s.Aggs[0].Alias != "total" {
			t.Errorf("first aggregate = %#v", s.Aggs[0])
		}
		if f, ok := s.Aggs[0].Arg.(*ast.FieldRef); !ok || f.Name != "bytes" {
			t.Errorf("first argument = %#v, want FieldRef bytes", s.Aggs[0].Arg)
		}
		if s.Aggs[1].Func != "dc" {
			t.Errorf("second aggregate = %#v, want dc", s.Aggs[1])
		}
		if len(s.By) != 2 {
			t.Fatalf("got %d by fields, want 2", len(s.By))
		}
		for i, want := range []string{"host", "user"} {
			f, ok := s.By[i].(*ast.FieldRef)
			if !ok || f.Name != want {
				t.Errorf("by field %d = %#v, want %q", i, s.By[i], want)
			}
		}
	})

	t.Run("distinct_count normalizes to dc", func(t *testing.T) {
		p := mustParse(t, `from main | stats distinct_count(user)`)
		s := p.Stages[1].(*ast.StatsStage)
		if s.Aggs[0].Func != "dc" {
			t.Fatalf("func = %q, want dc", s.Aggs[0].Func)
		}
	})
}

func TestRefusals(t *testing.T) {
	tests := []struct {
		src  string
		code string
		line int
		col  int
	}{
		{"", "empty-query", 1, 1},
		{"   ", "empty-query", 1, 1},

		// Pipeline shape.
		{"| from a", "syntax", 1, 1},
		{"where a=1", "syntax", 1, 1},
		{"main | stats count()", "syntax", 1, 1},
		{"from", "syntax", 1, 5},
		{"from *", "syntax", 1, 6},
		{"from logs *", "syntax", 1, 11},
		{"from a b", "syntax", 1, 8},
		{"from a | from b", "syntax", 1, 10},
		{"from a |", "syntax", 1, 8},
		{"from a || b", "syntax", 1, 9},

		// Commands outside the subset.
		{"from a | eval x=1", "unsupported-command", 1, 10},
		{"from a | sort -x", "unsupported-command", 1, 10},
		{"from a | frobnicate", "unknown-command", 1, 10},

		// Datasets.
		{"from datamodel:events", "unsupported-dataset", 1, 6},
		{"from lookup:foo", "unsupported-dataset", 1, 6},
		{"from DATA:foo", "unsupported-dataset", 1, 6},

		// Expressions.
		{"from a | where status=*", "syntax", 1, 23},
		{"from a | where [x]", "unsupported-subsearch", 1, 16},
		{"from a | where", "syntax", 1, 15},
		{"from a | where a AND", "syntax", 1, 21},
		{"from a | where a b", "syntax", 1, 18},
		{"from a | where =", "syntax", 1, 16},
		{"from a | where (a=1", "syntax", 1, 20},
		{"from a | where status IN (200,500)", "syntax", 1, 23},

		// Aggregates need the parentheses and a complete list.
		{"from a | stats count(*)", "syntax", 1, 22},
		{"from a | stats count", "syntax", 1, 21},
		{"from a | stats count() by", "syntax", 1, 26},
		{"from a | stats count() by host,", "syntax", 1, 32},

		// Lexical errors win over syntax, wherever they are.
		{"from a $ b", "syntax", 1, 8},
		{"from 1x", "syntax", 1, 6},
		{"from a | eval x=$", "syntax", 1, 17},
	}

	for _, tc := range tests {
		t.Run(tc.src, func(t *testing.T) {
			_, err := Parse(tc.src)
			if err == nil {
				t.Fatalf("Parse(%q) succeeded, want %s at %d:%d", tc.src, tc.code, tc.line, tc.col)
			}
			ds, ok := diag.FromError(err)
			if !ok || len(ds) == 0 {
				t.Fatalf("error is not a *diag.Error: %v", err)
			}
			d := ds[0]
			if d.Code != tc.code || d.Pos.Line != tc.line || d.Pos.Col != tc.col {
				t.Fatalf("got %s at %d:%d, want %s at %d:%d (%s)",
					d.Code, d.Pos.Line, d.Pos.Col, tc.code, tc.line, tc.col, d.Message)
			}
		})
	}
}

// TestPositions pins byte-based columns after multibyte text and across CRLF.
func TestPositions(t *testing.T) {
	t.Run("multibyte string", func(t *testing.T) {
		p := mustParse(t, `from m | where host="東京" AND status=500`)
		top := p.Stages[1].(*ast.SearchStage).Expr.(*ast.BinaryExpr)
		if top.Op != "and" {
			t.Fatalf("operator = %q, want and", top.Op)
		}
		if want := (ast.Pos{Line: 1, Col: 30, Offset: 29}); top.Pos != want {
			t.Errorf("AND position = %#v, want %#v", top.Pos, want)
		}
		right := top.Right.(*ast.BinaryExpr)
		if want := (ast.Pos{Line: 1, Col: 34, Offset: 33}); right.Left.Position() != want {
			t.Errorf("status position = %#v, want %#v", right.Left.Position(), want)
		}
	})

	t.Run("CRLF", func(t *testing.T) {
		p := mustParse(t, "from m\r\n| where host=\"東京\" | stats count()")
		if want := (ast.Pos{Line: 2, Col: 3, Offset: 10}); p.Stages[1].Position() != want {
			t.Errorf("where position = %#v, want %#v", p.Stages[1].Position(), want)
		}
		if want := (ast.Pos{Line: 2, Col: 25, Offset: 32}); p.Stages[2].Position() != want {
			t.Errorf("stats position = %#v, want %#v", p.Stages[2].Position(), want)
		}
	})
}
