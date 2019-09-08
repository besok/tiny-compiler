// Code generated from C:/projects/tiny-compiler/src/antlr\InterLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // InterLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by InterLangParser.
type InterLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by InterLangParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by InterLangParser#arg.
	VisitArg(ctx *ArgContext) interface{}

	// Visit a parse tree produced by InterLangParser#retTp.
	VisitRetTp(ctx *RetTpContext) interface{}

	// Visit a parse tree produced by InterLangParser#type.
	VisitType(ctx *TypeContext) interface{}
}
