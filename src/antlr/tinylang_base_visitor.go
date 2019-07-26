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

func (v *BaseTinyLangVisitor) VisitArgs(ctx *ArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitReturnSt(ctx *ReturnStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitFuncInvoc(ctx *FuncInvocContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitArgsInvoc(ctx *ArgsInvocContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitArgInvoc(ctx *ArgInvocContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitUpdVar(ctx *UpdVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitNewVar(ctx *NewVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitVarType(ctx *VarTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitArrayElem(ctx *ArrayElemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitArrayInit(ctx *ArrayInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitArrayInitEmp(ctx *ArrayInitEmpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitArrayInitVal(ctx *ArrayInitValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitArrayInitElems(ctx *ArrayInitElemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitExprOp(ctx *ExprOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitBoolExprOp(ctx *BoolExprOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitCommonBody(ctx *CommonBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitIfElse(ctx *IfElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitIF(ctx *IFContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitElseIf(ctx *ElseIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitElsePart(ctx *ElsePartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitWhile(ctx *WhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTinyLangVisitor) VisitForSt(ctx *ForStContext) interface{} {
	return v.VisitChildren(ctx)
}
