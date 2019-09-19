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

// EnterFile is called when production file is entered.
func (s *BaseInterLangListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseInterLangListener) ExitFile(ctx *FileContext) {}

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

// EnterStatement is called when production statement is entered.
func (s *BaseInterLangListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseInterLangListener) ExitStatement(ctx *StatementContext) {}

// EnterInternalArrArg is called when production internalArrArg is entered.
func (s *BaseInterLangListener) EnterInternalArrArg(ctx *InternalArrArgContext) {}

// ExitInternalArrArg is called when production internalArrArg is exited.
func (s *BaseInterLangListener) ExitInternalArrArg(ctx *InternalArrArgContext) {}

// EnterInternalVar is called when production internalVar is entered.
func (s *BaseInterLangListener) EnterInternalVar(ctx *InternalVarContext) {}

// ExitInternalVar is called when production internalVar is exited.
func (s *BaseInterLangListener) ExitInternalVar(ctx *InternalVarContext) {}

// EnterNewVar is called when production newVar is entered.
func (s *BaseInterLangListener) EnterNewVar(ctx *NewVarContext) {}

// ExitNewVar is called when production newVar is exited.
func (s *BaseInterLangListener) ExitNewVar(ctx *NewVarContext) {}

// EnterNewArrVar is called when production newArrVar is entered.
func (s *BaseInterLangListener) EnterNewArrVar(ctx *NewArrVarContext) {}

// ExitNewArrVar is called when production newArrVar is exited.
func (s *BaseInterLangListener) ExitNewArrVar(ctx *NewArrVarContext) {}

// EnterUpdVar is called when production updVar is entered.
func (s *BaseInterLangListener) EnterUpdVar(ctx *UpdVarContext) {}

// ExitUpdVar is called when production updVar is exited.
func (s *BaseInterLangListener) ExitUpdVar(ctx *UpdVarContext) {}

// EnterInitIntVar is called when production initIntVar is entered.
func (s *BaseInterLangListener) EnterInitIntVar(ctx *InitIntVarContext) {}

// ExitInitIntVar is called when production initIntVar is exited.
func (s *BaseInterLangListener) ExitInitIntVar(ctx *InitIntVarContext) {}

// EnterInitItem is called when production initItem is entered.
func (s *BaseInterLangListener) EnterInitItem(ctx *InitItemContext) {}

// ExitInitItem is called when production initItem is exited.
func (s *BaseInterLangListener) ExitInitItem(ctx *InitItemContext) {}

// EnterInitPrim is called when production initPrim is entered.
func (s *BaseInterLangListener) EnterInitPrim(ctx *InitPrimContext) {}

// ExitInitPrim is called when production initPrim is exited.
func (s *BaseInterLangListener) ExitInitPrim(ctx *InitPrimContext) {}

// EnterInitParam is called when production initParam is entered.
func (s *BaseInterLangListener) EnterInitParam(ctx *InitParamContext) {}

// ExitInitParam is called when production initParam is exited.
func (s *BaseInterLangListener) ExitInitParam(ctx *InitParamContext) {}

// EnterInitCall is called when production initCall is entered.
func (s *BaseInterLangListener) EnterInitCall(ctx *InitCallContext) {}

// ExitInitCall is called when production initCall is exited.
func (s *BaseInterLangListener) ExitInitCall(ctx *InitCallContext) {}

// EnterInitBoolOp is called when production initBoolOp is entered.
func (s *BaseInterLangListener) EnterInitBoolOp(ctx *InitBoolOpContext) {}

// ExitInitBoolOp is called when production initBoolOp is exited.
func (s *BaseInterLangListener) ExitInitBoolOp(ctx *InitBoolOpContext) {}

// EnterInitNumOp is called when production initNumOp is entered.
func (s *BaseInterLangListener) EnterInitNumOp(ctx *InitNumOpContext) {}

// ExitInitNumOp is called when production initNumOp is exited.
func (s *BaseInterLangListener) ExitInitNumOp(ctx *InitNumOpContext) {}

// EnterInitArrEl is called when production initArrEl is entered.
func (s *BaseInterLangListener) EnterInitArrEl(ctx *InitArrElContext) {}

// ExitInitArrEl is called when production initArrEl is exited.
func (s *BaseInterLangListener) ExitInitArrEl(ctx *InitArrElContext) {}

// EnterExtArrEl is called when production extArrEl is entered.
func (s *BaseInterLangListener) EnterExtArrEl(ctx *ExtArrElContext) {}

// ExitExtArrEl is called when production extArrEl is exited.
func (s *BaseInterLangListener) ExitExtArrEl(ctx *ExtArrElContext) {}

// EnterReturnTp is called when production returnTp is entered.
func (s *BaseInterLangListener) EnterReturnTp(ctx *ReturnTpContext) {}

// ExitReturnTp is called when production returnTp is exited.
func (s *BaseInterLangListener) ExitReturnTp(ctx *ReturnTpContext) {}

// EnterGotoIn is called when production gotoIn is entered.
func (s *BaseInterLangListener) EnterGotoIn(ctx *GotoInContext) {}

// ExitGotoIn is called when production gotoIn is exited.
func (s *BaseInterLangListener) ExitGotoIn(ctx *GotoInContext) {}

// EnterGotoTp is called when production gotoTp is entered.
func (s *BaseInterLangListener) EnterGotoTp(ctx *GotoTpContext) {}

// ExitGotoTp is called when production gotoTp is exited.
func (s *BaseInterLangListener) ExitGotoTp(ctx *GotoTpContext) {}

// EnterIfFalse is called when production ifFalse is entered.
func (s *BaseInterLangListener) EnterIfFalse(ctx *IfFalseContext) {}

// ExitIfFalse is called when production ifFalse is exited.
func (s *BaseInterLangListener) ExitIfFalse(ctx *IfFalseContext) {}

// EnterTypeTp is called when production typeTp is entered.
func (s *BaseInterLangListener) EnterTypeTp(ctx *TypeTpContext) {}

// ExitTypeTp is called when production typeTp is exited.
func (s *BaseInterLangListener) ExitTypeTp(ctx *TypeTpContext) {}

// EnterLine is called when production line is entered.
func (s *BaseInterLangListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseInterLangListener) ExitLine(ctx *LineContext) {}
