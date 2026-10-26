// Code generated from pkg/parser/spl1/grammar/SPL1.g4 by ANTLR 4.13.2. DO NOT EDIT.

package spl1gen // SPL1
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SPL1Parser.
type SPL1Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SPL1Parser#stageCommand.
	VisitStageCommand(ctx *StageCommandContext) interface{}

	// Visit a parse tree produced by SPL1Parser#stageSearch.
	VisitStageSearch(ctx *StageSearchContext) interface{}

	// Visit a parse tree produced by SPL1Parser#stageEnd.
	VisitStageEnd(ctx *StageEndContext) interface{}

	// Visit a parse tree produced by SPL1Parser#CmdSearch.
	VisitCmdSearch(ctx *CmdSearchContext) interface{}

	// Visit a parse tree produced by SPL1Parser#CmdWhere.
	VisitCmdWhere(ctx *CmdWhereContext) interface{}

	// Visit a parse tree produced by SPL1Parser#CmdStats.
	VisitCmdStats(ctx *CmdStatsContext) interface{}

	// Visit a parse tree produced by SPL1Parser#CmdTimechart.
	VisitCmdTimechart(ctx *CmdTimechartContext) interface{}

	// Visit a parse tree produced by SPL1Parser#CmdEval.
	VisitCmdEval(ctx *CmdEvalContext) interface{}

	// Visit a parse tree produced by SPL1Parser#CmdRex.
	VisitCmdRex(ctx *CmdRexContext) interface{}

	// Visit a parse tree produced by SPL1Parser#CmdTable.
	VisitCmdTable(ctx *CmdTableContext) interface{}

	// Visit a parse tree produced by SPL1Parser#CmdFields.
	VisitCmdFields(ctx *CmdFieldsContext) interface{}

	// Visit a parse tree produced by SPL1Parser#CmdSort.
	VisitCmdSort(ctx *CmdSortContext) interface{}

	// Visit a parse tree produced by SPL1Parser#CmdHead.
	VisitCmdHead(ctx *CmdHeadContext) interface{}

	// Visit a parse tree produced by SPL1Parser#CmdBin.
	VisitCmdBin(ctx *CmdBinContext) interface{}

	// Visit a parse tree produced by SPL1Parser#CmdRename.
	VisitCmdRename(ctx *CmdRenameContext) interface{}

	// Visit a parse tree produced by SPL1Parser#CmdTstats.
	VisitCmdTstats(ctx *CmdTstatsContext) interface{}

	// Visit a parse tree produced by SPL1Parser#searchExpr.
	VisitSearchExpr(ctx *SearchExprContext) interface{}

	// Visit a parse tree produced by SPL1Parser#searchOrTail.
	VisitSearchOrTail(ctx *SearchOrTailContext) interface{}

	// Visit a parse tree produced by SPL1Parser#searchAnd.
	VisitSearchAnd(ctx *SearchAndContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SAndExplicit.
	VisitSAndExplicit(ctx *SAndExplicitContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SAndImplicit.
	VisitSAndImplicit(ctx *SAndImplicitContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SNotUnary.
	VisitSNotUnary(ctx *SNotUnaryContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SNotPrimary.
	VisitSNotPrimary(ctx *SNotPrimaryContext) interface{}

	// Visit a parse tree produced by SPL1Parser#searchComparison.
	VisitSearchComparison(ctx *SearchComparisonContext) interface{}

	// Visit a parse tree produced by SPL1Parser#searchComparisonTail.
	VisitSearchComparisonTail(ctx *SearchComparisonTailContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SOpString.
	VisitSOpString(ctx *SOpStringContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SOpValue.
	VisitSOpValue(ctx *SOpValueContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SOpExpr.
	VisitSOpExpr(ctx *SOpExprContext) interface{}

	// Visit a parse tree produced by SPL1Parser#searchOperandValue.
	VisitSearchOperandValue(ctx *SearchOperandValueContext) interface{}

	// Visit a parse tree produced by SPL1Parser#rhsTail.
	VisitRhsTail(ctx *RhsTailContext) interface{}

	// Visit a parse tree produced by SPL1Parser#searchOperandExpr.
	VisitSearchOperandExpr(ctx *SearchOperandExprContext) interface{}

	// Visit a parse tree produced by SPL1Parser#searchOperandUnary.
	VisitSearchOperandUnary(ctx *SearchOperandUnaryContext) interface{}

	// Visit a parse tree produced by SPL1Parser#searchOperandPrimary.
	VisitSearchOperandPrimary(ctx *SearchOperandPrimaryContext) interface{}

	// Visit a parse tree produced by SPL1Parser#searchAdditive.
	VisitSearchAdditive(ctx *SearchAdditiveContext) interface{}

	// Visit a parse tree produced by SPL1Parser#searchAddTail.
	VisitSearchAddTail(ctx *SearchAddTailContext) interface{}

	// Visit a parse tree produced by SPL1Parser#searchMultiplicative.
	VisitSearchMultiplicative(ctx *SearchMultiplicativeContext) interface{}

	// Visit a parse tree produced by SPL1Parser#searchMulTail.
	VisitSearchMulTail(ctx *SearchMulTailContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SUnarySign.
	VisitSUnarySign(ctx *SUnarySignContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SUnaryPrimary.
	VisitSUnaryPrimary(ctx *SUnaryPrimaryContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SPrimParen.
	VisitSPrimParen(ctx *SPrimParenContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SPrimString.
	VisitSPrimString(ctx *SPrimStringContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SPrimWord.
	VisitSPrimWord(ctx *SPrimWordContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SPrimNumber.
	VisitSPrimNumber(ctx *SPrimNumberContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SPrimDuration.
	VisitSPrimDuration(ctx *SPrimDurationContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SPrimStarRun.
	VisitSPrimStarRun(ctx *SPrimStarRunContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SPrimStar.
	VisitSPrimStar(ctx *SPrimStarContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SPrimSubsearch.
	VisitSPrimSubsearch(ctx *SPrimSubsearchContext) interface{}

	// Visit a parse tree produced by SPL1Parser#searchTermTail.
	VisitSearchTermTail(ctx *SearchTermTailContext) interface{}

	// Visit a parse tree produced by SPL1Parser#valueExpr.
	VisitValueExpr(ctx *ValueExprContext) interface{}

	// Visit a parse tree produced by SPL1Parser#valueOrTail.
	VisitValueOrTail(ctx *ValueOrTailContext) interface{}

	// Visit a parse tree produced by SPL1Parser#valueAnd.
	VisitValueAnd(ctx *ValueAndContext) interface{}

	// Visit a parse tree produced by SPL1Parser#VAndExplicit.
	VisitVAndExplicit(ctx *VAndExplicitContext) interface{}

	// Visit a parse tree produced by SPL1Parser#VAndImplicit.
	VisitVAndImplicit(ctx *VAndImplicitContext) interface{}

	// Visit a parse tree produced by SPL1Parser#VNotUnary.
	VisitVNotUnary(ctx *VNotUnaryContext) interface{}

	// Visit a parse tree produced by SPL1Parser#VNotPrimary.
	VisitVNotPrimary(ctx *VNotPrimaryContext) interface{}

	// Visit a parse tree produced by SPL1Parser#valueComparison.
	VisitValueComparison(ctx *ValueComparisonContext) interface{}

	// Visit a parse tree produced by SPL1Parser#valueComparisonTail.
	VisitValueComparisonTail(ctx *ValueComparisonTailContext) interface{}

	// Visit a parse tree produced by SPL1Parser#valueAdd.
	VisitValueAdd(ctx *ValueAddContext) interface{}

	// Visit a parse tree produced by SPL1Parser#valueAddTail.
	VisitValueAddTail(ctx *ValueAddTailContext) interface{}

	// Visit a parse tree produced by SPL1Parser#valueMul.
	VisitValueMul(ctx *ValueMulContext) interface{}

	// Visit a parse tree produced by SPL1Parser#valueMulTail.
	VisitValueMulTail(ctx *ValueMulTailContext) interface{}

	// Visit a parse tree produced by SPL1Parser#VUnarySign.
	VisitVUnarySign(ctx *VUnarySignContext) interface{}

	// Visit a parse tree produced by SPL1Parser#VUnaryPrimary.
	VisitVUnaryPrimary(ctx *VUnaryPrimaryContext) interface{}

	// Visit a parse tree produced by SPL1Parser#VPrimParen.
	VisitVPrimParen(ctx *VPrimParenContext) interface{}

	// Visit a parse tree produced by SPL1Parser#VPrimCall.
	VisitVPrimCall(ctx *VPrimCallContext) interface{}

	// Visit a parse tree produced by SPL1Parser#VPrimWord.
	VisitVPrimWord(ctx *VPrimWordContext) interface{}

	// Visit a parse tree produced by SPL1Parser#VPrimString.
	VisitVPrimString(ctx *VPrimStringContext) interface{}

	// Visit a parse tree produced by SPL1Parser#VPrimNumber.
	VisitVPrimNumber(ctx *VPrimNumberContext) interface{}

	// Visit a parse tree produced by SPL1Parser#VPrimDuration.
	VisitVPrimDuration(ctx *VPrimDurationContext) interface{}

	// Visit a parse tree produced by SPL1Parser#VPrimStar.
	VisitVPrimStar(ctx *VPrimStarContext) interface{}

	// Visit a parse tree produced by SPL1Parser#VPrimSubsearch.
	VisitVPrimSubsearch(ctx *VPrimSubsearchContext) interface{}

	// Visit a parse tree produced by SPL1Parser#subsearch.
	VisitSubsearch(ctx *SubsearchContext) interface{}

	// Visit a parse tree produced by SPL1Parser#absentOperand.
	VisitAbsentOperand(ctx *AbsentOperandContext) interface{}

	// Visit a parse tree produced by SPL1Parser#comparisonOp.
	VisitComparisonOp(ctx *ComparisonOpContext) interface{}

	// Visit a parse tree produced by SPL1Parser#operandStar.
	VisitOperandStar(ctx *OperandStarContext) interface{}

	// Visit a parse tree produced by SPL1Parser#star.
	VisitStar(ctx *StarContext) interface{}

	// Visit a parse tree produced by SPL1Parser#minus.
	VisitMinus(ctx *MinusContext) interface{}

	// Visit a parse tree produced by SPL1Parser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by SPL1Parser#addOp.
	VisitAddOp(ctx *AddOpContext) interface{}

	// Visit a parse tree produced by SPL1Parser#mulOp.
	VisitMulOp(ctx *MulOpContext) interface{}

	// Visit a parse tree produced by SPL1Parser#searchMulOp.
	VisitSearchMulOp(ctx *SearchMulOpContext) interface{}

	// Visit a parse tree produced by SPL1Parser#sign.
	VisitSign(ctx *SignContext) interface{}

	// Visit a parse tree produced by SPL1Parser#identLike.
	VisitIdentLike(ctx *IdentLikeContext) interface{}

	// Visit a parse tree produced by SPL1Parser#exprWord.
	VisitExprWord(ctx *ExprWordContext) interface{}

	// Visit a parse tree produced by SPL1Parser#aggList.
	VisitAggList(ctx *AggListContext) interface{}

	// Visit a parse tree produced by SPL1Parser#agg.
	VisitAgg(ctx *AggContext) interface{}

	// Visit a parse tree produced by SPL1Parser#aggListNotSpan.
	VisitAggListNotSpan(ctx *AggListNotSpanContext) interface{}

	// Visit a parse tree produced by SPL1Parser#aggNotSpan.
	VisitAggNotSpan(ctx *AggNotSpanContext) interface{}

	// Visit a parse tree produced by SPL1Parser#aliasClause.
	VisitAliasClause(ctx *AliasClauseContext) interface{}

	// Visit a parse tree produced by SPL1Parser#byClause.
	VisitByClause(ctx *ByClauseContext) interface{}

	// Visit a parse tree produced by SPL1Parser#tstatsBody.
	VisitTstatsBody(ctx *TstatsBodyContext) interface{}

	// Visit a parse tree produced by SPL1Parser#tstatsAggList.
	VisitTstatsAggList(ctx *TstatsAggListContext) interface{}

	// Visit a parse tree produced by SPL1Parser#tstatsAgg.
	VisitTstatsAgg(ctx *TstatsAggContext) interface{}

	// Visit a parse tree produced by SPL1Parser#tstatsAggWord.
	VisitTstatsAggWord(ctx *TstatsAggWordContext) interface{}

	// Visit a parse tree produced by SPL1Parser#tstatsOption.
	VisitTstatsOption(ctx *TstatsOptionContext) interface{}

	// Visit a parse tree produced by SPL1Parser#tstatsOptionValue.
	VisitTstatsOptionValue(ctx *TstatsOptionValueContext) interface{}

	// Visit a parse tree produced by SPL1Parser#tstatsFrom.
	VisitTstatsFrom(ctx *TstatsFromContext) interface{}

	// Visit a parse tree produced by SPL1Parser#tstatsDatamodel.
	VisitTstatsDatamodel(ctx *TstatsDatamodelContext) interface{}

	// Visit a parse tree produced by SPL1Parser#tstatsWhere.
	VisitTstatsWhere(ctx *TstatsWhereContext) interface{}

	// Visit a parse tree produced by SPL1Parser#tstatsBy.
	VisitTstatsBy(ctx *TstatsByContext) interface{}

	// Visit a parse tree produced by SPL1Parser#tstatsByItem.
	VisitTstatsByItem(ctx *TstatsByItemContext) interface{}

	// Visit a parse tree produced by SPL1Parser#TcSpanFirst.
	VisitTcSpanFirst(ctx *TcSpanFirstContext) interface{}

	// Visit a parse tree produced by SPL1Parser#TcAggFirst.
	VisitTcAggFirst(ctx *TcAggFirstContext) interface{}

	// Visit a parse tree produced by SPL1Parser#timechartSpan.
	VisitTimechartSpan(ctx *TimechartSpanContext) interface{}

	// Visit a parse tree produced by SPL1Parser#spanAssign.
	VisitSpanAssign(ctx *SpanAssignContext) interface{}

	// Visit a parse tree produced by SPL1Parser#rexField.
	VisitRexField(ctx *RexFieldContext) interface{}

	// Visit a parse tree produced by SPL1Parser#fieldExprList.
	VisitFieldExprList(ctx *FieldExprListContext) interface{}

	// Visit a parse tree produced by SPL1Parser#fieldExpr.
	VisitFieldExpr(ctx *FieldExprContext) interface{}

	// Visit a parse tree produced by SPL1Parser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by SPL1Parser#spanValue.
	VisitSpanValue(ctx *SpanValueContext) interface{}

	// Visit a parse tree produced by SPL1Parser#tableItem.
	VisitTableItem(ctx *TableItemContext) interface{}

	// Visit a parse tree produced by SPL1Parser#fieldsSign.
	VisitFieldsSign(ctx *FieldsSignContext) interface{}

	// Visit a parse tree produced by SPL1Parser#fieldsItem.
	VisitFieldsItem(ctx *FieldsItemContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SortLimit.
	VisitSortLimit(ctx *SortLimitContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SortFirstField.
	VisitSortFirstField(ctx *SortFirstFieldContext) interface{}

	// Visit a parse tree produced by SPL1Parser#SortEmpty.
	VisitSortEmpty(ctx *SortEmptyContext) interface{}

	// Visit a parse tree produced by SPL1Parser#sortItem.
	VisitSortItem(ctx *SortItemContext) interface{}

	// Visit a parse tree produced by SPL1Parser#sortFirstItem.
	VisitSortFirstItem(ctx *SortFirstItemContext) interface{}

	// Visit a parse tree produced by SPL1Parser#sortField.
	VisitSortField(ctx *SortFieldContext) interface{}

	// Visit a parse tree produced by SPL1Parser#LimitKeyword.
	VisitLimitKeyword(ctx *LimitKeywordContext) interface{}

	// Visit a parse tree produced by SPL1Parser#LimitNumber.
	VisitLimitNumber(ctx *LimitNumberContext) interface{}

	// Visit a parse tree produced by SPL1Parser#limitTail.
	VisitLimitTail(ctx *LimitTailContext) interface{}

	// Visit a parse tree produced by SPL1Parser#BinSpanItem.
	VisitBinSpanItem(ctx *BinSpanItemContext) interface{}

	// Visit a parse tree produced by SPL1Parser#BinFieldItem.
	VisitBinFieldItem(ctx *BinFieldItemContext) interface{}

	// Visit a parse tree produced by SPL1Parser#renamePair.
	VisitRenamePair(ctx *RenamePairContext) interface{}

	// Visit a parse tree produced by SPL1Parser#renameField.
	VisitRenameField(ctx *RenameFieldContext) interface{}

	// Visit a parse tree produced by SPL1Parser#binField.
	VisitBinField(ctx *BinFieldContext) interface{}

	// Visit a parse tree produced by SPL1Parser#identLikeNotSpan.
	VisitIdentLikeNotSpan(ctx *IdentLikeNotSpanContext) interface{}

	// Visit a parse tree produced by SPL1Parser#identLikeNotLimit.
	VisitIdentLikeNotLimit(ctx *IdentLikeNotLimitContext) interface{}
}
