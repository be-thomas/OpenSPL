package spl2

import (
	"fmt"
	"unicode/utf8"

	"github.com/antlr4-go/antlr/v4"

	"github.com/be-thomas/openspl/pkg/diag"
	"github.com/be-thomas/openspl/pkg/parser/internal/antlrpos"
	"github.com/be-thomas/openspl/pkg/parser/spl2/gen"
)

// newTokenStream runs the generated lexer over src, fills the stream so every
// token is available up front, and applies the invalid-token pass. Lexing the
// whole input before parsing keeps the front-end's priority: a lexical error
// in a later stage beats a syntax error in an earlier one.
func newTokenStream(m *antlrpos.SourceMap, src string) (*antlr.CommonTokenStream, []antlr.Token, error) {
	lexer := spl2gen.NewSPL2Lexer(antlr.NewInputStream(src))
	lexer.RemoveErrorListeners()
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	stream.Fill()
	toks := stream.GetAllTokens()
	if err := checkTokens(m, toks, src); err != nil {
		return nil, nil, err
	}
	return stream, toks, nil
}

// checkTokens reports the first lexically invalid token in source order.
//
// ANTLR's lexer never fails: a character no rule accepts becomes UNKNOWN, and
// a numeric literal immediately followed by a letter lexes as a valid number
// plus an identifier. Both are lexical errors, reported at the start of the
// offending token before any parsing begins.
func checkTokens(m *antlrpos.SourceMap, toks []antlr.Token, src string) error {
	for i, t := range toks {
		switch t.GetTokenType() {
		case spl2gen.SPL2ParserUNKNOWN:
			return invalidCharacter(m, t, src)
		case spl2gen.SPL2ParserNUMBER, spl2gen.SPL2ParserDURATION:
			next := nextToken(toks, i)
			if next == nil || next.GetTokenType() == antlr.TokenEOF {
				continue
			}
			if t.GetStop()+1 != next.GetStart() {
				continue
			}
			at := m.ByteOffset(next.GetStart())
			if at >= len(src) || !isLetter(src[at]) {
				continue
			}
			start := m.ByteOffset(t.GetStart())
			end := at + 1
			if end > len(src) {
				end = len(src)
			}
			return diag.NewError(m.Pos(t), "syntax",
				"invalid numeric literal %q", src[start:end])
		}
	}
	return nil
}

func nextToken(toks []antlr.Token, i int) antlr.Token {
	if i+1 < len(toks) {
		return toks[i+1]
	}
	return nil
}

func invalidCharacter(m *antlrpos.SourceMap, t antlr.Token, src string) error {
	at := m.ByteOffset(t.GetStart())
	if at >= len(src) {
		return diag.NewError(m.Pos(t), "syntax", "unexpected end of query")
	}
	r, size := utf8.DecodeRuneInString(src[at:])
	if r == utf8.RuneError && size == 1 {
		return diag.NewError(m.Pos(t), "syntax", "unexpected byte 0x%02X", src[at])
	}
	return diag.NewError(m.Pos(t), "syntax", "unexpected character %q", string(r))
}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

// describe names a token the way a diagnostic should, using SPL2 vocabulary.
func describe(t antlr.Token) string {
	if t == nil || t.GetTokenType() == antlr.TokenEOF {
		return "end of query"
	}
	switch t.GetTokenType() {
	case spl2gen.SPL2ParserIDENT, spl2gen.SPL2ParserFROM, spl2gen.SPL2ParserWHERE,
		spl2gen.SPL2ParserSTATS, spl2gen.SPL2ParserBY, spl2gen.SPL2ParserAS,
		spl2gen.SPL2ParserAND, spl2gen.SPL2ParserOR, spl2gen.SPL2ParserNOT:
		return fmt.Sprintf("identifier %q", t.GetText())
	case spl2gen.SPL2ParserSTRING:
		return fmt.Sprintf("string %q", decodeString(t.GetText()))
	case spl2gen.SPL2ParserNUMBER:
		return fmt.Sprintf("number %q", t.GetText())
	case spl2gen.SPL2ParserDURATION:
		return fmt.Sprintf("duration %q", t.GetText())
	}
	return fmt.Sprintf("%q", t.GetText())
}

// decodeString returns the value of a string literal token: the delimiters are
// dropped and only an escaped delimiter is unescaped. Every other backslash is
// kept so regular expressions survive intact.
func decodeString(raw string) string {
	if len(raw) < 2 {
		return ""
	}
	quote := raw[0]
	body := raw[1 : len(raw)-1]
	out := make([]byte, 0, len(body))
	for i := 0; i < len(body); i++ {
		if body[i] == '\\' && i+1 < len(body) && body[i+1] == quote {
			out = append(out, quote)
			i++
			continue
		}
		out = append(out, body[i])
	}
	return string(out)
}
