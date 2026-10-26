// Code generated from pkg/parser/spl1/grammar/SPL1.g4 by ANTLR 4.13.2. DO NOT EDIT.

package spl1gen // SPL1
import "github.com/antlr4-go/antlr/v4"

type BaseSPL1Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSPL1Visitor) VisitStageCommand(ctx *StageCommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitStageSearch(ctx *StageSearchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitStageEnd(ctx *StageEndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitCmdSearch(ctx *CmdSearchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitCmdWhere(ctx *CmdWhereContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitCmdStats(ctx *CmdStatsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitCmdTimechart(ctx *CmdTimechartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitCmdEval(ctx *CmdEvalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitCmdRex(ctx *CmdRexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitCmdTable(ctx *CmdTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitCmdFields(ctx *CmdFieldsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitCmdSort(ctx *CmdSortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitCmdHead(ctx *CmdHeadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitCmdBin(ctx *CmdBinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSearchExpr(ctx *SearchExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSearchOrTail(ctx *SearchOrTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSearchAnd(ctx *SearchAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSAndExplicit(ctx *SAndExplicitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSAndImplicit(ctx *SAndImplicitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSNotUnary(ctx *SNotUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSNotPrimary(ctx *SNotPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSearchComparison(ctx *SearchComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSearchComparisonTail(ctx *SearchComparisonTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSOpString(ctx *SOpStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSOpValue(ctx *SOpValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSOpExpr(ctx *SOpExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSearchOperandValue(ctx *SearchOperandValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitRhsTail(ctx *RhsTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSearchOperandExpr(ctx *SearchOperandExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSearchOperandUnary(ctx *SearchOperandUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSearchOperandPrimary(ctx *SearchOperandPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSearchAdditive(ctx *SearchAdditiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSearchAddTail(ctx *SearchAddTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSearchMultiplicative(ctx *SearchMultiplicativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSearchMulTail(ctx *SearchMulTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSUnarySign(ctx *SUnarySignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSUnaryPrimary(ctx *SUnaryPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSPrimParen(ctx *SPrimParenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSPrimString(ctx *SPrimStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSPrimWord(ctx *SPrimWordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSPrimNumber(ctx *SPrimNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSPrimDuration(ctx *SPrimDurationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSPrimStarRun(ctx *SPrimStarRunContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSPrimStar(ctx *SPrimStarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSPrimSubsearch(ctx *SPrimSubsearchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSearchTermTail(ctx *SearchTermTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitValueExpr(ctx *ValueExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitValueOrTail(ctx *ValueOrTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitValueAnd(ctx *ValueAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitVAndExplicit(ctx *VAndExplicitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitVAndImplicit(ctx *VAndImplicitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitVNotUnary(ctx *VNotUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitVNotPrimary(ctx *VNotPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitValueComparison(ctx *ValueComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitValueComparisonTail(ctx *ValueComparisonTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitValueAdd(ctx *ValueAddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitValueAddTail(ctx *ValueAddTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitValueMul(ctx *ValueMulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitValueMulTail(ctx *ValueMulTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitVUnarySign(ctx *VUnarySignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitVUnaryPrimary(ctx *VUnaryPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitVPrimParen(ctx *VPrimParenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitVPrimCall(ctx *VPrimCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitVPrimWord(ctx *VPrimWordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitVPrimString(ctx *VPrimStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitVPrimNumber(ctx *VPrimNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitVPrimDuration(ctx *VPrimDurationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitVPrimStar(ctx *VPrimStarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitVPrimSubsearch(ctx *VPrimSubsearchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSubsearch(ctx *SubsearchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitAbsentOperand(ctx *AbsentOperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitComparisonOp(ctx *ComparisonOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitStar(ctx *StarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitMinus(ctx *MinusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitAddOp(ctx *AddOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitMulOp(ctx *MulOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSign(ctx *SignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitIdentLike(ctx *IdentLikeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitAggList(ctx *AggListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitAgg(ctx *AggContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitAggListNotSpan(ctx *AggListNotSpanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitAggNotSpan(ctx *AggNotSpanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitAliasClause(ctx *AliasClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitByClause(ctx *ByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitTcSpanFirst(ctx *TcSpanFirstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitTcAggFirst(ctx *TcAggFirstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitTimechartSpan(ctx *TimechartSpanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSpanAssign(ctx *SpanAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitRexField(ctx *RexFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitFieldExprList(ctx *FieldExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitFieldExpr(ctx *FieldExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSpanValue(ctx *SpanValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitTableItem(ctx *TableItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitFieldsSign(ctx *FieldsSignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitFieldsItem(ctx *FieldsItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSortLimit(ctx *SortLimitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSortFirstField(ctx *SortFirstFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSortEmpty(ctx *SortEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSortItem(ctx *SortItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSortFirstItem(ctx *SortFirstItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitSortField(ctx *SortFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitLimitKeyword(ctx *LimitKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitLimitNumber(ctx *LimitNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitLimitTail(ctx *LimitTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitBinSpanItem(ctx *BinSpanItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitBinFieldItem(ctx *BinFieldItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitBinField(ctx *BinFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitIdentLikeNotSpan(ctx *IdentLikeNotSpanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSPL1Visitor) VisitIdentLikeNotLimit(ctx *IdentLikeNotLimitContext) interface{} {
	return v.VisitChildren(ctx)
}
