package lexerparity

import (
	"fmt"
	"testing"

	"github.com/antlr4-go/antlr/v4"

	spl1gen "github.com/be-thomas/openspl/pkg/parser/spl1/gen"
	spl2gen "github.com/be-thomas/openspl/pkg/parser/spl2/gen"
)

// tok is the parts of a token both lexers must agree on. The token type is
// compared through its symbolic name, because the two grammars declare
// different keyword sets and therefore number the shared rules differently.
type tok struct {
	class string
	text  string
	start int
	stop  int
	line  int
	col   int
}

// corpus is lexical material only: identifiers, numbers, durations, strings,
// operators, whitespace, and characters no rule accepts. Words that are
// keywords in one dialect but not the other are excluded on purpose, because
// those differences are the dialects' business; the shared keyword spellings
// are included so their rules cannot drift.
var corpus = []struct {
	name string
	src  string
}{
	{"identifiers", `_x @timestamp a.b a:b A1 x_y @@@ a.b:c`},
	{"numbers", `0 12 3.14 1e10 1E-3 1.5e+2 007 1.0`},
	{"durations", `1mon 2ms 3s 4m 5h 6d 7w 8y 1.5h 10mon`},
	{"number-letter adjacency", `1month 5monx 1e+foo 1.2.3 1. 1e 1e+ 1MON 1M`},
	{"strings", `"a b" 'a b' "a\"b" 'a\'b' "\\d+" "café" "東京" "" ''`},
	{"unterminated string", `"a\\"`},
	{"multiline string", "x=\"line\nbreak\" y"},
	{"operators", `== != <= >= = < > ! | , ( ) [ ] + - * / %`},
	{"unknown characters", `$ # ~ & ? ` + "`" + ` .`},
	{"adjacency", `logs-2024-01 web*prod *error* a** a-b a - b a- b a -b`},
	{"spaces and tabs", "a\tb  c"},
	{"CRLF", "a\r\nb\r\nc"},
	{"multibyte around tokens", `a="café" b 東京 c`},
	{"shared keywords", `where stats by as and or not`},
}

func TestLexerParity(t *testing.T) {
	for _, tc := range corpus {
		t.Run(tc.name, func(t *testing.T) {
			got1, got2 := lexSPL1(tc.src), lexSPL2(tc.src)
			if len(got1) != len(got2) {
				t.Fatalf("token count: SPL1 %d, SPL2 %d\nSPL1: %v\nSPL2: %v",
					len(got1), len(got2), got1, got2)
			}
			for i := range got1 {
				if got1[i] != got2[i] {
					t.Errorf("token %d: SPL1 %+v, SPL2 %+v", i, got1[i], got2[i])
				}
			}
		})
	}
}

// TestCorpusCoversEverySharedRule fails if a corpus entry is edited in a way
// that stops exercising a duplicated rule, which would let the parity test pass
// while a rule drifts untested.
func TestCorpusCoversEverySharedRule(t *testing.T) {
	seen := make(map[string]bool)
	for _, tc := range corpus {
		for _, tk := range lexSPL1(tc.src) {
			seen[tk.class] = true
		}
	}

	want := []string{
		"IDENT", "NUMBER", "DURATION", "STRING",
		"EQ2", "NE", "LE", "GE", "EQ", "LT", "GT", "BANG",
		"PIPE", "COMMA", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET",
		"PLUS", "MINUS", "STAR", "SLASH", "PERCENT", "UNKNOWN",
		"WHERE", "STATS", "BY", "AS", "AND", "OR", "NOT",
	}
	for _, name := range want {
		if !seen[name] {
			t.Errorf("corpus does not produce a %s token", name)
		}
	}
}

func lexSPL1(src string) []tok {
	lx := spl1gen.NewSPL1Lexer(antlr.NewInputStream(src))
	lx.RemoveErrorListeners()
	ts := antlr.NewCommonTokenStream(lx, antlr.TokenDefaultChannel)
	ts.Fill()
	return collect(ts.GetAllTokens(), lx.GetSymbolicNames())
}

func lexSPL2(src string) []tok {
	lx := spl2gen.NewSPL2Lexer(antlr.NewInputStream(src))
	lx.RemoveErrorListeners()
	ts := antlr.NewCommonTokenStream(lx, antlr.TokenDefaultChannel)
	ts.Fill()
	return collect(ts.GetAllTokens(), lx.GetSymbolicNames())
}

func collect(toks []antlr.Token, names []string) []tok {
	out := make([]tok, 0, len(toks))
	for _, t := range toks {
		out = append(out, tok{
			class: class(t, names),
			text:  t.GetText(),
			start: t.GetStart(),
			stop:  t.GetStop(),
			line:  t.GetLine(),
			col:   t.GetColumn(),
		})
	}
	return out
}

func class(t antlr.Token, names []string) string {
	tt := t.GetTokenType()
	if tt == antlr.TokenEOF {
		return "EOF"
	}
	if tt > 0 && tt < len(names) {
		return names[tt]
	}
	return fmt.Sprintf("token %d", tt)
}
