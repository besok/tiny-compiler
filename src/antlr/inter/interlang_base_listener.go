// Code generated from C:/projects/tiny-compiler/src/antlr\InterLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // InterLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseInterLangListener is a complete listener for a parse tree produced by InterLangParser.
type BaseInterLangListener struct{}

var _ InterLangListener = &BaseInterLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseInterLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseInterLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseInterLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseInterLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseInterLangListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseInterLangListener) ExitFunction(ctx *FunctionContext) {}

// EnterArg is called when production arg is entered.
func (s *BaseInterLangListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BaseInterLangListener) ExitArg(ctx *ArgContext) {}

// EnterRetTp is called when production retTp is entered.
func (s *BaseInterLangListener) EnterRetTp(ctx *RetTpContext) {}

// ExitRetTp is called when production retTp is exited.
func (s *BaseInterLangListener) ExitRetTp(ctx *RetTpContext) {}

// EnterType is called when production type is entered.
func (s *BaseInterLangListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseInterLangListener) ExitType(ctx *TypeContext) {}
