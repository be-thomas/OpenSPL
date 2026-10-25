package spl1

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/be-thomas/openspl/pkg/ast"
	"github.com/be-thomas/openspl/pkg/diag"
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

type lexer struct {
	src  string
	i    int
	line int
	col  int
	toks []Token
}

// Lex tokenizes a classic SPL query. It is exported for the SPL2 front-end and
// for tooling that needs token-level access.
func Lex(src string) ([]Token, error) {
	l := &lexer{src: src, line: 1, col: 1}
	for {
		l.skipSpace()
		if l.i >= len(l.src) {
			l.toks = append(l.toks, Token{Kind: EOF, Pos: l.pos()})
			return l.toks, nil
		}

		pos := l.pos()
		ch := l.src[l.i]
		switch {
		case ch == '|':
			l.advance()
			l.emit(Pipe, "|", pos)
		case ch == ',':
			l.advance()
			l.emit(Comma, ",", pos)
		case ch == '(':
			l.advance()
			l.emit(LParen, "(", pos)
		case ch == ')':
			l.advance()
			l.emit(RParen, ")", pos)
		case ch == '[':
			l.advance()
			l.emit(LBracket, "[", pos)
		case ch == ']':
			l.advance()
			l.emit(RBracket, "]", pos)
		case ch == '"' || ch == '\'':
			text, err := l.scanString(ch, pos)
			if err != nil {
				return nil, err
			}
			l.emitString(text, pos)
		case isDigit(ch):
			if err := l.scanNumber(pos); err != nil {
				return nil, err
			}
		case isIdentStart(ch):
			start := l.i
			for l.i < len(l.src) && isIdentPart(l.src[l.i]) {
				l.advance()
			}
			l.emit(Ident, l.src[start:l.i], pos)
		case ch == '=':
			l.advance()
			text := "="
			if l.i < len(l.src) && l.src[l.i] == '=' {
				l.advance()
				text = "=="
			}
			l.emit(Eq, text, pos)
		case ch == '!':
			l.advance()
			if l.i < len(l.src) && l.src[l.i] == '=' {
				l.advance()
				l.emit(Ne, "!=", pos)
			} else {
				l.emit(Bang, "!", pos)
			}
		case ch == '<':
			l.advance()
			if l.i < len(l.src) && l.src[l.i] == '=' {
				l.advance()
				l.emit(Le, "<=", pos)
			} else {
				l.emit(Lt, "<", pos)
			}
		case ch == '>':
			l.advance()
			if l.i < len(l.src) && l.src[l.i] == '=' {
				l.advance()
				l.emit(Ge, ">=", pos)
			} else {
				l.emit(Gt, ">", pos)
			}
		case ch == '+':
			l.advance()
			l.emit(Plus, "+", pos)
		case ch == '-':
			l.advance()
			l.emit(Minus, "-", pos)
		case ch == '*':
			l.advance()
			l.emit(Star, "*", pos)
		case ch == '/':
			l.advance()
			l.emit(Slash, "/", pos)
		case ch == '%':
			l.advance()
			l.emit(Percent, "%", pos)
		default:
			r, size := utf8.DecodeRuneInString(l.src[l.i:])
			if r == utf8.RuneError && size == 1 {
				return nil, diag.NewError(pos, "syntax", "unexpected byte 0x%02X", ch)
			}
			return nil, diag.NewError(pos, "syntax", "unexpected character %q", string(r))
		}
	}
}

func (l *lexer) pos() ast.Pos {
	return ast.Pos{Line: l.line, Col: l.col, Offset: l.i}
}

func (l *lexer) emit(k Kind, text string, pos ast.Pos) {
	l.toks = append(l.toks, Token{Kind: k, Text: text, Pos: pos})
}

func (l *lexer) emitString(text string, pos ast.Pos) {
	l.toks = append(l.toks, Token{Kind: String, Text: text, Pos: pos})
}

func (l *lexer) advance() byte {
	ch := l.src[l.i]
	l.i++
	if ch == '\n' {
		l.line++
		l.col = 1
	} else {
		l.col++
	}
	return ch
}

func (l *lexer) skipSpace() {
	for l.i < len(l.src) {
		switch l.src[l.i] {
		case ' ', '\t', '\r', '\n':
			l.advance()
		default:
			return
		}
	}
}

func (l *lexer) scanString(quote byte, pos ast.Pos) (string, error) {
	l.advance() // opening quote
	var b strings.Builder
	for l.i < len(l.src) {
		ch := l.advance()
		if ch == quote {
			return b.String(), nil
		}
		// Only the quote itself is unescaped. Backslashes are preserved so
		// regular expressions survive intact (`\d`, `\S`, and `\\`).
		if ch == '\\' && l.i < len(l.src) && l.src[l.i] == quote {
			l.advance()
			b.WriteByte(quote)
			continue
		}
		b.WriteByte(ch)
	}
	return "", diag.NewError(pos, "syntax", "unterminated string literal")
}

func (l *lexer) scanNumber(pos ast.Pos) error {
	start := l.i
	for l.i < len(l.src) && isDigit(l.src[l.i]) {
		l.advance()
	}
	if l.i < len(l.src) && l.src[l.i] == '.' && l.i+1 < len(l.src) && isDigit(l.src[l.i+1]) {
		l.advance()
		for l.i < len(l.src) && isDigit(l.src[l.i]) {
			l.advance()
		}
	}
	// Scientific notation: 1e3, 2.5E-2.
	if l.i < len(l.src) && (l.src[l.i] == 'e' || l.src[l.i] == 'E') {
		j := l.i + 1
		if j < len(l.src) && (l.src[j] == '+' || l.src[j] == '-') {
			j++
		}
		if j < len(l.src) && isDigit(l.src[j]) {
			for l.i < j {
				l.advance()
			}
			for l.i < len(l.src) && isDigit(l.src[l.i]) {
				l.advance()
			}
		}
	}

	num := l.src[start:l.i]
	unit := l.scanDurationUnit()
	if unit != "" {
		l.toks = append(l.toks, Token{Kind: Duration, Text: num + unit, Pos: pos})
		return nil
	}
	if l.i < len(l.src) && isLetter(l.src[l.i]) {
		return diag.NewError(pos, "syntax", "invalid numeric literal %q", l.src[start:l.i+1])
	}
	if _, err := strconv.ParseFloat(num, 64); err != nil {
		return diag.NewError(pos, "syntax", "invalid number %q", num)
	}
	l.toks = append(l.toks, Token{Kind: Number, Text: num, Pos: pos})
	return nil
}

// scanDurationUnit consumes the longest valid SPL time unit at the cursor.
func (l *lexer) scanDurationUnit() string {
	rest := l.src[l.i:]
	for _, unit := range []string{"mon", "ms", "s", "m", "h", "d", "w", "y"} {
		if strings.HasPrefix(rest, unit) {
			after := l.i + len(unit)
			if after < len(l.src) && isLetter(l.src[after]) {
				continue
			}
			for range unit {
				l.advance()
			}
			return unit
		}
	}
	return ""
}

func isDigit(b byte) bool { return b >= '0' && b <= '9' }

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isIdentStart(b byte) bool {
	return b == '_' || b == '@' || isLetter(b)
}

func isIdentPart(b byte) bool {
	return isIdentStart(b) || isDigit(b) || b == '.' || b == ':'
}
