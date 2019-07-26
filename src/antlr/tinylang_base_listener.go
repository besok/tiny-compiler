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

// EnterArgs is called when production args is entered.
func (s *BaseTinyLangListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseTinyLangListener) ExitArgs(ctx *ArgsContext) {}

// EnterReturnSt is called when production returnSt is entered.
func (s *BaseTinyLangListener) EnterReturnSt(ctx *ReturnStContext) {}

// ExitReturnSt is called when production returnSt is exited.
func (s *BaseTinyLangListener) ExitReturnSt(ctx *ReturnStContext) {}

// EnterFuncInvoc is called when production funcInvoc is entered.
func (s *BaseTinyLangListener) EnterFuncInvoc(ctx *FuncInvocContext) {}

// ExitFuncInvoc is called when production funcInvoc is exited.
func (s *BaseTinyLangListener) ExitFuncInvoc(ctx *FuncInvocContext) {}

// EnterArgsInvoc is called when production argsInvoc is entered.
func (s *BaseTinyLangListener) EnterArgsInvoc(ctx *ArgsInvocContext) {}

// ExitArgsInvoc is called when production argsInvoc is exited.
func (s *BaseTinyLangListener) ExitArgsInvoc(ctx *ArgsInvocContext) {}

// EnterArgInvoc is called when production argInvoc is entered.
func (s *BaseTinyLangListener) EnterArgInvoc(ctx *ArgInvocContext) {}

// ExitArgInvoc is called when production argInvoc is exited.
func (s *BaseTinyLangListener) ExitArgInvoc(ctx *ArgInvocContext) {}

// EnterUpdVar is called when production updVar is entered.
func (s *BaseTinyLangListener) EnterUpdVar(ctx *UpdVarContext) {}

// ExitUpdVar is called when production updVar is exited.
func (s *BaseTinyLangListener) ExitUpdVar(ctx *UpdVarContext) {}

// EnterNewVar is called when production newVar is entered.
func (s *BaseTinyLangListener) EnterNewVar(ctx *NewVarContext) {}

// ExitNewVar is called when production newVar is exited.
func (s *BaseTinyLangListener) ExitNewVar(ctx *NewVarContext) {}

// EnterVarType is called when production varType is entered.
func (s *BaseTinyLangListener) EnterVarType(ctx *VarTypeContext) {}

// ExitVarType is called when production varType is exited.
func (s *BaseTinyLangListener) ExitVarType(ctx *VarTypeContext) {}

// EnterArrayElem is called when production arrayElem is entered.
func (s *BaseTinyLangListener) EnterArrayElem(ctx *ArrayElemContext) {}

// ExitArrayElem is called when production arrayElem is exited.
func (s *BaseTinyLangListener) ExitArrayElem(ctx *ArrayElemContext) {}

// EnterArrayInit is called when production arrayInit is entered.
func (s *BaseTinyLangListener) EnterArrayInit(ctx *ArrayInitContext) {}

// ExitArrayInit is called when production arrayInit is exited.
func (s *BaseTinyLangListener) ExitArrayInit(ctx *ArrayInitContext) {}

// EnterArrayInitEmp is called when production arrayInitEmp is entered.
func (s *BaseTinyLangListener) EnterArrayInitEmp(ctx *ArrayInitEmpContext) {}

// ExitArrayInitEmp is called when production arrayInitEmp is exited.
func (s *BaseTinyLangListener) ExitArrayInitEmp(ctx *ArrayInitEmpContext) {}

// EnterArrayInitVal is called when production arrayInitVal is entered.
func (s *BaseTinyLangListener) EnterArrayInitVal(ctx *ArrayInitValContext) {}

// ExitArrayInitVal is called when production arrayInitVal is exited.
func (s *BaseTinyLangListener) ExitArrayInitVal(ctx *ArrayInitValContext) {}

// EnterArrayInitElems is called when production arrayInitElems is entered.
func (s *BaseTinyLangListener) EnterArrayInitElems(ctx *ArrayInitElemsContext) {}

// ExitArrayInitElems is called when production arrayInitElems is exited.
func (s *BaseTinyLangListener) ExitArrayInitElems(ctx *ArrayInitElemsContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseTinyLangListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseTinyLangListener) ExitExpr(ctx *ExprContext) {}

// EnterExprOp is called when production exprOp is entered.
func (s *BaseTinyLangListener) EnterExprOp(ctx *ExprOpContext) {}

// ExitExprOp is called when production exprOp is exited.
func (s *BaseTinyLangListener) ExitExprOp(ctx *ExprOpContext) {}

// EnterBoolExpr is called when production boolExpr is entered.
func (s *BaseTinyLangListener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production boolExpr is exited.
func (s *BaseTinyLangListener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterBoolExprOp is called when production boolExprOp is entered.
func (s *BaseTinyLangListener) EnterBoolExprOp(ctx *BoolExprOpContext) {}

// ExitBoolExprOp is called when production boolExprOp is exited.
func (s *BaseTinyLangListener) ExitBoolExprOp(ctx *BoolExprOpContext) {}

// EnterCommonBody is called when production commonBody is entered.
func (s *BaseTinyLangListener) EnterCommonBody(ctx *CommonBodyContext) {}

// ExitCommonBody is called when production commonBody is exited.
func (s *BaseTinyLangListener) ExitCommonBody(ctx *CommonBodyContext) {}

// EnterIfElse is called when production ifElse is entered.
func (s *BaseTinyLangListener) EnterIfElse(ctx *IfElseContext) {}

// ExitIfElse is called when production ifElse is exited.
func (s *BaseTinyLangListener) ExitIfElse(ctx *IfElseContext) {}

// EnterIF is called when production iF is entered.
func (s *BaseTinyLangListener) EnterIF(ctx *IFContext) {}

// ExitIF is called when production iF is exited.
func (s *BaseTinyLangListener) ExitIF(ctx *IFContext) {}

// EnterElseIf is called when production elseIf is entered.
func (s *BaseTinyLangListener) EnterElseIf(ctx *ElseIfContext) {}

// ExitElseIf is called when production elseIf is exited.
func (s *BaseTinyLangListener) ExitElseIf(ctx *ElseIfContext) {}

// EnterElsePart is called when production elsePart is entered.
func (s *BaseTinyLangListener) EnterElsePart(ctx *ElsePartContext) {}

// ExitElsePart is called when production elsePart is exited.
func (s *BaseTinyLangListener) ExitElsePart(ctx *ElsePartContext) {}

// EnterWhile is called when production while is entered.
func (s *BaseTinyLangListener) EnterWhile(ctx *WhileContext) {}

// ExitWhile is called when production while is exited.
func (s *BaseTinyLangListener) ExitWhile(ctx *WhileContext) {}

// EnterForSt is called when production forSt is entered.
func (s *BaseTinyLangListener) EnterForSt(ctx *ForStContext) {}

// ExitForSt is called when production forSt is exited.
func (s *BaseTinyLangListener) ExitForSt(ctx *ForStContext) {}
