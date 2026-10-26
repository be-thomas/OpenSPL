package spl1

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/antlr4-go/antlr/v4"

	"github.com/be-thomas/openspl/pkg/ast"
	"github.com/be-thomas/openspl/pkg/diag"
	"github.com/be-thomas/openspl/pkg/parser/internal/antlrpos"
	"github.com/be-thomas/openspl/pkg/parser/spl1/gen"
)

// Kind is a lexical token kind.
type Kind int

const (
	EOF Kind = iota
	Ident
	String
	Number
	Duration
	Pipe
	Comma
	LParen
	RParen
	LBracket
	RBracket
	Eq
	Ne
	Lt
	Le
	Gt
	Ge
	Plus
	Minus
	Star
	Slash
	Percent
	Bang
)

func (k Kind) String() string {
	switch k {
	case EOF:
		return "end of query"
	case Ident:
		return "identifier"
	case String:
		return "string literal"
	case Number:
		return "number"
	case Duration:
		return "duration"
	case Pipe:
		return "'|'"
	case Comma:
		return "','"
	case LParen:
		return "'('"
	case RParen:
		return "')'"
	case LBracket:
		return "'['"
	case RBracket:
		return "']'"
	case Eq:
		return "'='"
	case Ne:
		return "'!='"
	case Lt:
		return "'<'"
	case Le:
		return "'<='"
	case Gt:
		return "'>'"
	case Ge:
		return "'>='"
	case Plus:
		return "'+'"
	case Minus:
		return "'-'"
	case Star:
		return "'*'"
	case Slash:
		return "'/'"
	case Percent:
		return "'%'"
	case Bang:
		return "'!'"
	}
	return "token"
}

// Token is one lexed unit with its source position.
type Token struct {
	Kind Kind
	Text string
	Pos  ast.Pos
}

func (t Token) describe() string {
	switch t.Kind {
	case EOF:
		return "end of query"
	case Ident:
		return fmt.Sprintf("identifier %q", t.Text)
	case String:
		return fmt.Sprintf("string %q", t.Text)
	case Number, Duration:
		return fmt.Sprintf("%s %q", t.Kind, t.Text)
	}
	return t.Kind.String()
}

// Lex tokenizes a classic SPL query. It is exported for tooling that needs
// token-level access and is a thin adapter over the generated ANTLR lexer:
// keyword tokens are reported as identifiers, and string literals carry their
// decoded value.
func Lex(src string) ([]Token, error) {
	m := antlrpos.New(src)
	stream, err := newTokenStream(m, src)
	if err != nil {
		return nil, err
	}
	toks := stream.GetAllTokens()
	out := make([]Token, 0, len(toks))
	for _, t := range toks {
		tok, err := convertToken(m, t)
		if err != nil {
			return nil, err
		}
		out = append(out, tok)
	}
	return out, nil
}

// newTokenStream runs the generated lexer over src, fills the stream so every
// token is available up front, and applies the invalid-token pass. Lexing the
// whole input before parsing keeps the hand parser's priority: a lexical error
// in a later stage beats a semantic error in an earlier one. The returned
// stream answers the parser's lookahead with the value-run token types (see
// tightStream); GetAllTokens still returns the lexer's own tokens.
func newTokenStream(m *antlrpos.SourceMap, src string) (*tightStream, error) {
	lexer := spl1gen.NewSPL1Lexer(antlr.NewInputStream(src))
	lexer.RemoveErrorListeners()
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	stream.Fill()
	toks := stream.GetAllTokens()
	if err := checkTokens(m, toks, src); err != nil {
		return nil, err
	}
	return &tightStream{CommonTokenStream: stream, types: tightTokenTypes(toks)}, nil
}

// checkTokens reports the first lexically invalid token in source order.
//
// ANTLR's lexer never fails: a character no rule accepts becomes UNKNOWN, a
// numeric literal immediately followed by a letter lexes as a valid number
// plus an identifier, and a number too large for a float64 still lexes as
// NUMBER. All three are lexical errors in SPL, reported at the start of the
// offending token before any parsing begins.
func checkTokens(m *antlrpos.SourceMap, toks []antlr.Token, src string) error {
	for i, t := range toks {
		switch t.GetTokenType() {
		case spl1gen.SPL1ParserUNKNOWN:
			return invalidCharacter(m, t, src)
		case spl1gen.SPL1ParserNUMBER, spl1gen.SPL1ParserDURATION:
			// A number outside float64 is invalid, like any other lexical
			// error, and is reported before parsing begins.
			if t.GetTokenType() == spl1gen.SPL1ParserNUMBER {
				if _, err := strconv.ParseFloat(t.GetText(), 64); err != nil {
					return diag.NewError(m.Pos(t), "syntax",
						"invalid number %q", t.GetText())
				}
			}
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
			// The literal is malformed because a letter-run continues it, so
			// quote the whole run rather than the numeric token alone.
			end := m.ByteOffset(next.GetStop() + 1)
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

// convertToken maps a generated token onto the public token kind.
func convertToken(m *antlrpos.SourceMap, t antlr.Token) (Token, error) {
	pos := m.Pos(t)
	switch t.GetTokenType() {
	case spl1gen.SPL1ParserEOF:
		return Token{Kind: EOF, Pos: pos}, nil
	case spl1gen.SPL1ParserIDENT,
		spl1gen.SPL1ParserSEARCH, spl1gen.SPL1ParserWHERE, spl1gen.SPL1ParserSTATS,
		spl1gen.SPL1ParserTIMECHART, spl1gen.SPL1ParserEVAL, spl1gen.SPL1ParserREX,
		spl1gen.SPL1ParserTABLE, spl1gen.SPL1ParserFIELDS, spl1gen.SPL1ParserSORT,
		spl1gen.SPL1ParserHEAD, spl1gen.SPL1ParserBIN, spl1gen.SPL1ParserBY,
		spl1gen.SPL1ParserAS, spl1gen.SPL1ParserSPAN, spl1gen.SPL1ParserLIMIT,
		spl1gen.SPL1ParserFIELD, spl1gen.SPL1ParserAND, spl1gen.SPL1ParserOR,
		spl1gen.SPL1ParserNOT:
		return Token{Kind: Ident, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserSTRING:
		return Token{Kind: String, Text: decodeString(t.GetText()), Pos: pos}, nil
	case spl1gen.SPL1ParserNUMBER:
		return Token{Kind: Number, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserDURATION:
		return Token{Kind: Duration, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserPIPE:
		return Token{Kind: Pipe, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserCOMMA:
		return Token{Kind: Comma, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserLPAREN:
		return Token{Kind: LParen, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserRPAREN:
		return Token{Kind: RParen, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserLBRACKET:
		return Token{Kind: LBracket, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserRBRACKET:
		return Token{Kind: RBracket, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserEQ2, spl1gen.SPL1ParserEQ:
		return Token{Kind: Eq, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserNE:
		return Token{Kind: Ne, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserLT:
		return Token{Kind: Lt, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserLE:
		return Token{Kind: Le, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserGT:
		return Token{Kind: Gt, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserGE:
		return Token{Kind: Ge, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserPLUS:
		return Token{Kind: Plus, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserMINUS:
		return Token{Kind: Minus, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserSTAR:
		return Token{Kind: Star, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserSLASH:
		return Token{Kind: Slash, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserPERCENT:
		return Token{Kind: Percent, Text: t.GetText(), Pos: pos}, nil
	case spl1gen.SPL1ParserBANG:
		return Token{Kind: Bang, Text: t.GetText(), Pos: pos}, nil
	}
	return Token{}, invalidCharacterFrom(m, t)
}

// invalidCharacterFrom is the defensive path for a token type convertToken
// does not know; checkTokens has already refused UNKNOWN, so reaching it means
// the grammar gained a token the adapter was not updated for.
func invalidCharacterFrom(m *antlrpos.SourceMap, t antlr.Token) error {
	return diag.NewError(m.Pos(t), "syntax", "unexpected %q", t.GetText())
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
	var b strings.Builder
	b.Grow(len(body))
	for i := 0; i < len(body); i++ {
		if body[i] == '\\' && i+1 < len(body) && body[i+1] == quote {
			b.WriteByte(quote)
			i++
			continue
		}
		b.WriteByte(body[i])
	}
	return b.String()
}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}
