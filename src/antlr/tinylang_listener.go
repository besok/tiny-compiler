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

	// EnterFuncArgs is called when entering the funcArgs production.
	EnterFuncArgs(c *FuncArgsContext)

	// EnterFuncReturn is called when entering the funcReturn production.
	EnterFuncReturn(c *FuncReturnContext)

	// EnterFuncInvoc is called when entering the funcInvoc production.
	EnterFuncInvoc(c *FuncInvocContext)

	// EnterFuncArgsInvoc is called when entering the funcArgsInvoc production.
	EnterFuncArgsInvoc(c *FuncArgsInvocContext)

	// EnterFuncInvocArgs is called when entering the funcInvocArgs production.
	EnterFuncInvocArgs(c *FuncInvocArgsContext)

	// EnterUpdVariable is called when entering the updVariable production.
	EnterUpdVariable(c *UpdVariableContext)

	// EnterNewVariable is called when entering the newVariable production.
	EnterNewVariable(c *NewVariableContext)

	// EnterVariableType is called when entering the variableType production.
	EnterVariableType(c *VariableTypeContext)

	// EnterArrayElem is called when entering the arrayElem production.
	EnterArrayElem(c *ArrayElemContext)

	// EnterArrayInit is called when entering the arrayInit production.
	EnterArrayInit(c *ArrayInitContext)

	// EnterArrayInitEmpty is called when entering the arrayInitEmpty production.
	EnterArrayInitEmpty(c *ArrayInitEmptyContext)

	// EnterArrayInitValue is called when entering the arrayInitValue production.
	EnterArrayInitValue(c *ArrayInitValueContext)

	// EnterArrayInitElems is called when entering the arrayInitElems production.
	EnterArrayInitElems(c *ArrayInitElemsContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExprOperand is called when entering the exprOperand production.
	EnterExprOperand(c *ExprOperandContext)

	// EnterBoolExpr is called when entering the boolExpr production.
	EnterBoolExpr(c *BoolExprContext)

	// EnterBoolExprSingle is called when entering the boolExprSingle production.
	EnterBoolExprSingle(c *BoolExprSingleContext)

	// EnterBoolExprOperand is called when entering the boolExprOperand production.
	EnterBoolExprOperand(c *BoolExprOperandContext)

	// EnterStatementBody is called when entering the statementBody production.
	EnterStatementBody(c *StatementBodyContext)

	// EnterIfElseSt is called when entering the ifElseSt production.
	EnterIfElseSt(c *IfElseStContext)

	// EnterIfSt is called when entering the ifSt production.
	EnterIfSt(c *IfStContext)

	// EnterElseIfSt is called when entering the elseIfSt production.
	EnterElseIfSt(c *ElseIfStContext)

	// EnterElseSt is called when entering the elseSt production.
	EnterElseSt(c *ElseStContext)

	// EnterWhileSt is called when entering the whileSt production.
	EnterWhileSt(c *WhileStContext)

	// EnterForSt is called when entering the forSt production.
	EnterForSt(c *ForStContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitFuncInit is called when exiting the funcInit production.
	ExitFuncInit(c *FuncInitContext)

	// ExitFuncArgs is called when exiting the funcArgs production.
	ExitFuncArgs(c *FuncArgsContext)

	// ExitFuncReturn is called when exiting the funcReturn production.
	ExitFuncReturn(c *FuncReturnContext)

	// ExitFuncInvoc is called when exiting the funcInvoc production.
	ExitFuncInvoc(c *FuncInvocContext)

	// ExitFuncArgsInvoc is called when exiting the funcArgsInvoc production.
	ExitFuncArgsInvoc(c *FuncArgsInvocContext)

	// ExitFuncInvocArgs is called when exiting the funcInvocArgs production.
	ExitFuncInvocArgs(c *FuncInvocArgsContext)

	// ExitUpdVariable is called when exiting the updVariable production.
	ExitUpdVariable(c *UpdVariableContext)

	// ExitNewVariable is called when exiting the newVariable production.
	ExitNewVariable(c *NewVariableContext)

	// ExitVariableType is called when exiting the variableType production.
	ExitVariableType(c *VariableTypeContext)

	// ExitArrayElem is called when exiting the arrayElem production.
	ExitArrayElem(c *ArrayElemContext)

	// ExitArrayInit is called when exiting the arrayInit production.
	ExitArrayInit(c *ArrayInitContext)

	// ExitArrayInitEmpty is called when exiting the arrayInitEmpty production.
	ExitArrayInitEmpty(c *ArrayInitEmptyContext)

	// ExitArrayInitValue is called when exiting the arrayInitValue production.
	ExitArrayInitValue(c *ArrayInitValueContext)

	// ExitArrayInitElems is called when exiting the arrayInitElems production.
	ExitArrayInitElems(c *ArrayInitElemsContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExprOperand is called when exiting the exprOperand production.
	ExitExprOperand(c *ExprOperandContext)

	// ExitBoolExpr is called when exiting the boolExpr production.
	ExitBoolExpr(c *BoolExprContext)

	// ExitBoolExprSingle is called when exiting the boolExprSingle production.
	ExitBoolExprSingle(c *BoolExprSingleContext)

	// ExitBoolExprOperand is called when exiting the boolExprOperand production.
	ExitBoolExprOperand(c *BoolExprOperandContext)

	// ExitStatementBody is called when exiting the statementBody production.
	ExitStatementBody(c *StatementBodyContext)

	// ExitIfElseSt is called when exiting the ifElseSt production.
	ExitIfElseSt(c *IfElseStContext)

	// ExitIfSt is called when exiting the ifSt production.
	ExitIfSt(c *IfStContext)

	// ExitElseIfSt is called when exiting the elseIfSt production.
	ExitElseIfSt(c *ElseIfStContext)

	// ExitElseSt is called when exiting the elseSt production.
	ExitElseSt(c *ElseStContext)

	// ExitWhileSt is called when exiting the whileSt production.
	ExitWhileSt(c *WhileStContext)

	// ExitForSt is called when exiting the forSt production.
	ExitForSt(c *ForStContext)
}
