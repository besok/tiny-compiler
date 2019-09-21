// Code generated from C:/projects/tiny-compiler/src/antlr\InterLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // InterLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by InterLangParser.
type InterLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by InterLangParser#file.
	VisitFile(ctx *FileContext) interface{}

	// Visit a parse tree produced by InterLangParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by InterLangParser#arg.
	VisitArg(ctx *ArgContext) interface{}

	// Visit a parse tree produced by InterLangParser#retTp.
	VisitRetTp(ctx *RetTpContext) interface{}

	// Visit a parse tree produced by InterLangParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by InterLangParser#internalArrArg.
	VisitInternalArrArg(ctx *InternalArrArgContext) interface{}

	// Visit a parse tree produced by InterLangParser#internalVar.
	VisitInternalVar(ctx *InternalVarContext) interface{}

	// Visit a parse tree produced by InterLangParser#newVar.
	VisitNewVar(ctx *NewVarContext) interface{}

	// Visit a parse tree produced by InterLangParser#updVar.
	VisitUpdVar(ctx *UpdVarContext) interface{}

	// Visit a parse tree produced by InterLangParser#initIntVar.
	VisitInitIntVar(ctx *InitIntVarContext) interface{}

	// Visit a parse tree produced by InterLangParser#initItem.
	VisitInitItem(ctx *InitItemContext) interface{}

	// Visit a parse tree produced by InterLangParser#initPrim.
	VisitInitPrim(ctx *InitPrimContext) interface{}

	// Visit a parse tree produced by InterLangParser#initParam.
	VisitInitParam(ctx *InitParamContext) interface{}

	// Visit a parse tree produced by InterLangParser#initCall.
	VisitInitCall(ctx *InitCallContext) interface{}

	// Visit a parse tree produced by InterLangParser#initBoolOp.
	VisitInitBoolOp(ctx *InitBoolOpContext) interface{}

	// Visit a parse tree produced by InterLangParser#initNumOp.
	VisitInitNumOp(ctx *InitNumOpContext) interface{}

	// Visit a parse tree produced by InterLangParser#initArrEl.
	VisitInitArrEl(ctx *InitArrElContext) interface{}

	// Visit a parse tree produced by InterLangParser#extArrEl.
	VisitExtArrEl(ctx *ExtArrElContext) interface{}

	// Visit a parse tree produced by InterLangParser#returnTp.
	VisitReturnTp(ctx *ReturnTpContext) interface{}

	// Visit a parse tree produced by InterLangParser#gotoIn.
	VisitGotoIn(ctx *GotoInContext) interface{}

	// Visit a parse tree produced by InterLangParser#gotoTp.
	VisitGotoTp(ctx *GotoTpContext) interface{}

	// Visit a parse tree produced by InterLangParser#ifFalse.
	VisitIfFalse(ctx *IfFalseContext) interface{}

	// Visit a parse tree produced by InterLangParser#typeTp.
	VisitTypeTp(ctx *TypeTpContext) interface{}

	// Visit a parse tree produced by InterLangParser#line.
	VisitLine(ctx *LineContext) interface{}
}
