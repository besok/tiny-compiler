// Code generated from C:/projects/tiny-compiler/src/antlr\InterLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // InterLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// InterLangListener is a complete listener for a parse tree produced by InterLangParser.
type InterLangListener interface {
	antlr.ParseTreeListener

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterRetTp is called when entering the retTp production.
	EnterRetTp(c *RetTpContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitRetTp is called when exiting the retTp production.
	ExitRetTp(c *RetTpContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)
}
