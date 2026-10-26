// Code generated from pkg/parser/spl2/grammar/SPL2.g4 by ANTLR 4.13.2. DO NOT EDIT.

package spl2gen // SPL2
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SPL2Parser.
type SPL2Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SPL2Parser#stageFrom.
	VisitStageFrom(ctx *StageFromContext) interface{}

	// Visit a parse tree produced by SPL2Parser#stageWhere.
	VisitStageWhere(ctx *StageWhereContext) interface{}

	// Visit a parse tree produced by SPL2Parser#stageStats.
	VisitStageStats(ctx *StageStatsContext) interface{}

	// Visit a parse tree produced by SPL2Parser#stageEnd.
	VisitStageEnd(ctx *StageEndContext) interface{}

	// Visit a parse tree produced by SPL2Parser#indexName.
	VisitIndexName(ctx *IndexNameContext) interface{}

	// Visit a parse tree produced by SPL2Parser#indexTail.
	VisitIndexTail(ctx *IndexTailContext) interface{}

	// Visit a parse tree produced by SPL2Parser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by SPL2Parser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by SPL2Parser#orTail.
	VisitOrTail(ctx *OrTailContext) interface{}

	// Visit a parse tree produced by SPL2Parser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by SPL2Parser#andTail.
	VisitAndTail(ctx *AndTailContext) interface{}

	// Visit a parse tree produced by SPL2Parser#notExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by SPL2Parser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by SPL2Parser#comparisonTail.
	VisitComparisonTail(ctx *ComparisonTailContext) interface{}

	// Visit a parse tree produced by SPL2Parser#additive.
	VisitAdditive(ctx *AdditiveContext) interface{}

	// Visit a parse tree produced by SPL2Parser#additiveTail.
	VisitAdditiveTail(ctx *AdditiveTailContext) interface{}

	// Visit a parse tree produced by SPL2Parser#multiplicative.
	VisitMultiplicative(ctx *MultiplicativeContext) interface{}

	// Visit a parse tree produced by SPL2Parser#multiplicativeTail.
	VisitMultiplicativeTail(ctx *MultiplicativeTailContext) interface{}

	// Visit a parse tree produced by SPL2Parser#unary.
	VisitUnary(ctx *UnaryContext) interface{}

	// Visit a parse tree produced by SPL2Parser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by SPL2Parser#subsearch.
	VisitSubsearch(ctx *SubsearchContext) interface{}

	// Visit a parse tree produced by SPL2Parser#absentOperand.
	VisitAbsentOperand(ctx *AbsentOperandContext) interface{}

	// Visit a parse tree produced by SPL2Parser#aggList.
	VisitAggList(ctx *AggListContext) interface{}

	// Visit a parse tree produced by SPL2Parser#agg.
	VisitAgg(ctx *AggContext) interface{}

	// Visit a parse tree produced by SPL2Parser#aliasClause.
	VisitAliasClause(ctx *AliasClauseContext) interface{}

	// Visit a parse tree produced by SPL2Parser#byClause.
	VisitByClause(ctx *ByClauseContext) interface{}

	// Visit a parse tree produced by SPL2Parser#fieldList.
	VisitFieldList(ctx *FieldListContext) interface{}

	// Visit a parse tree produced by SPL2Parser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by SPL2Parser#compOp.
	VisitCompOp(ctx *CompOpContext) interface{}

	// Visit a parse tree produced by SPL2Parser#addOp.
	VisitAddOp(ctx *AddOpContext) interface{}

	// Visit a parse tree produced by SPL2Parser#mulOp.
	VisitMulOp(ctx *MulOpContext) interface{}

	// Visit a parse tree produced by SPL2Parser#sign.
	VisitSign(ctx *SignContext) interface{}

	// Visit a parse tree produced by SPL2Parser#identLike.
	VisitIdentLike(ctx *IdentLikeContext) interface{}
}
