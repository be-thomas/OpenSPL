// Package spl2 is the front-end for the supported subset of Splunk's SPL2.
//
// The subset is `from` followed by any number of `where` and `stats` stages:
//
//	from main | where status=500 | stats count() AS count BY host
//
// SPL2 differs from classic SPL in both surface syntax and semantics: a
// pipeline always starts at `from`, a bare word is a field reference even on
// the right of a comparison, and conditions need an explicit `AND`. The
// dialect therefore has its own grammar (grammar/SPL2.g4), generated parser,
// and lowering; it shares only the source map and the error listener with the
// classic front-end, and neither front-end imports the other.
//
// Everything outside the subset is refused with a positioned diagnostic rather
// than approximated: namespaced datasets (a `datamodel:` or any other
// colon-qualified name) cannot become an Elasticsearch index, commands other
// than `from`, `where`, and `stats` are recognized and refused, and a `*`
// inside an expression is a syntax error because the subset has no meaning for
// a wildcard there.
package spl2

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/be-thomas/openspl/pkg/ast"
	"github.com/be-thomas/openspl/pkg/diag"
	"github.com/be-thomas/openspl/pkg/parser/commands"
	"github.com/be-thomas/openspl/pkg/parser/internal/antlrpos"
	"github.com/be-thomas/openspl/pkg/parser/internal/antlrsupport"
	"github.com/be-thomas/openspl/pkg/parser/spl2/gen"
)

// Parse parses an SPL2 query into the unified IR.
func Parse(src string) (pl *ast.Pipeline, err error) {
	if strings.TrimSpace(src) == "" {
		return nil, diag.NewError(ast.Pos{Line: 1, Col: 1}, "empty-query", "query is empty")
	}

	m := antlrpos.New(src)
	stream, toks, err := newTokenStream(m, src)
	if err != nil {
		return nil, err
	}

	parser := spl2gen.NewSPL2Parser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(antlrsupport.New(m))

	b := &builder{m: m, toks: toks, parser: parser}

	// The listener aborts the parse by panicking on the first syntax error;
	// any other panic is a bug and is re-raised by Recover.
	defer func() {
		if r := recover(); r != nil {
			if e, ok := antlrsupport.Recover(r, m); ok {
				pl, err = nil, e
			}
		}
	}()
	return b.pipeline()
}

// pipeline drives the entry rules one stage at a time. The grammar owns
// dataset names, expressions, and aggregates; the driver owns pipeline shape
// and command classification, because the first stage must be `from` and a
// recognized command has to be refused before it is parsed as something else.
func (b *builder) pipeline() (*ast.Pipeline, error) {
	ts := b.parser.GetTokenStream()
	cur := ts.LT(1)

	// Unlike classic SPL, a leading pipe is not accepted: the pipeline starts
	// at `from`, so any other first token is an error at that token.
	if cur.GetTokenType() != spl2gen.SPL2ParserFROM {
		return nil, diag.NewError(b.pos(cur), "syntax",
			"an SPL2 pipeline must start with 'from'")
	}

	pl := &ast.Pipeline{}
	for first := true; ; first = false {
		stage, end, err := b.stage(first)
		if err != nil {
			return nil, err
		}
		pl.Stages = append(pl.Stages, stage)

		// stageEnd is part of every entry rule and was just consumed. The
		// stream does not advance past EOF, so the rule context, not LT(-1),
		// says whether the stage ended at a pipe.
		if end.EOF() != nil {
			return pl, nil
		}
		if ts.LA(1) == antlr.TokenEOF {
			return nil, diag.NewError(b.pos(end.PIPE().GetSymbol()), "syntax",
				"expected a command after '|'")
		}
	}
}

// stage classifies the stage start and parses exactly one stage.
func (b *builder) stage(first bool) (ast.Stage, spl2gen.IStageEndContext, error) {
	ts := b.parser.GetTokenStream()
	cur := ts.LT(1)

	if isIdentLikeToken(cur.GetTokenType()) {
		name := strings.ToLower(cur.GetText())
		if commands.SPL2Implemented(name) {
			return b.stageCommand(name, first)
		}
		if commands.SPL2Known(name) {
			return nil, nil, diag.NewError(b.pos(cur), "unsupported-command",
				"command %q is recognized but not implemented for SPL2 yet; see docs/COMPATIBILITY.md", name)
		}
		return nil, nil, diag.NewError(b.pos(cur), "unknown-command", "unknown command %q", name)
	}

	// Not a word: let the grammar blame the token. `stageWhere` requires a
	// command keyword, so the error lands on this token.
	ctx := b.parser.StageWhere()
	st, err := b.where(ctx)
	return st, ctx.StageEnd(), err
}

// stageCommand parses one implemented stage. The name comes from the command
// table, so a command added there without a case here is a programming error,
// not a user-visible condition.
func (b *builder) stageCommand(name string, first bool) (ast.Stage, spl2gen.IStageEndContext, error) {
	switch name {
	case "from":
		cur := b.parser.GetTokenStream().LT(1)
		if !first {
			return nil, nil, diag.NewError(b.pos(cur), "syntax",
				"'from' must be the first stage")
		}
		ctx := b.parser.StageFrom()
		st, err := b.from(ctx)
		return st, ctx.StageEnd(), err

	case "where":
		ctx := b.parser.StageWhere()
		st, err := b.where(ctx)
		return st, ctx.StageEnd(), err

	case "stats":
		ctx := b.parser.StageStats()
		st, err := b.stats(ctx)
		return st, ctx.StageEnd(), err
	}

	pos := b.pos(b.parser.GetTokenStream().LT(1))
	return nil, nil, diag.NewError(pos, "internal", "command %q has no SPL2 grammar rule", name)
}

// isIdentLikeToken mirrors the grammar's identLike rule for stage dispatch.
func isIdentLikeToken(t int) bool {
	switch t {
	case spl2gen.SPL2ParserIDENT, spl2gen.SPL2ParserFROM, spl2gen.SPL2ParserWHERE,
		spl2gen.SPL2ParserSTATS, spl2gen.SPL2ParserBY, spl2gen.SPL2ParserAS,
		spl2gen.SPL2ParserAND, spl2gen.SPL2ParserOR, spl2gen.SPL2ParserNOT:
		return true
	}
	return false
}
