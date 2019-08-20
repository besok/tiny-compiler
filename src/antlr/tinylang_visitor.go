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

	// Visit a parse tree produced by TinyLangParser#funcArg.
	VisitFuncArg(ctx *FuncArgContext) interface{}

	// Visit a parse tree produced by TinyLangParser#funcArgs.
	VisitFuncArgs(ctx *FuncArgsContext) interface{}

	// Visit a parse tree produced by TinyLangParser#funcReturn.
	VisitFuncReturn(ctx *FuncReturnContext) interface{}

	// Visit a parse tree produced by TinyLangParser#funcReturnType.
	VisitFuncReturnType(ctx *FuncReturnTypeContext) interface{}

	// Visit a parse tree produced by TinyLangParser#funcInvoc.
	VisitFuncInvoc(ctx *FuncInvocContext) interface{}

	// Visit a parse tree produced by TinyLangParser#funcArgsInvoc.
	VisitFuncArgsInvoc(ctx *FuncArgsInvocContext) interface{}

	// Visit a parse tree produced by TinyLangParser#funcInvocArgs.
	VisitFuncInvocArgs(ctx *FuncInvocArgsContext) interface{}

	// Visit a parse tree produced by TinyLangParser#updVariable.
	VisitUpdVariable(ctx *UpdVariableContext) interface{}

	// Visit a parse tree produced by TinyLangParser#newVariable.
	VisitNewVariable(ctx *NewVariableContext) interface{}

	// Visit a parse tree produced by TinyLangParser#val.
	VisitVal(ctx *ValContext) interface{}

	// Visit a parse tree produced by TinyLangParser#variableType.
	VisitVariableType(ctx *VariableTypeContext) interface{}

	// Visit a parse tree produced by TinyLangParser#arrayElem.
	VisitArrayElem(ctx *ArrayElemContext) interface{}

	// Visit a parse tree produced by TinyLangParser#arrayInit.
	VisitArrayInit(ctx *ArrayInitContext) interface{}

	// Visit a parse tree produced by TinyLangParser#arrayInitEmpty.
	VisitArrayInitEmpty(ctx *ArrayInitEmptyContext) interface{}

	// Visit a parse tree produced by TinyLangParser#arrayInitValue.
	VisitArrayInitValue(ctx *ArrayInitValueContext) interface{}

	// Visit a parse tree produced by TinyLangParser#arrayInitElems.
	VisitArrayInitElems(ctx *ArrayInitElemsContext) interface{}

	// Visit a parse tree produced by TinyLangParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by TinyLangParser#exprOperand.
	VisitExprOperand(ctx *ExprOperandContext) interface{}

	// Visit a parse tree produced by TinyLangParser#boolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by TinyLangParser#boolExprSingleExtra.
	VisitBoolExprSingleExtra(ctx *BoolExprSingleExtraContext) interface{}

	// Visit a parse tree produced by TinyLangParser#boolExprSingle.
	VisitBoolExprSingle(ctx *BoolExprSingleContext) interface{}

	// Visit a parse tree produced by TinyLangParser#boolExprOperand.
	VisitBoolExprOperand(ctx *BoolExprOperandContext) interface{}

	// Visit a parse tree produced by TinyLangParser#breakOrContinue.
	VisitBreakOrContinue(ctx *BreakOrContinueContext) interface{}

	// Visit a parse tree produced by TinyLangParser#statementBody.
	VisitStatementBody(ctx *StatementBodyContext) interface{}

	// Visit a parse tree produced by TinyLangParser#ifElseSt.
	VisitIfElseSt(ctx *IfElseStContext) interface{}

	// Visit a parse tree produced by TinyLangParser#ifSt.
	VisitIfSt(ctx *IfStContext) interface{}

	// Visit a parse tree produced by TinyLangParser#elseIfSt.
	VisitElseIfSt(ctx *ElseIfStContext) interface{}

	// Visit a parse tree produced by TinyLangParser#elseSt.
	VisitElseSt(ctx *ElseStContext) interface{}

	// Visit a parse tree produced by TinyLangParser#whileSt.
	VisitWhileSt(ctx *WhileStContext) interface{}

	// Visit a parse tree produced by TinyLangParser#forSt.
	VisitForSt(ctx *ForStContext) interface{}
}
