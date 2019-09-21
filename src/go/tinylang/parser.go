package parsing

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"tiny-compiler/src/antlr/tiny"
)

func IR(path string) *os.File {
	s := Parse(path)

	for _,f := range s.Functions{
		f.Process()
	}

	file, _ := CreateFile(path)
	e := WriteFile(file)
	if e != nil{
		log.Fatalf("error to save to file: %s ",e)
	}
	return file
}

func Parse(path string) Script {
	open, e := os.Open(path)
	if e!= nil {
		log.Fatalf(" error to open a file :%s, error: %s",path,e)
	}

	e = open.Close()

	inp, _ := antlr.NewFileStream(path)
	lexer := parser.NewTinyLangLexer(inp)
	p := parser.NewTinyLangParser(antlr.NewCommonTokenStream(lexer, 0))

	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	file := p.File()

	antlr.ParseTreeWalkerDefault.Walk(NewListener(), file)

	return script
}

type CommonTinyListener struct {
	*parser.BaseTinyLangListener
}

func NewListener() *CommonTinyListener {
	return &CommonTinyListener{}
}

//func (l *CommonTinyListener) VisitTerminal(node antlr.TerminalNode) {
//}

func (l *CommonTinyListener) VisitErrorNode(node antlr.ErrorNode) {
	if pop, b := Pop(); b {
		ne := pop
		ln := (*ne).line()
		for ; ln == -1; {
			ne, b = Pop()
			if !b {
				break
			}
			ln = (*ne).line()
		}
		tp := reflect.TypeOf((*ne).get()).Name()
		log.Fatalf(" compile error: %s in %s, line:%d", node.GetText(),tp ,ln)
	}else{
		log.Fatalf(" compile error: %s , line:%d", node.GetText(),0)
	}
}

//
//func (l *CommonTinyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
//	panic("implement me")
//}
//
//func (l *CommonTinyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
//	panic("implement me")
//}
//
//func (l *CommonTinyListener) EnterFile(c *parser.FileContext) {
//	panic("implement me")
//}
//
func (l *CommonTinyListener) EnterFuncInit(c *parser.FuncInitContext) {
	line := c.GetStart().GetLine()
	funcName := c.ITEM().GetText()

	fd := NewFuncDefinition(line, funcName)
	dCtx := &FuncDefinitionCtx{F: fd, S: Start}
	Push(dCtx)
}

//func (l *CommonTinyListener) EnterFuncArgs(c *parser.FuncArgsContext) {
//}

func (l *CommonTinyListener) EnterFuncReturn(c *parser.FuncReturnContext) {
	ctx := FuncReturn{Line: c.GetStart().GetLine()}
	Push(&ctx)

}

func (l *CommonTinyListener) EnterFuncInvoc(c *parser.FuncInvocContext) {
	name := ""
	if sf := c.SYS_FUNC(); sf != nil {
		name = sf.GetText()
	} else {
		name = c.ITEM().GetText()
	}

	ct := FuncInvoc{Name: name, Line: c.GetStart().GetLine(), Args: make([]interface{}, 0)}
	Push(&ct)

}

//func (l *CommonTinyListener) EnterFuncArgsInvoc(c *parser.FuncArgsInvocContext) {
//	panic("implement me")
//}
//func (l *CommonTinyListener) EnterFuncInvocArgs(c *parser.FuncInvocArgsContext) {
//	panic("implement me")
//}
func (l *CommonTinyListener) EnterUpdVariable(c *parser.UpdVariableContext) {
	v := UpdVarCtx{V: UpdVar{Line: c.GetStart().GetLine()}, S: Start}
	item := c.ITEM()
	if item != nil {
		v.V.Left = struct {
			IsArrEl bool
			ArrElem ArrayElem
		}{IsArrEl: false, ArrElem: ArrayElem{Name: item.GetText()}}
		v.S = End
	} else {
		v.V.Left = struct {
			IsArrEl bool
			ArrElem ArrayElem
		}{IsArrEl: true}
	}
	Push(&v)
}

func (l *CommonTinyListener) EnterNewVariable(c *parser.NewVariableContext) {
	line := c.GetStart().GetLine()
	name := c.ITEM().GetText()
	tp := c.VariableType().GetText()
	isArr := strings.HasPrefix(tp, "[]")
	if isArr {
		tp = strings.TrimPrefix(tp, "[]")
	}
	t := Type{IsArray: isArr, T: tp}
	newV := &NewVariable{Line: line, Name: name, Type: t}
	Push(newV)
}
func (l *CommonTinyListener) EnterVal(c *parser.ValContext) {
	v := Val{}
	if bl := c.TRUE(); bl != nil {
		v.T = "b"
		v.V = "true"
	} else if bl := c.FALSE(); bl != nil {
		v.T = "b"
		v.V = "false"
	} else if raw := c.STRING_RAW(); raw != nil {
		v.T = "s"
		v.V = raw.GetText()
	} else if raw := c.NUMBER(); raw != nil {
		v.T = "n"
		v.V = raw.GetText()
	} else if raw := c.ITEM(); raw != nil {
		v.T = "i"
		v.V = raw.GetText()
	}

	Release(v)
}

//func (l *CommonTinyListener) EnterVariableType(c *parser.VariableTypeContext) {
//	panic("implement me")
//}
//
func (l *CommonTinyListener) EnterArrayElem(c *parser.ArrayElemContext) {
	name := c.ITEM(0).GetText()
	elem := ArrayElem{Name: name, HasPos: false}

	if nm := c.NUMBER(); nm != nil {
		number, _ := strconv.Atoi(nm.GetText())
		elem.HasPos = true
		elem.Pos = number
	}

	if nextItem := c.ITEM(1); nextItem != nil {
		elem.HasPos = false
		elem.Calc = nextItem.GetText()
	}

	Push(&elem)

}

//
//func (l *CommonTinyListener) EnterArrayInit(c *parser.ArrayInitContext) {
//	panic("implement me")
//}
//
func (l *CommonTinyListener) EnterArrayInitEmpty(c *parser.ArrayInitEmptyContext) {
	cp := c.NUMBER()
	capacity := 0
	if cp != nil {
		caps, err := strconv.Atoi(cp.GetText())
		if err != nil {
			log.Fatalln("error for array init, line: ", c.GetStart().GetLine())
		}
		capacity = caps
	}

	t := Type{}
	if c.NUM() != nil {
		t.T = c.NUM().GetText()
	}
	if c.STRING() != nil {
		t.T = c.STRING().GetText()
	}
	if c.BOOL() != nil {
		t.T = c.BOOL().GetText()
	}

	arrInit := ArrayInit{T: t, Cap: capacity}
	Release(arrInit)
}

//
//func (l *CommonTinyListener) EnterArrayInitValue(c *parser.ArrayInitValueContext) {
//	panic("implement me")
//}
//
func (l *CommonTinyListener) EnterArrayInitElems(c *parser.ArrayInitElemsContext) {

	bAll := c.AllBOOL_VAL()
	nAll := c.AllNUMBER()
	sAll := c.AllSTRING_RAW()

	commonLen := 0
	if len(bAll) > 0 {
		commonLen++
	}
	if len(sAll) > 0 {
		commonLen++
	}
	if len(nAll) > 0 {
		commonLen++
	}

	if commonLen > 1 {
		log.Fatalln("array should has only one type, line: ", c.GetStart().GetLine())
	}

	arrInit := ArrayInit{}

	if len(bAll) > 0 {
		ln := len(bAll)
		boolVals := make([]interface{}, 0)
		for i := 0; i < ln; i++ {
			text := c.BOOL_VAL(i).GetText()
			f := false
			if text == "true" {
				f = true
			}
			boolVals = append(boolVals, f)
		}
		arrInit = ArrayInit{T: Type{T: "bool"}, Cap: ln, Val: boolVals}

	}
	if len(nAll) > 0 {
		ln := len(nAll)
		intVals := make([]interface{}, 0)
		for i := 0; i < ln; i++ {
			text := c.NUMBER(i).GetText()
			f, _ := strconv.Atoi(text)
			intVals = append(intVals, f)
		}
		arrInit = ArrayInit{T: Type{T: "num"}, Cap: ln, Val: intVals}
	}
	if len(sAll) > 0 {
		ln := len(sAll)
		intVals := make([]interface{}, 0)
		for i := 0; i < ln; i++ {
			text := c.STRING_RAW(i).GetText()
			intVals = append(intVals, text)
		}
		arrInit = ArrayInit{T: Type{T: "string"}, Cap: ln, Val: intVals}
	}

	Release(arrInit)

}

//
func (l *CommonTinyListener) EnterExpr(c *parser.ExprContext) {
	sg := c.NUM_SIGN().GetText()
	e := ExprCtx{S: Start, E: Expr{Line: c.GetStart().GetLine(), Sign: sg}}
	Push(&e)
}
func (l *CommonTinyListener) EnterBoolExpr(c *parser.BoolExprContext) {
	ctx := BoolExpr{Line: c.GetStart().GetLine(), BoolExpr: make([]BoolExprSingleExtra, 0)}
	Push(&ctx)
}
func (l *CommonTinyListener) EnterBreakOrContinue(c *parser.BreakOrContinueContext) {
	pop, _ := Pop()
	if br := c.BREAK(); br != nil {
		(*pop).PutItem(BreakOrContinue{IsBreak: true})
	} else {
		(*pop).PutItem(BreakOrContinue{IsBreak: false})
	}
	Push(*pop)
}

func (l *CommonTinyListener) EnterExprOperand(c *parser.ExprOperandContext) {
	op := ExprOperand{}
	if i := c.NUMBER(); i != nil {
		op.T = "n"
		op.V = i.GetText()
	} else if i := c.STRING_RAW(); i != nil {
		op.T = "s"
		op.V = i.GetText()
	} else if i := c.ITEM(); i != nil {
		op.T = "i"
		op.V = i.GetText()
	}
	Push(&op)
}

func (l *CommonTinyListener) EnterBoolExprSingle(c *parser.BoolExprSingleContext) {
	bexpS := BoolExprSingleCtx{S: Start, V: BoolExprSingle{Sign: c.BOOL_SIGN().GetText()}}
	Push(&bexpS)
}
func (l *CommonTinyListener) EnterBoolExprSingleExtra(c *parser.BoolExprSingleExtraContext) {
	pl := c.BOOL_PL().GetText()
	bexpS := BoolExprSingleExtra{Op: pl}
	Push(&bexpS)
}
func (l *CommonTinyListener) EnterBoolExprOperand(c *parser.BoolExprOperandContext) {
	op := BoolExprOperand{}
	if tr := c.TRUE(); tr != nil {
		op.IsPrim = true
		op.V = true
	} else if fl := c.FALSE(); fl != nil {
		op.IsPrim = true
		op.V = false
	} else {
		op.IsPrim = false
	}
	Push(&op)
}

func (l *CommonTinyListener) EnterStatementBody(c *parser.StatementBodyContext) {
	bodyCtx := StatementBody{V: make([]interface{}, 0)}
	Push(&bodyCtx)
}

func (l *CommonTinyListener) EnterIfElseSt(c *parser.IfElseStContext) {
	stCtx := IfElseIfStCtx{S: Start, V: IfElseIfSt{Line: c.GetStart().GetLine(), ElseIf: make([]IfSt, 0)}}
	Push(&stCtx)
}

func (l *CommonTinyListener) EnterIfSt(c *parser.IfStContext) {
	ctx := IfStCtx{S: Start}
	if it := c.ITEM(); it != nil {
		ctx.V.Cond = it.GetText()
		ctx.S = End
	}
	if it := c.FALSE(); it != nil {
		ctx.V.Cond = it.GetText()
		ctx.S = End
	}
	if it := c.TRUE(); it != nil {
		ctx.V.Cond = it.GetText()
		ctx.S = End
	}
	Push(&ctx)
}

func (l *CommonTinyListener) EnterElseIfSt(c *parser.ElseIfStContext) {

}

func (l *CommonTinyListener) EnterElseSt(c *parser.ElseStContext) {
	prev, _ := Pop()
	(*prev).setState(End)
	Push(*prev)
	e := ElseSt{}
	Push(&e)
}

func (l *CommonTinyListener) EnterWhileSt(c *parser.WhileStContext) {
	wCtx := WhileStCtx{S: Start, W: WhileSt{Line: c.GetStart().GetLine()}}

	if it := c.ITEM(); it != nil {
		wCtx.W.BoolExpr = it.GetText()
		wCtx.W.BoolExprT = "i"
		wCtx.S = End
	} else if it := c.FALSE(); it != nil {
		wCtx.W.BoolExpr = it.GetText()
		wCtx.W.BoolExprT = "b"
		wCtx.S = End
	} else if it := c.TRUE(); it != nil {
		wCtx.W.BoolExpr = it.GetText()
		wCtx.W.BoolExprT = "b"
		wCtx.S = End
	} else {
		wCtx.W.BoolExprT = "e"
	}
	Push(&wCtx)
}

//
func (l *CommonTinyListener) EnterForSt(c *parser.ForStContext) {
	stCtx := ForStCtx{F: ForSt{Line: c.GetStart().GetLine()}, S: Start}
	Push(&stCtx)
}

//func (l *CommonTinyListener) ExitFile(c *parser.FileContext) {
//	panic("implement me")
//}

func (l *CommonTinyListener) EnterFuncArg(ctx *parser.FuncArgContext) {
	item := ctx.ITEM().GetText()
	tp := ctx.VariableType().GetText()
	isArr := strings.HasPrefix(tp, "[]")
	if isArr {
		tp = strings.TrimPrefix(tp, "[]")
	}
	v := Var{Name: item, Type: Type{IsArray: isArr, T: tp}}

	Release(v)
}

func (l *CommonTinyListener) ExitFuncInit(c *parser.FuncInitContext) {
	ctx, ok := Pop()
	if !ok {
		log.Panic("error!")
	}
	(*ctx).(*FuncDefinitionCtx).Close()
}

func (l *CommonTinyListener) ExitFuncArgs(c *parser.FuncArgsContext) {
	ctx, _ := Pop()
	(*ctx).(*FuncDefinitionCtx).setState(Middle)
	Push(*ctx)
}
func (l *CommonTinyListener) EnterFuncReturnType(c *parser.FuncReturnTypeContext) {
	v := Var{Name: "return", Type: Type{IsArray: false, T: ""}}
	vt := c.VariableType()

	if vt == nil {
		v.Type.T = "void"
	} else {
		tp := vt.GetText()
		isArr := strings.HasPrefix(tp, "[]")
		if isArr {
			tp = strings.TrimPrefix(tp, "[]")
		}
		v.Type.IsArray = isArr
		v.Type.T = tp
	}

	Release(v)
}
func (l *CommonTinyListener) ExitFuncReturnType(c *parser.FuncReturnTypeContext) {
	ctx, _ := Pop()
	(*ctx).setState(Next)
	Push(*ctx)
}

//
func (l *CommonTinyListener) ExitFuncReturn(c *parser.FuncReturnContext) {
	fr, _ := Pop()
	Release(*fr)
}

func (l *CommonTinyListener) ExitFuncInvoc(c *parser.FuncInvocContext) {
	fInv, _ := Pop()
	Release(*fInv)
}

//
//func (l *CommonTinyListener) ExitFuncArgsInvoc(c *parser.FuncArgsInvocContext) {
//	panic("implement me")
//}
//
//func (l *CommonTinyListener) ExitFuncInvocArgs(c *parser.FuncInvocArgsContext) {
//	panic("implement me")
//}.
//
func (l *CommonTinyListener) ExitUpdVariable(c *parser.UpdVariableContext) {
	ctx, _ := Pop()
	Release(*ctx)
}

//
func (l *CommonTinyListener) ExitNewVariable(c *parser.NewVariableContext) {
	ctx, _ := Pop()
	Release(*ctx)
}

//
//func (l *CommonTinyListener) ExitVariableType(c *parser.VariableTypeContext) {
//	panic("implement me")
//}
//
func (l *CommonTinyListener) ExitArrayElem(c *parser.ArrayElemContext) {
	ctx, _ := Pop()
	Release(*ctx)
}

//
//func (l *CommonTinyListener) ExitArrayInit(c *parser.ArrayInitContext) {
//	panic("implement me")
//}
//
//func (l *CommonTinyListener) ExitArrayInitEmpty(c *parser.ArrayInitEmptyContext) {
//	panic("implement me")
//}
//
//func (l *CommonTinyListener) ExitArrayInitValue(c *parser.ArrayInitValueContext) {
//	panic("implement me")
//}
//
//func (l *CommonTinyListener) ExitArrayInitElems(c *parser.ArrayInitElemsContext) {
//	panic("implement me")
//}
//
func (l *CommonTinyListener) ExitExpr(c *parser.ExprContext) {
	ctx, _ := Pop()
	Release(*ctx)
}
func (l *CommonTinyListener) ExitBoolExpr(c *parser.BoolExprContext) {
	ctx, _ := Pop()
	Release(*ctx)
}

//
func (l *CommonTinyListener) ExitExprOperand(c *parser.ExprOperandContext) {
	ctx, _ := Pop()
	Release(*ctx)
}

func (l *CommonTinyListener) ExitBoolExprSingle(c *parser.BoolExprSingleContext) {
	bEs, _ := Pop()
	Release(*bEs)
}
func (l *CommonTinyListener) ExitBoolExprSingleExtra(c *parser.BoolExprSingleExtraContext) {
	bEs, _ := Pop()
	Release(*bEs)
}

func (l *CommonTinyListener) ExitBoolExprOperand(c *parser.BoolExprOperandContext) {
	ct, _ := Pop()
	Release(*ct)
}
func (l *CommonTinyListener) ExitStatementBody(c *parser.StatementBodyContext) {
	ct, _ := Pop()
	Release(*ct)
}

func (l *CommonTinyListener) ExitIfElseSt(c *parser.IfElseStContext) {
	ctx, _ := Pop()
	Release(*ctx)
}

func (l *CommonTinyListener) ExitIfSt(c *parser.IfStContext) {
	ctx, _ := Pop()
	Release(*ctx)
}
func (l *CommonTinyListener) ExitElseIfSt(c *parser.ElseIfStContext) {
	//ctx, _ := Pop()
	//Release(*ctx)
}
func (l *CommonTinyListener) ExitElseSt(c *parser.ElseStContext) {
	ctx, _ := Pop()
	Release(*ctx)
}

func (l *CommonTinyListener) ExitWhileSt(c *parser.WhileStContext) {
	p, _ := Pop()
	Release(*p)
}

func (l *CommonTinyListener) ExitForSt(c *parser.ForStContext) {
	p, _ := Pop()
	Release(*p)
}
