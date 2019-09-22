package env

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
	"os"
	"strconv"
	"strings"
	"tiny-compiler/src/antlr/inter"
)

func ParseIR(path string) Functions {
	open, e := os.Open(path)
	if e != nil {
		log.Fatalf(" error to open a file :%s, error: %s", path, e)
	}

	e = open.Close()

	inp, _ := antlr.NewFileStream(path)
	lexer := parser.NewInterLangLexer(inp)
	p := parser.NewInterLangParser(antlr.NewCommonTokenStream(lexer, 0))

	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	file := p.File()

	antlr.ParseTreeWalkerDefault.Walk(&InterListener{}, file)

	return funcs
}

type InterListener struct {
	*parser.BaseInterLangListener
}

//
//func (l *InterListener) VisitTerminal(node antlr.TerminalNode) {
//	panic("implement me")
//}
//
//func (l *InterListener) VisitErrorNode(node antlr.ErrorNode) {
//	panic("implement me")
//}
//
//func (l *InterListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
//	panic("implement me")
//}

//func (l *InterListener) EnterFile(c *parser.FileContext) {
//
//}

func (l *InterListener) EnterFunction(c *parser.FunctionContext) {
	line := c.GetStart().GetLine()
	name := c.ITEM().GetText()
	f := Function{Name: name, Line: line, Args: make([]Arg, 0), Body: make([]BodyStatement, 0)}
	Push(f)
}

func (l *InterListener) EnterArg(c *parser.ArgContext) {

	line := c.GetStart().GetLine()
	name := c.ITEM().GetText()

	tp := c.TypeTp().(*parser.TypeTpContext)

	typeVal := Type{IsArray: false}

	if tp.ARRAY() != nil {
		typeVal.IsArray = true
	}

	if tp.BOOL() != nil {
		typeVal.T = B
	}
	if tp.NUM() != nil {
		typeVal.T = N
	}
	if tp.STR() != nil {
		typeVal.T = S
	}

	arg := Arg{Name: name, Line: line, Type: typeVal}

	function := (*Pop()).(Function)
	function.PutArg(arg)
	Push(function)

}

func (l *InterListener) EnterRetTp(c *parser.RetTpContext) {
	returnType := ReturnType{IsVoid: false}

	if c.VOID() != nil {
		returnType.IsVoid = true
	} else {
		tp := c.TypeTp().(*parser.TypeTpContext)
		typeVal := Type{IsArray: false}

		if tp.ARRAY() != nil {
			typeVal.IsArray = true
		}

		if tp.BOOL() != nil {
			typeVal.T = B
		}
		if tp.NUM() != nil {
			typeVal.T = N
		}
		if tp.STR() != nil {
			typeVal.T = S
		}
		returnType.Type = typeVal
	}

	Put(func(f Function) Function {
		f.ReturnType = returnType
		return f
	})

}

//
//func (l *InterListener) EnterStatement(c *parser.StatementContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) EnterInternalArrArg(c *parser.InternalArrArgContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) EnterInternalVar(c *parser.InternalVarContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) EnterNewVar(c *parser.NewVarContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) EnterNewArrVar(c *parser.NewArrVarContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) EnterUpdVar(c *parser.UpdVarContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) EnterInitItem(c *parser.InitItemContext) {
//	panic("implement me")
//}

func (l *InterListener) EnterInitIntVar(c *parser.InitIntVarContext) {
	left := c.InternalVar(0).(*parser.InternalVarContext)
	right := c.InternalVar(1).(*parser.InternalVarContext)

	leftVar := makeIVar(left)
	rightVar := makeIVar(right)
	st := InitInternalVarSt{Left: leftVar, Right: rightVar}
	Put(ToBody(st))
}

//
func (l *InterListener) EnterInitPrim(c *parser.InitPrimContext) {
	line := c.GetStart().GetLine()
	v := makeIVar(c.InternalVar().(*parser.InternalVarContext))
	st := InitPrimSt{Line: line, Left: v}

	if c.NUMBER() != nil {
		st.T = N
		nu, _ := strconv.Atoi(c.NUMBER().GetText())
		st.Val = nu
	}

	if c.TRUE() != nil {
		st.T = B
		st.Val = true
	}
	if c.FALSE() != nil {
		st.T = B
		st.Val = false
	}
	if c.STRING_RAW() != nil {
		raw := c.STRING_RAW().GetText()
		raw = strings.TrimPrefix(raw,"\"")
		raw = strings.TrimSuffix(raw,"\"")
		st.T = S
		st.Val = raw

	}

	Put(ToBody(st))

}

//
//func (l *InterListener) EnterInitParam(c *parser.InitParamContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) EnterInitCall(c *parser.InitCallContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) EnterInitBoolOp(c *parser.InitBoolOpContext) {
//	panic("implement me")
//}
//
func (l *InterListener) EnterInitNumOp(c *parser.InitNumOpContext) {
	line := c.GetStart().GetLine()
	left := makeIVar(c.InternalVar(0).(*parser.InternalVarContext))
	right := makeIVar(c.InternalVar(1).(*parser.InternalVarContext))
	text := c.NUM_SIGN().GetText()

	Put(ToBody(InitNumBoolOpSt{Line: line, Left: left, Right: right, Sign: text, IsBool: false}))
}

func (l *InterListener) EnterInitArrEl(c *parser.InitArrElContext) {
	line := c.GetStart().GetLine()
	name := c.ITEM().GetText()
	left := makeIVar(c.InternalVar(0).(*parser.InternalVarContext))
	right := makeIVar(c.InternalVar(0).(*parser.InternalVarContext))
	st := InitArrElemSt{Line: line, Name: name, Left: left, Right: right}
	Put(ToBody(st))
}

func (l *InterListener) EnterExtArrEl(c *parser.ExtArrElContext) {
	line := c.GetStart().GetLine()
	name := c.ITEM().GetText()
	st := ExtArrElemSt{Line: line, Name: name, IsNum: true}
	n := c.NUMBER()
	if n != nil {
		num, _ := strconv.Atoi(n.GetText())
		st.Num = num
		iVar := makeIVar(c.InternalVar(0).(*parser.InternalVarContext))
		st.Right = iVar
	} else {
		l := makeIVar(c.InternalVar(0).(*parser.InternalVarContext))
		r := makeIVar(c.InternalVar(1).(*parser.InternalVarContext))
		st.IsNum = false
		st.Right = r
		st.Left = l
	}

	Put(ToBody(st))
}

func (l *InterListener) EnterReturnTp(c *parser.ReturnTpContext) {
	line := c.GetStart().GetLine()
	inVar := c.InternalVar().(*parser.InternalVarContext)
	iv := makeIVar(inVar)

	Put(ToBody(ReturnSt{Line: line, Var: iv}))

}

func makeIVar(inVar *parser.InternalVarContext) InternalVar {
	number, _ := strconv.Atoi(inVar.NUMBER().GetText())
	innerVar := strings.TrimPrefix(inVar.INNER_VAR().GetText(), "_")
	iv := InternalVar{T: innerVar, N: number}
	return iv
}

func (l *InterListener) EnterGotoIn(c *parser.GotoInContext) {
	gt, _ := strconv.Atoi(c.NUMBER().GetText())
	line := c.GetStart().GetLine()

	Put(ToBody(GotoSt{Line: line, Goto: gt}))
}

//
//func (l *InterListener) EnterGoto(c *parser.GotoContext) {
//	panic("implement me")
//}
//
func (l *InterListener) EnterIfFalse(c *parser.IfFalseContext) {
	line := c.GetStart().GetLine()
	iVar := makeIVar(c.InternalVar().(*parser.InternalVarContext))
	gt, _ := strconv.Atoi(c.GotoIn().(*parser.GotoInContext).NUMBER().GetText())

	Put(ToBody(IfFalseSt{Line: line, Var: iVar, Goto: gt}))

}

//
//func (l *InterListener) EnterType(c *parser.TypeContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) EnterLine(c *parser.LineContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitFile(c *parser.FileContext) {
//	panic("implement me")
//}

func (l *InterListener) ExitFunction(c *parser.FunctionContext) {
	f := *Pop()
	switch f.(type) {
	case Function:
		funcs.List = append(funcs.List, f.(Function))
	default:
		panic("should be func")
	}
}

//func (l *InterListener) ExitArg(c *parser.ArgContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitRetTp(c *parser.RetTpContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitStatement(c *parser.StatementContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitInternalArrArg(c *parser.InternalArrArgContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitInternalVar(c *parser.InternalVarContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitNewVar(c *parser.NewVarContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitNewArrVar(c *parser.NewArrVarContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitUpdVar(c *parser.UpdVarContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitInitItem(c *parser.InitItemContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitInitPrim(c *parser.InitPrimContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitInitParam(c *parser.InitParamContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitInitCall(c *parser.InitCallContext) {
//	panic("implement me")
//}
//
func (l *InterListener) ExitInitBoolOp(c *parser.InitBoolOpContext) {
	line := c.GetStart().GetLine()
	left := makeIVar(c.InternalVar(0).(*parser.InternalVarContext))
	right := makeIVar(c.InternalVar(1).(*parser.InternalVarContext))
	text := c.BOOL_SIGN().GetText()

	Put(ToBody(InitNumBoolOpSt{Line: line, Left: left, Right: right, Sign: text, IsBool: true}))
}

//
//func (l *InterListener) ExitInitNumOp(c *parser.InitNumOpContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitInitArrEl(c *parser.InitArrElContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitExtArrEl(c *parser.ExtArrElContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitReturn(c *parser.ReturnContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitGotoIn(c *parser.GotoInContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitGoto(c *parser.GotoContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitIfFalse(c *parser.IfFalseContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitType(c *parser.TypeContext) {
//	panic("implement me")
//}
//
//func (l *InterListener) ExitLine(c *parser.LineContext) {
//	panic("implement me")
//}

type Ctx interface{}

type Functions struct {
	List []Function
}

var funcs = Functions{List: make([]Function, 0)}

type Function struct {
	Line       int
	Name       string
	Args       []Arg
	ReturnType ReturnType
	Body       []BodyStatement
}

func (f *Function) PutArg(arg Arg) {
	f.Args = append(f.Args, arg)
}
func (f *Function) PutSt(st BodyStatement) {
	f.Body = append(f.Body, st)
}

func ToBody(st BodyStatement) func(Function) Function {
	return func(function Function) Function {
		function.PutSt(st)
		return function
	}
}

func Put(f func(Function) Function) {
	Push(f((*Pop()).(Function)))
}

type Arg struct {
	Line int
	Name string
	Type Type
}

type PrimitiveType string

const (
	S PrimitiveType = "S"
	B PrimitiveType = "B"
	N PrimitiveType = "N"
)

type Type struct {
	T       PrimitiveType
	IsArray bool
}

type ReturnType struct {
	Type   Type
	IsVoid bool
}

type ReturnSt struct {
	Line int
	Var  InternalVar
}

type BodyStatement interface{}

type InitInternalVarSt struct {
	Left  InternalVar
	Right InternalVar
}
type IfFalseSt struct {
	Line int
	Var  InternalVar
	Goto int
}

type GotoSt struct {
	Line int
	Goto int
}

type InitPrimSt struct {
	Line int
	Left InternalVar
	T    PrimitiveType
	Val  interface{}
}

type ExtArrElemSt struct {
	Line  int
	Name  string
	Num   int
	IsNum bool
	Left  InternalVar
	Right InternalVar
}
type InitNumBoolOpSt struct {
	Line   int
	Left   InternalVar
	Right  InternalVar
	Sign   string
	IsBool bool
}
type InitArrElemSt struct {
	Line  int
	Name  string
	Left  InternalVar
	Right InternalVar
}

type InternalVar struct {
	N int
	T string
}

var ctx = make([]Ctx, 0)

func Push(c Ctx) {
	ctx = append(ctx, c)
}
func Pop() *Ctx {
	l := len(ctx)
	if l == 0 {
		return nil
	}
	ret := ctx[l-1]
	ctx = ctx[:l-1]
	return &ret
}
