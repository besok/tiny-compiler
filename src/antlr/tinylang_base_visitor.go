// Code generated from C:/projects/tiny-compiler/src/antlr\TinyLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // TinyLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseTinyLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTinyLangVisitor) VisitFile(ctx *FileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitFuncInit(ctx *FuncInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitFuncArg(ctx *FuncArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitFuncArgs(ctx *FuncArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitFuncReturn(ctx *FuncReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitFuncReturnType(ctx *FuncReturnTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitFuncInvoc(ctx *FuncInvocContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitFuncArgsInvoc(ctx *FuncArgsInvocContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitFuncInvocArgs(ctx *FuncInvocArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitUpdVariable(ctx *UpdVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitNewVariable(ctx *NewVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitVariableType(ctx *VariableTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitArrayElem(ctx *ArrayElemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitArrayInit(ctx *ArrayInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitArrayInitEmpty(ctx *ArrayInitEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitArrayInitValue(ctx *ArrayInitValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitArrayInitElems(ctx *ArrayInitElemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitExprOperand(ctx *ExprOperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitBoolExprSingle(ctx *BoolExprSingleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitBoolExprOperand(ctx *BoolExprOperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitStatementBody(ctx *StatementBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitIfElseSt(ctx *IfElseStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitIfSt(ctx *IfStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitElseIfSt(ctx *ElseIfStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitElseSt(ctx *ElseStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitWhileSt(ctx *WhileStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitForSt(ctx *ForStContext) interface{} {
	return v.VisitChildren(ctx)
}
