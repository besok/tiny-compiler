// Code generated from C:/projects/tiny-compiler/src/antlr\InterLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // InterLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseInterLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseInterLangVisitor) VisitFile(ctx *FileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitArg(ctx *ArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitRetTp(ctx *RetTpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitInternalArrArg(ctx *InternalArrArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitInternalVar(ctx *InternalVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitNewVar(ctx *NewVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitUpdVar(ctx *UpdVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitInitIntVar(ctx *InitIntVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitInitItem(ctx *InitItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitInitPrim(ctx *InitPrimContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitInitParam(ctx *InitParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitInitCall(ctx *InitCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitInitBoolOp(ctx *InitBoolOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitInitNumOp(ctx *InitNumOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitInitArrEl(ctx *InitArrElContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitExtArrEl(ctx *ExtArrElContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitReturnTp(ctx *ReturnTpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitGotoIn(ctx *GotoInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitGotoTp(ctx *GotoTpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitIfFalse(ctx *IfFalseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitTypeTp(ctx *TypeTpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInterLangVisitor) VisitLine(ctx *LineContext) interface{} {
	return v.VisitChildren(ctx)
}
