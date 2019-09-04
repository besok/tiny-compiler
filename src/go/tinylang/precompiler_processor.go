package parsing

import (
	"fmt"
	"strconv"
)

// init
// init_arr
// param
// call
// goto
// ifTrue x goto line
// ifFalse x goto line
// goto
// bool sign
// arr item
// return

type View interface {
	Process() string
}

func (uv UpdVar) Process() string {
	right := uv.Right.(View).Process()
	left := uv.Left
	if left.IsArrEl {
		addLine(fmt.Sprintf("%s[%d] = %s", left.ArrElem.Name, left.ArrElem.Pos, right))
	} else {
		addLine(fmt.Sprintf("%s = %s", left.ArrElem.Name, right))
	}
	return ""
}

func (beo BoolExprOperand) Process() string {
	if beo.IsPrim {
		return strconv.FormatBool(beo.V.(bool))
	}
	return beo.V.(View).Process()
}

func (be BoolExpr) Process() string {
	exprs := be.BoolExpr
	prev := ""
	for _, e := range exprs {
		sign := e.V.Sign
		left := createVar()
		leftRes := e.V.Left.(View).Process()
		addLine(fmt.Sprintf("%s = %s", left, leftRes))
		right := createVar()
		rightRes := e.V.Right.(View).Process()
		addLine(fmt.Sprintf("%s = %s", right, rightRes))

		fin := createVar()
		addLine(fmt.Sprintf("%s = %s %s %s", fin, left, sign, right))
		var afterFin = fin
		if prev != "" {
			afterFin = createVar()
			addLine(fmt.Sprintf("%s = %s %s %s", afterFin, prev, e.Op, fin))
		}
		prev = afterFin
	}

	return prev
}

func (fs ForSt) Process() string {
	var startLine = fmt.Sprintf("%d", nextNumber())

	_ = fs.InitVar.Process()
	fl := fs.Cond.Process()
	lineToFix := addLine(fmt.Sprintf("ifFalse %s goto ____", fl))
	initGotoCtx()
	addJump(lineToFix, false)
	fs.Body.Process()
	fs.UpdVar.Process()
	endLine := fmt.Sprintf("%d", nextNumber())
	releaseGotoCtx(startLine, endLine)
	return ""
}

func (ife IfElseIfSt) Process() string {

	ends := make([]int, 0)

	processIfSt(ife.If, &ends)
	sts := ife.ElseIf
	for _, s := range sts {
		processIfSt(s, &ends)
	}

	ife.Else.Body.Process()
	for _, e := range ends {
		changeLine(e, "____", fmt.Sprintf("%d", nextNumber()))
	}

	return ""
}

func processIfSt(ifSt IfSt, ends *[]int) {
	c := ifSt.Cond
	var vr string
	switch c.(type) {
	case string:
		vr = createVar()
	default:
		vr = c.(View).Process()
	}
	line := addLine(fmt.Sprintf("ifFalse %s goto ____", vr))
	ifSt.Body.Process()
	*ends = append(*ends, addLine(fmt.Sprintf("goto ____")))
	changeLine(line, "____", fmt.Sprintf("%d", nextNumber()))
}

func (wh WhileSt) Process() string {
	var startLine = fmt.Sprintf("%d", nextNumber())
	var vr string
	switch wh.BoolExprT {
	case "i", "b":
		e := wh.BoolExpr
		vr = createVar()
		addLine(fmt.Sprintf("%s = %s", vr, e))
	default:
		vr = wh.BoolExpr.(View).Process()
	}

	initGotoCtx()

	lineToFix := addLine(fmt.Sprintf("ifFalse %s goto ____", vr))
	addJump(lineToFix, false)
	wh.Body.Process()
	addLine(fmt.Sprintf("goto %s", startLine))
	endLine := fmt.Sprintf("%d", nextNumber())
	releaseGotoCtx(startLine, endLine)

	return ""
}
func (ae ArrayElem) Process() string {
	el := ae.Name
	p := ""
	if !ae.HasPos {

		switch ae.Calc.(type) {
		case string:
			vr := createVar()
			p = fmt.Sprintf("%d",addLine(fmt.Sprintf("%s = %s", vr, ae.Calc)))
		default:
			 p = ae.Calc.(View).Process()
		}

	} else {
		p = createVar()
		addLine(fmt.Sprintf("%s = %d", p, ae.Pos))
	}
	vr := createVar()
	addLine(fmt.Sprintf("%s = %s[%s]", vr, el, p))
	return vr
}
func (e ExprOperand) Process() string {
	vr := createVar()
	var r string
	switch e.T {
	case "n", "s", "i":
		r = e.V.(string)
	case "f", "a":
		r = e.V.(View).Process()
	}
	addLine(fmt.Sprintf("%s = %s", vr, r))
	return vr
}
func (e Expr) Process() string {
	left := e.Left.(View).Process()
	right := e.Right.(View).Process()
	vr := createVar()
	addLine(fmt.Sprintf("%s = %s %s %s", vr, left, e.Sign, right))

	return vr
}
func (fi FuncInvoc) Process() string {
	name := fi.Name
	count := 0
	for _, arg := range fi.Args {
		vr := createVar()
		el := arg.(View).Process()
		addLine(fmt.Sprintf("%s = param %s", vr, el))
		count++
	}
	vr := createVar()
	addLine(fmt.Sprintf("%s = call %s %d", vr, name, count))
	return vr
}
func (v Val) Process() string {
	nV := createVar()
	val := v.V
	addLine(fmt.Sprintf("%s = %s", nV, val))
	return nV
}
func (ai ArrayInit) Process() string {
	for i := 0; i < ai.Cap; i++ {
		vr := createVar()
		var res string
		switch ai.Val[i].(type) {
		case int:
			res = fmt.Sprintf("%d", ai.Val[i])
		case string:
			res = fmt.Sprintf("%s", ai.Val[i])
		}
		addLine(fmt.Sprintf("%s = init_arr %s", vr, res))
	}
	return ""
}
func (nv NewVariable) Process() string {
	nV := nv.Right.(View).Process()
	rt := nv.Type
	arr := ""
	if rt.IsArray {
		arr = "]"
	}
	addLine(fmt.Sprintf("%s = init %s%s", nv.Name, arr, rt.T))
	if nV != "" {
		addLine(fmt.Sprintf("%s = %s", nv.Name, nV))
	}
	_ = addLocal(Variable{name: nv.Name, t: nv.Type, v: nV})
	return ""
}
func (bc BreakOrContinue) Process() string {
	line := addLine(fmt.Sprintf("goto ____"))
	if bc.IsBreak {
		addJump(line, false)
	} else {
		addJump(line, true)
	}
	return ""
}

//func (sb Return) Process() string {}
func (sb StatementBody) Process() string {
	sts := sb.V

	if sts == nil {
		return ""
	}

	for _, v := range sts {
		v.(View).Process()
	}
	return ""
}

func (fd FuncReturn) Process() string {
	nx := fd.V.(View).Process()
	addLine(fmt.Sprintf("return %s",nx))
	return nx
}
func (fd FuncDefinition) Process() string {
	f := fmt.Sprintf("function %s", fd.Name)
	addLine(f)
	for _, vr := range fd.Args {
		arr := ""
		if vr.Type.IsArray {
			arr = "]"
		}
		f = fmt.Sprintf("argument %s = %s%s", vr.Name, arr, vr.Type.T)
		addLine(f)
	}
	rt := fd.ReturnType
	arr := ""
	if rt.IsArray {
		arr = "]"
	}
	f = fmt.Sprintf("return_type = %s%s", arr, rt.T)
	addLine(f)
	fd.Body.Process()
	return ""
}

type Variable struct {
	name string
	t    Type
	v    string
}
