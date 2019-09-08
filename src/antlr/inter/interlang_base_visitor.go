// Code generated from C:/projects/tiny-compiler/src/antlr\InterLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // InterLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseInterLangVisitor struct {
	*antlr.BaseParseTreeVisitor
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

func (v *BaseInterLangVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}
