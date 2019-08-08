package parsing

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"tiny-compiler/src/antlr"
)

type CommonTinyListener struct {
	*parser.BaseTinyLangListener
}

func NewListener() *CommonTinyListener {
	return &CommonTinyListener{}
}

func (l *CommonTinyListener) VisitTerminal(node antlr.TerminalNode) {
	panic("implement me")
}

func (l *CommonTinyListener) VisitErrorNode(node antlr.ErrorNode) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterFile(c *parser.FileContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterFuncInit(c *parser.FuncInitContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterFuncArgs(c *parser.FuncArgsContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterFuncReturn(c *parser.FuncReturnContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterFuncInvoc(c *parser.FuncInvocContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterFuncArgsInvoc(c *parser.FuncArgsInvocContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterFuncInvocArgs(c *parser.FuncInvocArgsContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterUpdVariable(c *parser.UpdVariableContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterNewVariable(c *parser.NewVariableContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterVariableType(c *parser.VariableTypeContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterArrayElem(c *parser.ArrayElemContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterArrayInit(c *parser.ArrayInitContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterArrayInitEmpty(c *parser.ArrayInitEmptyContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterArrayInitValue(c *parser.ArrayInitValueContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterArrayInitElems(c *parser.ArrayInitElemsContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterExpr(c *parser.ExprContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterExprOperand(c *parser.ExprOperandContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterBoolExpr(c *parser.BoolExprContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterBoolExprSingle(c *parser.BoolExprSingleContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterBoolExprOperand(c *parser.BoolExprOperandContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterStatementBody(c *parser.StatementBodyContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterIfElseSt(c *parser.IfElseStContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterIfSt(c *parser.IfStContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterElseIfSt(c *parser.ElseIfStContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterElseSt(c *parser.ElseStContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterWhileSt(c *parser.WhileStContext) {
	panic("implement me")
}

func (l *CommonTinyListener) EnterForSt(c *parser.ForStContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitFile(c *parser.FileContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitFuncInit(c *parser.FuncInitContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitFuncArgs(c *parser.FuncArgsContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitFuncReturn(c *parser.FuncReturnContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitFuncInvoc(c *parser.FuncInvocContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitFuncArgsInvoc(c *parser.FuncArgsInvocContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitFuncInvocArgs(c *parser.FuncInvocArgsContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitUpdVariable(c *parser.UpdVariableContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitNewVariable(c *parser.NewVariableContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitVariableType(c *parser.VariableTypeContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitArrayElem(c *parser.ArrayElemContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitArrayInit(c *parser.ArrayInitContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitArrayInitEmpty(c *parser.ArrayInitEmptyContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitArrayInitValue(c *parser.ArrayInitValueContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitArrayInitElems(c *parser.ArrayInitElemsContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitExpr(c *parser.ExprContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitExprOperand(c *parser.ExprOperandContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitBoolExpr(c *parser.BoolExprContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitBoolExprSingle(c *parser.BoolExprSingleContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitBoolExprOperand(c *parser.BoolExprOperandContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitStatementBody(c *parser.StatementBodyContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitIfElseSt(c *parser.IfElseStContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitIfSt(c *parser.IfStContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitElseIfSt(c *parser.ElseIfStContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitElseSt(c *parser.ElseStContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitWhileSt(c *parser.WhileStContext) {
	panic("implement me")
}

func (l *CommonTinyListener) ExitForSt(c *parser.ForStContext) {
	panic("implement me")
}




