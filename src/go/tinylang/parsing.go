package parsing

import (
	"tiny-compiler/src/antlr"
)

type PrintTinyListener struct {
	*parser.BaseTinyLangListener
}

//func (*PrintTinyListener) VisitTerminal(node antlr.TerminalNode) {
//	log.Println("visit terminal:", node.GetText())
//}
//
//func (*PrintTinyListener) VisitErrorNode(node antlr.ErrorNode) {
//	log.Println("visit errorNode:", node.GetText())
//}

//func (*PrintTinyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
//	log.Printf("enter rule:%s, text:%s \n", ctx.GetRuleContext().GetText(),ctx.GetText())
//}

//func (*PrintTinyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
//	panic("implement me")
//}

//func (*PrintTinyListener) EnterFile(c *parser.FileContext) {
//	panic("implement me")
//}

//func (*PrintTinyListener) EnterFuncInit(c *parser.FuncInitContext) {
//	panic("implement me")
//}
//
//variableType
//
//func (*PrintTinyListener) EnterFuncReturn(c *parser.FuncReturnContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterFuncInvoc(c *parser.FuncInvocContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterFuncArgsInvoc(c *parser.FuncArgsInvocContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterFuncInvocArgs(c *parser.FuncInvocArgsContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterUpdVariable(c *parser.UpdVariableContext) {
//	log.Printf(">> upd variable:%s",c.GetText())
//}

//func (*PrintTinyListener) ExitUpdVariable(c *parser.UpdVariableContext) {
//}
//
//func (*PrintTinyListener) EnterNewVariable(c *parser.NewVariableContext) {
//	panic("implement me")
//}
//
func (*PrintTinyListener) EnterVariableType(c *parser.VariableTypeContext) {
}
func (*PrintTinyListener) ExitVariableType(c *parser.VariableTypeContext) {
}
//
//func (*PrintTinyListener) EnterArrayElem(c *parser.ArrayElemContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterArrayInit(c *parser.ArrayInitContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterArrayInitEmpty(c *parser.ArrayInitEmptyContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterArrayInitValue(c *parser.ArrayInitValueContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterArrayInitElems(c *parser.ArrayInitElemsContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterExpr(c *parser.ExprContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterExprOperand(c *parser.ExprOperandContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterBoolExpr(c *parser.BoolExprContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterBoolExprOperand(c *parser.BoolExprOperandContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterStatementBody(c *parser.StatementBodyContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterIfElseSt(c *parser.IfElseStContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterIfSt(c *parser.IfStContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterElseIfSt(c *parser.ElseIfStContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterElseSt(c *parser.ElseStContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterWhileSt(c *parser.WhileStContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) EnterForSt(c *parser.ForStContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitFile(c *parser.FileContext) {
//	panic("implement me")
//}
//
func (*PrintTinyListener) ExitFuncInit(c *parser.FuncInitContext) {
	item := c.ITEM().GetText()
	line := c.GetStart().GetLine()
	_ = addFunc(FuncDefinition{line, item})
}
//
//func (*PrintTinyListener) ExitFuncArgs(c *parser.FuncArgsContext) {
//	panic("implement me")
//}

//func (*PrintTinyListener) ExitFuncReturn(c *parser.FuncReturnContext) {
//	log.Printf("exit func return: %s, ",c.GetText())
//}
//
//func (*PrintTinyListener) ExitFuncInvoc(c *parser.FuncInvocContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitFuncArgsInvoc(c *parser.FuncArgsInvocContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitFuncInvocArgs(c *parser.FuncInvocArgsContext) {
//	panic("implement me")
//}


//func (*PrintTinyListener) ExitNewVariable(c *parser.NewVariableContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitArrayElem(c *parser.ArrayElemContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitArrayInit(c *parser.ArrayInitContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitArrayInitEmpty(c *parser.ArrayInitEmptyContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitArrayInitValue(c *parser.ArrayInitValueContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitArrayInitElems(c *parser.ArrayInitElemsContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitExpr(c *parser.ExprContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitExprOperand(c *parser.ExprOperandContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitBoolExpr(c *parser.BoolExprContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitBoolExprOperand(c *parser.BoolExprOperandContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitStatementBody(c *parser.StatementBodyContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitIfElseSt(c *parser.IfElseStContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitIfSt(c *parser.IfStContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitElseIfSt(c *parser.ElseIfStContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitElseSt(c *parser.ElseStContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitWhileSt(c *parser.WhileStContext) {
//	panic("implement me")
//}
//
//func (*PrintTinyListener) ExitForSt(c *parser.ForStContext) {
//	panic("implement me")
//}

func NewPrintListener() *PrintTinyListener {
	return new(PrintTinyListener)
}







