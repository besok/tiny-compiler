package env

import (
	"fmt"
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

//func (l *InterListener) VisitErrorNode(node antlr.ErrorNode) {
//	panic("implement me")
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

//func (l *InterListener) EnterInternalArrArg(c *parser.InternalArrArgContext) {
//
//}

func (l *InterListener) EnterNewVar(c *parser.NewVarContext) {
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

	st := NewVarSt{Line: line, Name:name, T: typeVal, ArrArgs: make([]IntArrArgSt, 0)}

	for _, el := range c.AllInternalArrArg() {
		ctx := el.(*parser.InternalArrArgContext)
		line := ctx.GetStart().GetLine()
		iVar := makeIVar(ctx.InternalVar())

		stV := IntArrArgSt{Var: iVar, Line: line}

		if ctx.NUMBER() != nil {
			n, _ := strconv.Atoi(ctx.NUMBER().GetText())
			stV.Value = n
			stV.Type = N
		}

		if ctx.TRUE() != nil {
			stV.Value = true
			stV.Type = B
		}
		if ctx.FALSE() != nil {
			stV.Value = false
			stV.Type = B
		}
		if ctx.STRING_RAW() != nil {
			raw := ctx.STRING_RAW().GetText()
			raw = strings.TrimPrefix(raw, "\"")
			raw = strings.TrimSuffix(raw, "\"")
			stV.Type = S
			stV.Value = raw
		}

		st.ArrArgs = append(st.ArrArgs,stV)
	}

	Put(ToBody(st))

}

func (l *InterListener) EnterUpdVar(c *parser.UpdVarContext) {
	line := c.GetStart().GetLine()
	text := c.ITEM().GetText()
	iVar := makeIVar(c.InternalVar())
	Put(ToBody(UpdVarSt{Line: line, Name: text, Var: iVar}))
}

func (l *InterListener) EnterInitItem(c *parser.InitItemContext) {
	line := c.GetStart().GetLine()
	text := c.ITEM().GetText()
	iVar := makeIVar(c.InternalVar())

	Put(ToBody(InitItemSt{Line: line, Item: text, Var: iVar}))
}

func (l *InterListener) EnterInitIntVar(c *parser.InitIntVarContext) {
	line := c.GetStart().GetLine()
	Put(ToBody(InitInternalVarSt{Line:line, Left: makeIVar(c.InternalVar(0)), Right: makeIVar(c.InternalVar(1))}))
}

func (l *InterListener) EnterInitPrim(c *parser.InitPrimContext) {
	line := c.GetStart().GetLine()
	v := makeIVar(c.InternalVar())
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
		raw = strings.TrimPrefix(raw, "\"")
		raw = strings.TrimSuffix(raw, "\"")
		st.T = S
		st.Val = raw

	}

	Put(ToBody(st))

}

func (l *InterListener) EnterInitParam(c *parser.InitParamContext) {
	line := c.GetStart().GetLine()
	left := makeIVar(c.InternalVar(0))
	right := makeIVar(c.InternalVar(1))
	Put(ToBody(ParamSt{Line: line, Left: left, Right: right}))
}
func (l *InterListener) EnterInitCall(c *parser.InitCallContext) {
	line := c.GetStart().GetLine()
	number, _ := strconv.Atoi(c.NUMBER().GetText())
	iVar := makeIVar(c.InternalVar())
	isSys := false
	name := ""
	if c.SYS_FUNC() != nil {
		isSys = true
		name = c.SYS_FUNC().GetText()
	} else {
		name = c.ITEM().GetText()
	}

	Put(ToBody(CallSt{Var: iVar, Line: line, Count: number, Func: name, IsSys: isSys}))

}

//func (l *InterListener) EnterInitBoolOp(c *parser.InitBoolOpContext) {
//	panic("implement me")
//}
func (l *InterListener) EnterInitNumOp(c *parser.InitNumOpContext) {
	line := c.GetStart().GetLine()
	left := makeIVar(c.InternalVar(0))
	right := makeIVar(c.InternalVar(1))
	text := c.NUM_SIGN().GetText()

	Put(ToBody(InitNumBoolOpSt{Line: line, Left: left, Right: right, Sign: text, IsBool: false}))
}

func (l *InterListener) EnterInitArrEl(c *parser.InitArrElContext) {
	line := c.GetStart().GetLine()
	name := c.ITEM().GetText()
	left := makeIVar(c.InternalVar(0))
	right := makeIVar(c.InternalVar(0))
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
		iVar := makeIVar(c.InternalVar(0))
		st.Right = iVar
	} else {
		l := makeIVar(c.InternalVar(0))
		r := makeIVar(c.InternalVar(1))
		st.IsNum = false
		st.Right = r
		st.Left = l
	}

	Put(ToBody(st))
}

func (l *InterListener) EnterReturnTp(c *parser.ReturnTpContext) {
	line := c.GetStart().GetLine()
	Put(ToBody(ReturnSt{Line: line, Var: makeIVar(c.InternalVar())}))

}

func makeIVar(ivCtx parser.IInternalVarContext) InternalVar {
	inVar := ivCtx.(*parser.InternalVarContext)
	number, _ := strconv.Atoi(inVar.NUMBER().GetText())
	innerVar := strings.TrimPrefix(inVar.INNER_VAR().GetText(), "_")
	iv := InternalVar{T: innerVar, N: number}
	return iv
}

func (l *InterListener) EnterGotoTp(c *parser.GotoTpContext) {
	in := c.GotoIn()
	gt, _ := strconv.Atoi(in.(*parser.GotoInContext).NUMBER().GetText())
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
	iVar := makeIVar(c.InternalVar())
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

func (l *InterListener) ExitInitBoolOp(c *parser.InitBoolOpContext) {
	line := c.GetStart().GetLine()
	left := makeIVar(c.InternalVar(0))
	right := makeIVar(c.InternalVar(1))
	text := c.BOOL_SIGN().GetText()

	Put(ToBody(InitNumBoolOpSt{Line: line, Left: left, Right: right, Sign: text, IsBool: true}))
}

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

type BodyStatement interface{
	handle(frame *RecordTable) (interface{}, bool)
	line() int
}

type InitInternalVarSt struct {
	Line  int
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
type ParamSt struct {
	Line  int
	Left  InternalVar
	Right InternalVar
}

type CallSt struct {
	Line  int
	Count int
	Var   InternalVar
	Func  string
	IsSys bool
}

type InitItemSt struct {
	Line int
	Var  InternalVar
	Item string
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
func(v *InternalVar) makeName() string {
	return fmt.Sprintf("%s%d",v.T,v.N)
}
type UpdVarSt struct {
	Line int
	Name string
	Var  InternalVar
}

type NewVarSt struct {
	Line    int
	Name    string
	T       Type
	ArrArgs []IntArrArgSt
}

type IntArrArgSt struct {
	Line  int
	Var   InternalVar
	Type  PrimitiveType
	Value interface{}
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
