// Code generated from C:/projects/tiny-compiler/src/antlr\TinyLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // TinyLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TinyLangListener is a complete listener for a parse tree produced by TinyLangParser.
type TinyLangListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterFuncInit is called when entering the funcInit production.
	EnterFuncInit(c *FuncInitContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterReturnSt is called when entering the returnSt production.
	EnterReturnSt(c *ReturnStContext)

	// EnterFuncInvoc is called when entering the funcInvoc production.
	EnterFuncInvoc(c *FuncInvocContext)

	// EnterArgsInvoc is called when entering the argsInvoc production.
	EnterArgsInvoc(c *ArgsInvocContext)

	// EnterArgInvoc is called when entering the argInvoc production.
	EnterArgInvoc(c *ArgInvocContext)

	// EnterUpdVar is called when entering the updVar production.
	EnterUpdVar(c *UpdVarContext)

	// EnterNewVar is called when entering the newVar production.
	EnterNewVar(c *NewVarContext)

	// EnterVarType is called when entering the varType production.
	EnterVarType(c *VarTypeContext)

	// EnterArrayElem is called when entering the arrayElem production.
	EnterArrayElem(c *ArrayElemContext)

	// EnterArrayInit is called when entering the arrayInit production.
	EnterArrayInit(c *ArrayInitContext)

	// EnterArrayInitEmp is called when entering the arrayInitEmp production.
	EnterArrayInitEmp(c *ArrayInitEmpContext)

	// EnterArrayInitVal is called when entering the arrayInitVal production.
	EnterArrayInitVal(c *ArrayInitValContext)

	// EnterArrayInitElems is called when entering the arrayInitElems production.
	EnterArrayInitElems(c *ArrayInitElemsContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExprOp is called when entering the exprOp production.
	EnterExprOp(c *ExprOpContext)

	// EnterBoolExpr is called when entering the boolExpr production.
	EnterBoolExpr(c *BoolExprContext)

	// EnterBoolExprOp is called when entering the boolExprOp production.
	EnterBoolExprOp(c *BoolExprOpContext)

	// EnterCommonBody is called when entering the commonBody production.
	EnterCommonBody(c *CommonBodyContext)

	// EnterIfElse is called when entering the ifElse production.
	EnterIfElse(c *IfElseContext)

	// EnterIF is called when entering the iF production.
	EnterIF(c *IFContext)

	// EnterElseIf is called when entering the elseIf production.
	EnterElseIf(c *ElseIfContext)

	// EnterElsePart is called when entering the elsePart production.
	EnterElsePart(c *ElsePartContext)

	// EnterWhile is called when entering the while production.
	EnterWhile(c *WhileContext)

	// EnterForSt is called when entering the forSt production.
	EnterForSt(c *ForStContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitFuncInit is called when exiting the funcInit production.
	ExitFuncInit(c *FuncInitContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitReturnSt is called when exiting the returnSt production.
	ExitReturnSt(c *ReturnStContext)

	// ExitFuncInvoc is called when exiting the funcInvoc production.
	ExitFuncInvoc(c *FuncInvocContext)

	// ExitArgsInvoc is called when exiting the argsInvoc production.
	ExitArgsInvoc(c *ArgsInvocContext)

	// ExitArgInvoc is called when exiting the argInvoc production.
	ExitArgInvoc(c *ArgInvocContext)

	// ExitUpdVar is called when exiting the updVar production.
	ExitUpdVar(c *UpdVarContext)

	// ExitNewVar is called when exiting the newVar production.
	ExitNewVar(c *NewVarContext)

	// ExitVarType is called when exiting the varType production.
	ExitVarType(c *VarTypeContext)

	// ExitArrayElem is called when exiting the arrayElem production.
	ExitArrayElem(c *ArrayElemContext)

	// ExitArrayInit is called when exiting the arrayInit production.
	ExitArrayInit(c *ArrayInitContext)

	// ExitArrayInitEmp is called when exiting the arrayInitEmp production.
	ExitArrayInitEmp(c *ArrayInitEmpContext)

	// ExitArrayInitVal is called when exiting the arrayInitVal production.
	ExitArrayInitVal(c *ArrayInitValContext)

	// ExitArrayInitElems is called when exiting the arrayInitElems production.
	ExitArrayInitElems(c *ArrayInitElemsContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExprOp is called when exiting the exprOp production.
	ExitExprOp(c *ExprOpContext)

	// ExitBoolExpr is called when exiting the boolExpr production.
	ExitBoolExpr(c *BoolExprContext)

	// ExitBoolExprOp is called when exiting the boolExprOp production.
	ExitBoolExprOp(c *BoolExprOpContext)

	// ExitCommonBody is called when exiting the commonBody production.
	ExitCommonBody(c *CommonBodyContext)

	// ExitIfElse is called when exiting the ifElse production.
	ExitIfElse(c *IfElseContext)

	// ExitIF is called when exiting the iF production.
	ExitIF(c *IFContext)

	// ExitElseIf is called when exiting the elseIf production.
	ExitElseIf(c *ElseIfContext)

	// ExitElsePart is called when exiting the elsePart production.
	ExitElsePart(c *ElsePartContext)

	// ExitWhile is called when exiting the while production.
	ExitWhile(c *WhileContext)

	// ExitForSt is called when exiting the forSt production.
	ExitForSt(c *ForStContext)
}
