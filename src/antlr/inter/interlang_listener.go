// Code generated from C:/projects/tiny-compiler/src/antlr\InterLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // InterLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// InterLangListener is a complete listener for a parse tree produced by InterLangParser.
type InterLangListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterRetTp is called when entering the retTp production.
	EnterRetTp(c *RetTpContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterInternalArrArg is called when entering the internalArrArg production.
	EnterInternalArrArg(c *InternalArrArgContext)

	// EnterInternalVar is called when entering the internalVar production.
	EnterInternalVar(c *InternalVarContext)

	// EnterNewVar is called when entering the newVar production.
	EnterNewVar(c *NewVarContext)

	// EnterNewArrVar is called when entering the newArrVar production.
	EnterNewArrVar(c *NewArrVarContext)

	// EnterUpdVar is called when entering the updVar production.
	EnterUpdVar(c *UpdVarContext)

	// EnterInitItem is called when entering the initItem production.
	EnterInitItem(c *InitItemContext)

	// EnterInitPrim is called when entering the initPrim production.
	EnterInitPrim(c *InitPrimContext)

	// EnterInitParam is called when entering the initParam production.
	EnterInitParam(c *InitParamContext)

	// EnterInitCall is called when entering the initCall production.
	EnterInitCall(c *InitCallContext)

	// EnterInitBoolOp is called when entering the initBoolOp production.
	EnterInitBoolOp(c *InitBoolOpContext)

	// EnterInitNumOp is called when entering the initNumOp production.
	EnterInitNumOp(c *InitNumOpContext)

	// EnterInitArrEl is called when entering the initArrEl production.
	EnterInitArrEl(c *InitArrElContext)

	// EnterExtArrEl is called when entering the extArrEl production.
	EnterExtArrEl(c *ExtArrElContext)

	// EnterReturn is called when entering the return production.
	EnterReturn(c *ReturnContext)

	// EnterGotoIn is called when entering the gotoIn production.
	EnterGotoIn(c *GotoInContext)

	// EnterGoto is called when entering the goto production.
	EnterGoto(c *GotoContext)

	// EnterIfFalse is called when entering the ifFalse production.
	EnterIfFalse(c *IfFalseContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitRetTp is called when exiting the retTp production.
	ExitRetTp(c *RetTpContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitInternalArrArg is called when exiting the internalArrArg production.
	ExitInternalArrArg(c *InternalArrArgContext)

	// ExitInternalVar is called when exiting the internalVar production.
	ExitInternalVar(c *InternalVarContext)

	// ExitNewVar is called when exiting the newVar production.
	ExitNewVar(c *NewVarContext)

	// ExitNewArrVar is called when exiting the newArrVar production.
	ExitNewArrVar(c *NewArrVarContext)

	// ExitUpdVar is called when exiting the updVar production.
	ExitUpdVar(c *UpdVarContext)

	// ExitInitItem is called when exiting the initItem production.
	ExitInitItem(c *InitItemContext)

	// ExitInitPrim is called when exiting the initPrim production.
	ExitInitPrim(c *InitPrimContext)

	// ExitInitParam is called when exiting the initParam production.
	ExitInitParam(c *InitParamContext)

	// ExitInitCall is called when exiting the initCall production.
	ExitInitCall(c *InitCallContext)

	// ExitInitBoolOp is called when exiting the initBoolOp production.
	ExitInitBoolOp(c *InitBoolOpContext)

	// ExitInitNumOp is called when exiting the initNumOp production.
	ExitInitNumOp(c *InitNumOpContext)

	// ExitInitArrEl is called when exiting the initArrEl production.
	ExitInitArrEl(c *InitArrElContext)

	// ExitExtArrEl is called when exiting the extArrEl production.
	ExitExtArrEl(c *ExtArrElContext)

	// ExitReturn is called when exiting the return production.
	ExitReturn(c *ReturnContext)

	// ExitGotoIn is called when exiting the gotoIn production.
	ExitGotoIn(c *GotoInContext)

	// ExitGoto is called when exiting the goto production.
	ExitGoto(c *GotoContext)

	// ExitIfFalse is called when exiting the ifFalse production.
	ExitIfFalse(c *IfFalseContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)
}
