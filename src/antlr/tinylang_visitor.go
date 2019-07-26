// Code generated from C:/projects/tiny-compiler/src/antlr\TinyLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // TinyLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by TinyLangParser.
type TinyLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TinyLangParser#file.
	VisitFile(ctx *FileContext) interface{}

	// Visit a parse tree produced by TinyLangParser#funcInit.
	VisitFuncInit(ctx *FuncInitContext) interface{}

	// Visit a parse tree produced by TinyLangParser#args.
	VisitArgs(ctx *ArgsContext) interface{}

	// Visit a parse tree produced by TinyLangParser#returnSt.
	VisitReturnSt(ctx *ReturnStContext) interface{}

	// Visit a parse tree produced by TinyLangParser#funcInvoc.
	VisitFuncInvoc(ctx *FuncInvocContext) interface{}

	// Visit a parse tree produced by TinyLangParser#argsInvoc.
	VisitArgsInvoc(ctx *ArgsInvocContext) interface{}

	// Visit a parse tree produced by TinyLangParser#argInvoc.
	VisitArgInvoc(ctx *ArgInvocContext) interface{}

	// Visit a parse tree produced by TinyLangParser#updVar.
	VisitUpdVar(ctx *UpdVarContext) interface{}

	// Visit a parse tree produced by TinyLangParser#newVar.
	VisitNewVar(ctx *NewVarContext) interface{}

	// Visit a parse tree produced by TinyLangParser#varType.
	VisitVarType(ctx *VarTypeContext) interface{}

	// Visit a parse tree produced by TinyLangParser#arrayElem.
	VisitArrayElem(ctx *ArrayElemContext) interface{}

	// Visit a parse tree produced by TinyLangParser#arrayInit.
	VisitArrayInit(ctx *ArrayInitContext) interface{}

	// Visit a parse tree produced by TinyLangParser#arrayInitEmp.
	VisitArrayInitEmp(ctx *ArrayInitEmpContext) interface{}

	// Visit a parse tree produced by TinyLangParser#arrayInitVal.
	VisitArrayInitVal(ctx *ArrayInitValContext) interface{}

	// Visit a parse tree produced by TinyLangParser#arrayInitElems.
	VisitArrayInitElems(ctx *ArrayInitElemsContext) interface{}

	// Visit a parse tree produced by TinyLangParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by TinyLangParser#exprOp.
	VisitExprOp(ctx *ExprOpContext) interface{}

	// Visit a parse tree produced by TinyLangParser#boolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by TinyLangParser#boolExprOp.
	VisitBoolExprOp(ctx *BoolExprOpContext) interface{}

	// Visit a parse tree produced by TinyLangParser#commonBody.
	VisitCommonBody(ctx *CommonBodyContext) interface{}

	// Visit a parse tree produced by TinyLangParser#ifElse.
	VisitIfElse(ctx *IfElseContext) interface{}

	// Visit a parse tree produced by TinyLangParser#iF.
	VisitIF(ctx *IFContext) interface{}

	// Visit a parse tree produced by TinyLangParser#elseIf.
	VisitElseIf(ctx *ElseIfContext) interface{}

	// Visit a parse tree produced by TinyLangParser#elsePart.
	VisitElsePart(ctx *ElsePartContext) interface{}

	// Visit a parse tree produced by TinyLangParser#while.
	VisitWhile(ctx *WhileContext) interface{}

	// Visit a parse tree produced by TinyLangParser#forSt.
	VisitForSt(ctx *ForStContext) interface{}
}
