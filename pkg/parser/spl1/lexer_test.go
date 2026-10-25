package spl1

import (
	"testing"

	"github.com/be-thomas/openspl/pkg/ast"
	"github.com/be-thomas/openspl/pkg/diag"
)

func TestLexPositions(t *testing.T) {
	toks, err := Lex("index=web\nmsg=\"café\" host=web")
	if err != nil {
		t.Fatal(err)
	}
	want := []struct {
		kind Kind
		text string
		pos  ast.Pos
	}{
		{Ident, "index", ast.Pos{Line: 1, Col: 1, Offset: 0}},
		{Eq, "=", ast.Pos{Line: 1, Col: 6, Offset: 5}},
		{Ident, "web", ast.Pos{Line: 1, Col: 7, Offset: 6}},
		{Ident, "msg", ast.Pos{Line: 2, Col: 1, Offset: 10}},
		{Eq, "=", ast.Pos{Line: 2, Col: 4, Offset: 13}},
		{String, "café", ast.Pos{Line: 2, Col: 5, Offset: 14}},
		// "café" is six bytes, so the next token's column counts bytes, not
		// runes: 12 bytes into the line is column 13.
		{Ident, "host", ast.Pos{Line: 2, Col: 13, Offset: 22}},
		{Eq, "=", ast.Pos{Line: 2, Col: 17, Offset: 26}},
		{Ident, "web", ast.Pos{Line: 2, Col: 18, Offset: 27}},
		{EOF, "", ast.Pos{Line: 2, Col: 21, Offset: 30}},
	}
	if len(toks) != len(want) {
		t.Fatalf("got %d tokens, want %d: %#v", len(toks), len(want), toks)
	}
	for i, w := range want {
		got := toks[i]
		if got.Kind != w.kind || got.Text != w.text || got.Pos != w.pos {
			t.Errorf("token %d = {%v %q %+v}, want {%v %q %+v}",
				i, got.Kind, got.Text, got.Pos, w.kind, w.text, w.pos)
		}
	}
}

func TestLexCRLFPositions(t *testing.T) {
	toks, err := Lex("a\r\nb")
	if err != nil {
		t.Fatal(err)
	}
	if toks[0].Pos != (ast.Pos{Line: 1, Col: 1, Offset: 0}) {
		t.Fatalf("first token = %+v", toks[0].Pos)
	}
	if toks[1].Pos != (ast.Pos{Line: 2, Col: 1, Offset: 3}) {
		t.Fatalf("second token = %+v, want line 2 col 1 offset 3", toks[1].Pos)
	}
}

func TestLexInvalidRuneReportsTheCharacter(t *testing.T) {
	_, err := Lex(`index=wéb`)
	if err == nil {
		t.Fatal("expected an error")
	}
	ds, ok := diag.FromError(err)
	if !ok || len(ds) != 1 {
		t.Fatalf("err = %v, want one diagnostic", err)
	}
	d := ds[0]
	if d.Code != "syntax" || d.Pos.Line != 1 || d.Pos.Col != 8 {
		t.Fatalf("diagnostic = %#v, want syntax at 1:8", d)
	}
	if want := `unexpected character "é"`; d.Message != want {
		t.Fatalf("message = %q, want %q", d.Message, want)
	}
}

func TestLexInvalidByteIsReportedAsAByte(t *testing.T) {
	_, err := Lex("index=w\xffb")
	if err == nil {
		t.Fatal("expected an error")
	}
	ds, ok := diag.FromError(err)
	if !ok || len(ds) != 1 {
		t.Fatalf("err = %v, want one diagnostic", err)
	}
	d := ds[0]
	if d.Code != "syntax" || d.Pos.Line != 1 || d.Pos.Col != 8 {
		t.Fatalf("diagnostic = %#v, want syntax at 1:8", d)
	}
	if want := "unexpected byte 0xFF"; d.Message != want {
		t.Fatalf("message = %q, want %q", d.Message, want)
	}
}

func TestParsePositionAfterMultibyteString(t *testing.T) {
	_, err := Parse(`index=web msg="café" +`)
	if err == nil {
		t.Fatal("expected an error")
	}
	ds, _ := diag.FromError(err)
	if len(ds) == 0 || ds[0].Code != "syntax" || ds[0].Pos.Line != 1 || ds[0].Pos.Col != 23 {
		t.Fatalf("diagnostics = %#v, want syntax at 1:23", ds)
	}
}
