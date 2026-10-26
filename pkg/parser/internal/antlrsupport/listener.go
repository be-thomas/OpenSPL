// Package antlrsupport turns ANTLR parse errors into positioned diagnostics.
//
// The SPL front-ends are fail-fast: the first syntax error ends the parse.
// ANTLR's stock error strategies recover and report several errors, and the
// ErrorStrategy interface carries an unexported method, so a strategy that
// bails cannot live outside the antlr package. Instead the listener records the
// first error as a diagnostic and panics with SyntaxAbort; the front-end
// recovers that panic in its entry point and returns the diagnostic.
package antlrsupport

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"

	"github.com/be-thomas/openspl/pkg/ast"
	"github.com/be-thomas/openspl/pkg/diag"
	"github.com/be-thomas/openspl/pkg/parser/internal/antlrpos"
)

// SyntaxAbort is the panic value the listener raises for the first syntax
// error. Err is the diagnostic the front-end returns.
type SyntaxAbort struct {
	Err *diag.Error
}

// Listener is an antlr.ErrorListener that aborts on the first syntax error.
//
// Recognizer-level messages (ANTLR's own ambiguity and context-sensitivity
// reports) are ignored: they describe the grammar, not the query.
type Listener struct {
	*antlr.DefaultErrorListener

	Map *antlrpos.SourceMap
}

// New returns a listener that maps positions through m.
func New(m *antlrpos.SourceMap) *Listener {
	return &Listener{DefaultErrorListener: antlr.NewDefaultErrorListener(), Map: m}
}

// SyntaxError records the first syntax error and panics. Message wording is
// ours, not ANTLR's, so diagnostics stay readable and stable across grammar
// changes; the code and position are what callers depend on.
func (l *Listener) SyntaxError(_ antlr.Recognizer, offendingSymbol interface{}, line, column int, _ string, _ antlr.RecognitionException) {
	pos := l.Map.PosAt(line, column)
	message := "unexpected token"
	if tok, ok := offendingSymbol.(antlr.Token); ok && tok != nil {
		pos = l.Map.Pos(tok)
		if tok.GetTokenType() == antlr.TokenEOF {
			message = "unexpected end of query"
		} else {
			message = fmt.Sprintf("unexpected %q", tok.GetText())
		}
	}
	panic(SyntaxAbort{Err: diag.NewError(pos, "syntax", "%s", message)})
}

// TokenAbort is a panic value raised by generated parser code that must end
// the parse at a specific token. Generated grammar code cannot import this
// package, so it declares its own type with these methods; the interface
// assertion in Recover is structural.
//
// Actions are preferable to predicates for this: ANTLR copies an action into
// the generated code and runs it only when the alternative is actually parsed,
// while a predicate is evaluated speculatively during prediction.
type TokenAbort interface {
	AbortToken() antlr.Token
	AbortCode() string
	AbortMessage() string
}

// Recover returns the diagnostic carried by an abort panic, mapping a
// TokenAbort through m. Any other panic value is re-raised so real bugs stay
// visible.
func Recover(r interface{}, m *antlrpos.SourceMap) (*diag.Error, bool) {
	switch abort := r.(type) {
	case SyntaxAbort:
		return abort.Err, true
	case TokenAbort:
		t := abort.AbortToken()
		return diag.NewError(m.Pos(t), abort.AbortCode(), "%s", abort.AbortMessage()), true
	}
	panic(r)
}

// Pos returns the position of a token, for front-end use.
func (l *Listener) Pos(t antlr.Token) ast.Pos { return l.Map.Pos(t) }
