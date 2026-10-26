package spl1

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/be-thomas/openspl/pkg/ast"
	"github.com/be-thomas/openspl/pkg/diag"
	"github.com/be-thomas/openspl/pkg/parser/commands"
	"github.com/be-thomas/openspl/pkg/parser/internal/antlrpos"
	"github.com/be-thomas/openspl/pkg/parser/internal/antlrsupport"
	"github.com/be-thomas/openspl/pkg/parser/spl1/gen"
)

// Parse parses a classic SPL query into the unified IR.
func Parse(src string) (pl *ast.Pipeline, err error) {
	if strings.TrimSpace(src) == "" {
		return nil, diag.NewError(ast.Pos{Line: 1, Col: 1}, "empty-query", "query is empty")
	}

	m := antlrpos.New(src)
	stream, err := newTokenStream(m, src)
	if err != nil {
		return nil, err
	}

	parser := spl1gen.NewSPL1Parser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(antlrsupport.New(m))

	b := &builder{m: m, toks: stream.GetAllTokens(), parser: parser}

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
// command and expression structure; the driver owns stage sequencing, command
// classification, and the leading/trailing-pipe rules, because fail-fast
// priority depends on stopping at the first stage-level problem.
func (b *builder) pipeline() (*ast.Pipeline, error) {
	ts := b.parser.GetTokenStream()

	// A leading pipe is tolerated: `| stats count` means the same as
	// `stats count` with an implicit match-all search.
	if ts.LA(1) == spl1gen.SPL1ParserPIPE {
		ts.Consume()
	}
	if ts.LA(1) == antlr.TokenEOF {
		return nil, diag.NewError(b.pos(ts.LT(1)), "empty-query", "query is empty")
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
func (b *builder) stage(first bool) (ast.Stage, spl1gen.IStageEndContext, error) {
	ts := b.parser.GetTokenStream()
	cur := ts.LT(1)

	if isIdentLikeToken(cur.GetTokenType()) {
		name := strings.ToLower(cur.GetText())
		if commands.SPL1Implemented(name) {
			ctx := b.parser.StageCommand()
			st, err := b.command(ctx)
			return st, ctx.StageEnd(), err
		}
		if commands.SPL1Known(name) {
			return nil, nil, diag.NewError(b.pos(cur), "unsupported-command",
				"command %q is recognized but not implemented yet; see docs/COMPATIBILITY.md", name)
		}
		if !first {
			return nil, nil, diag.NewError(b.pos(cur), "unknown-command", "unknown command %q", name)
		}
		ctx := b.parser.StageSearch()
		st, err := b.searchStage(ctx)
		return st, ctx.StageEnd(), err
	}

	if !first {
		// Not a command keyword: the grammar reports the unexpected token.
		ctx := b.parser.StageCommand()
		st, err := b.command(ctx)
		return st, ctx.StageEnd(), err
	}
	ctx := b.parser.StageSearch()
	st, err := b.searchStage(ctx)
	return st, ctx.StageEnd(), err
}

// isIdentLikeToken mirrors the grammar's identLike rule for stage dispatch.
func isIdentLikeToken(t int) bool {
	switch t {
	case spl1gen.SPL1ParserIDENT, spl1gen.SPL1ParserSEARCH, spl1gen.SPL1ParserWHERE,
		spl1gen.SPL1ParserSTATS, spl1gen.SPL1ParserTIMECHART, spl1gen.SPL1ParserEVAL,
		spl1gen.SPL1ParserREX, spl1gen.SPL1ParserTABLE, spl1gen.SPL1ParserFIELDS,
		spl1gen.SPL1ParserSORT, spl1gen.SPL1ParserHEAD, spl1gen.SPL1ParserBIN,
		spl1gen.SPL1ParserRENAME, spl1gen.SPL1ParserBY, spl1gen.SPL1ParserAS,
		spl1gen.SPL1ParserSPAN, spl1gen.SPL1ParserLIMIT, spl1gen.SPL1ParserFIELD,
		spl1gen.SPL1ParserAND, spl1gen.SPL1ParserOR, spl1gen.SPL1ParserNOT,
		spl1gen.SPL1ParserNOT_WORD:
		return true
	}
	return false
}
