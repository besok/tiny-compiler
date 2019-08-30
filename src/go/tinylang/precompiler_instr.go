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

type View interface {
	Process() string
}

func (uv UpdVar) Process() string {
	right := uv.Right.(View).Process()
	left := uv.Left
	if left.IsArrEl {
		addLine(fmt.Sprintf("%s[%d] = %s",left.ArrElem.Name,left.ArrElem.Pos,right))
	}else{
		addLine(fmt.Sprintf("%s = %s",left.ArrElem.Name,right))
	}
	return ""
}


func (beo BoolExprOperand) Process() string {
	if beo.IsPrim{
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



func (wh WhileSt) Process() string {
	var startLine = nextNumber()
	var vr string
	switch wh.BoolExprT {
	case "i", "b":
		e := wh.BoolExpr
		vr = createVar()
		addLine(fmt.Sprintf("%s = %s", vr, e))
	default:
		vr = wh.BoolExpr.(View).Process()
	}

	lineToFix := addLine(fmt.Sprintf("ifFalse %s goto ____", vr))

	sb := wh.Body



	addLine(fmt.Sprintf("goto %d", startLine))
	changeLine(lineToFix, "____", fmt.Sprintf("%d", nextNumber()))
	return ""
}
func (ae ArrayElem) Process() string {
	el := ae.Name
	p := ""
	if !ae.HasPos {
		p = ae.Calc.(View).Process()
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
	if bc.IsBreak{
		return "break"
	}
	return "continue"
}
func (sb StatementBody) Process() string {
	sts := sb.V
	r:=""

	for _, v := range sts {
		res := v.(View).Process()
		if res == "break" || res == "continue"{
			r = res
		}
	}
	return r
}

func (fd FuncDefinition) Process() string {
	f := fmt.Sprintf("function %s", fd.Name)
	addLine(f)
	for _, vr := range fd.Args {
		arr := ""
		if vr.Type.IsArray {
			arr = "]"
		}
		f = fmt.Sprintf("arg %s = %s%s", vr.Name, arr, vr.Type.T)
		addLine(f)
	}
	rt := fd.ReturnType
	arr := ""
	if rt.IsArray {
		arr = "]"
	}
	f = fmt.Sprintf("return type = %s%s", arr, rt.T)
	addLine(f)
	fd.Body.Process()
	addLine("")
	return ""
}

type Variable struct {
	name string
	t    Type
	v    string
}
