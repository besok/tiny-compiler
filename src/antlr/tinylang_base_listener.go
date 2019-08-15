// Code generated from C:/projects/tiny-compiler/src/antlr\TinyLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // TinyLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTinyLangListener is a complete listener for a parse tree produced by TinyLangParser.
type BaseTinyLangListener struct{}

var _ TinyLangListener = &BaseTinyLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTinyLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTinyLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTinyLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTinyLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseTinyLangListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseTinyLangListener) ExitFile(ctx *FileContext) {}

// EnterFuncInit is called when production funcInit is entered.
func (s *BaseTinyLangListener) EnterFuncInit(ctx *FuncInitContext) {}

// ExitFuncInit is called when production funcInit is exited.
func (s *BaseTinyLangListener) ExitFuncInit(ctx *FuncInitContext) {}

// EnterFuncArg is called when production funcArg is entered.
func (s *BaseTinyLangListener) EnterFuncArg(ctx *FuncArgContext) {}

// ExitFuncArg is called when production funcArg is exited.
func (s *BaseTinyLangListener) ExitFuncArg(ctx *FuncArgContext) {}

// EnterFuncArgs is called when production funcArgs is entered.
func (s *BaseTinyLangListener) EnterFuncArgs(ctx *FuncArgsContext) {}

// ExitFuncArgs is called when production funcArgs is exited.
func (s *BaseTinyLangListener) ExitFuncArgs(ctx *FuncArgsContext) {}

// EnterFuncReturn is called when production funcReturn is entered.
func (s *BaseTinyLangListener) EnterFuncReturn(ctx *FuncReturnContext) {}

// ExitFuncReturn is called when production funcReturn is exited.
func (s *BaseTinyLangListener) ExitFuncReturn(ctx *FuncReturnContext) {}

// EnterFuncReturnType is called when production funcReturnType is entered.
func (s *BaseTinyLangListener) EnterFuncReturnType(ctx *FuncReturnTypeContext) {}

// ExitFuncReturnType is called when production funcReturnType is exited.
func (s *BaseTinyLangListener) ExitFuncReturnType(ctx *FuncReturnTypeContext) {}

// EnterFuncInvoc is called when production funcInvoc is entered.
func (s *BaseTinyLangListener) EnterFuncInvoc(ctx *FuncInvocContext) {}

// ExitFuncInvoc is called when production funcInvoc is exited.
func (s *BaseTinyLangListener) ExitFuncInvoc(ctx *FuncInvocContext) {}

// EnterFuncArgsInvoc is called when production funcArgsInvoc is entered.
func (s *BaseTinyLangListener) EnterFuncArgsInvoc(ctx *FuncArgsInvocContext) {}

// ExitFuncArgsInvoc is called when production funcArgsInvoc is exited.
func (s *BaseTinyLangListener) ExitFuncArgsInvoc(ctx *FuncArgsInvocContext) {}

// EnterFuncInvocArgs is called when production funcInvocArgs is entered.
func (s *BaseTinyLangListener) EnterFuncInvocArgs(ctx *FuncInvocArgsContext) {}

// ExitFuncInvocArgs is called when production funcInvocArgs is exited.
func (s *BaseTinyLangListener) ExitFuncInvocArgs(ctx *FuncInvocArgsContext) {}

// EnterUpdVariable is called when production updVariable is entered.
func (s *BaseTinyLangListener) EnterUpdVariable(ctx *UpdVariableContext) {}

// ExitUpdVariable is called when production updVariable is exited.
func (s *BaseTinyLangListener) ExitUpdVariable(ctx *UpdVariableContext) {}

// EnterNewVariable is called when production newVariable is entered.
func (s *BaseTinyLangListener) EnterNewVariable(ctx *NewVariableContext) {}

// ExitNewVariable is called when production newVariable is exited.
func (s *BaseTinyLangListener) ExitNewVariable(ctx *NewVariableContext) {}

// EnterVal is called when production val is entered.
func (s *BaseTinyLangListener) EnterVal(ctx *ValContext) {}

// ExitVal is called when production val is exited.
func (s *BaseTinyLangListener) ExitVal(ctx *ValContext) {}

// EnterVariableType is called when production variableType is entered.
func (s *BaseTinyLangListener) EnterVariableType(ctx *VariableTypeContext) {}

// ExitVariableType is called when production variableType is exited.
func (s *BaseTinyLangListener) ExitVariableType(ctx *VariableTypeContext) {}

// EnterArrayElem is called when production arrayElem is entered.
func (s *BaseTinyLangListener) EnterArrayElem(ctx *ArrayElemContext) {}

// ExitArrayElem is called when production arrayElem is exited.
func (s *BaseTinyLangListener) ExitArrayElem(ctx *ArrayElemContext) {}

// EnterArrayInit is called when production arrayInit is entered.
func (s *BaseTinyLangListener) EnterArrayInit(ctx *ArrayInitContext) {}

// ExitArrayInit is called when production arrayInit is exited.
func (s *BaseTinyLangListener) ExitArrayInit(ctx *ArrayInitContext) {}

// EnterArrayInitEmpty is called when production arrayInitEmpty is entered.
func (s *BaseTinyLangListener) EnterArrayInitEmpty(ctx *ArrayInitEmptyContext) {}

// ExitArrayInitEmpty is called when production arrayInitEmpty is exited.
func (s *BaseTinyLangListener) ExitArrayInitEmpty(ctx *ArrayInitEmptyContext) {}

// EnterArrayInitValue is called when production arrayInitValue is entered.
func (s *BaseTinyLangListener) EnterArrayInitValue(ctx *ArrayInitValueContext) {}

// ExitArrayInitValue is called when production arrayInitValue is exited.
func (s *BaseTinyLangListener) ExitArrayInitValue(ctx *ArrayInitValueContext) {}

// EnterArrayInitElems is called when production arrayInitElems is entered.
func (s *BaseTinyLangListener) EnterArrayInitElems(ctx *ArrayInitElemsContext) {}

// ExitArrayInitElems is called when production arrayInitElems is exited.
func (s *BaseTinyLangListener) ExitArrayInitElems(ctx *ArrayInitElemsContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseTinyLangListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseTinyLangListener) ExitExpr(ctx *ExprContext) {}

// EnterExprOperand is called when production exprOperand is entered.
func (s *BaseTinyLangListener) EnterExprOperand(ctx *ExprOperandContext) {}

// ExitExprOperand is called when production exprOperand is exited.
func (s *BaseTinyLangListener) ExitExprOperand(ctx *ExprOperandContext) {}

// EnterBoolExpr is called when production boolExpr is entered.
func (s *BaseTinyLangListener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production boolExpr is exited.
func (s *BaseTinyLangListener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterBoolExprSingle is called when production boolExprSingle is entered.
func (s *BaseTinyLangListener) EnterBoolExprSingle(ctx *BoolExprSingleContext) {}

// ExitBoolExprSingle is called when production boolExprSingle is exited.
func (s *BaseTinyLangListener) ExitBoolExprSingle(ctx *BoolExprSingleContext) {}

// EnterBoolExprOperand is called when production boolExprOperand is entered.
func (s *BaseTinyLangListener) EnterBoolExprOperand(ctx *BoolExprOperandContext) {}

// ExitBoolExprOperand is called when production boolExprOperand is exited.
func (s *BaseTinyLangListener) ExitBoolExprOperand(ctx *BoolExprOperandContext) {}

// EnterStatementBody is called when production statementBody is entered.
func (s *BaseTinyLangListener) EnterStatementBody(ctx *StatementBodyContext) {}

// ExitStatementBody is called when production statementBody is exited.
func (s *BaseTinyLangListener) ExitStatementBody(ctx *StatementBodyContext) {}

// EnterIfElseSt is called when production ifElseSt is entered.
func (s *BaseTinyLangListener) EnterIfElseSt(ctx *IfElseStContext) {}

// ExitIfElseSt is called when production ifElseSt is exited.
func (s *BaseTinyLangListener) ExitIfElseSt(ctx *IfElseStContext) {}

// EnterIfSt is called when production ifSt is entered.
func (s *BaseTinyLangListener) EnterIfSt(ctx *IfStContext) {}

// ExitIfSt is called when production ifSt is exited.
func (s *BaseTinyLangListener) ExitIfSt(ctx *IfStContext) {}

// EnterElseIfSt is called when production elseIfSt is entered.
func (s *BaseTinyLangListener) EnterElseIfSt(ctx *ElseIfStContext) {}

// ExitElseIfSt is called when production elseIfSt is exited.
func (s *BaseTinyLangListener) ExitElseIfSt(ctx *ElseIfStContext) {}

// EnterElseSt is called when production elseSt is entered.
func (s *BaseTinyLangListener) EnterElseSt(ctx *ElseStContext) {}

// ExitElseSt is called when production elseSt is exited.
func (s *BaseTinyLangListener) ExitElseSt(ctx *ElseStContext) {}

// EnterWhileSt is called when production whileSt is entered.
func (s *BaseTinyLangListener) EnterWhileSt(ctx *WhileStContext) {}

// ExitWhileSt is called when production whileSt is exited.
func (s *BaseTinyLangListener) ExitWhileSt(ctx *WhileStContext) {}

// EnterForSt is called when production forSt is entered.
func (s *BaseTinyLangListener) EnterForSt(ctx *ForStContext) {}

// ExitForSt is called when production forSt is exited.
func (s *BaseTinyLangListener) ExitForSt(ctx *ForStContext) {}
