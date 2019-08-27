package parsing

import (
	"fmt"
)

type View interface {
	Process() string
}

func (ae ArrayElem) Process() string {
	el := ae .Name
	p :=""
	if !ae.HasPos{
		p = ae.Calc.(View).Process()
	}else{
		p = createVar()
		addLine(fmt.Sprintf("%s = %d",p,ae.Pos))
	}
	vr := createVar()
	addLine(fmt.Sprintf("%s = %s[%s]",vr,el,p))
	return vr
}
func (e Expr) Process() string {
	return ""
}
func (fi FuncInvoc) Process() string {
	name := fi.Name
	count := 0
	for _, arg := range fi.Args {
		vr := createVar()
		el := arg.(View).Process()
		addLine(fmt.Sprintf("%s = param %s",vr, el))
		count++
	}
	vr := createVar()
	addLine(fmt.Sprintf("%s = call %s %d",vr, name, count))
	return vr
}
func (v Val) Process() string {
	nV := createVar()
	val := v.V
	switch v.T {
	case "s":
		val = val.(string)[1 : len(val.(string))-1]
	}
	addLine(fmt.Sprintf("%s = %s", nV, val))
	return nV
}
func (nv NewVariable) Process() string {
	nV := nv.Right.(View).Process()
	rt := nv.Type
	arr := ""
	if rt.IsArray {
		arr = "]"
	}
	addLine(fmt.Sprintf("%s = init %s%s", nv.Name, arr, rt.T))
	addLine(fmt.Sprintf("%s = %s", nv.Name, nV))
	_ = addLocal(Variable{name: nv.Name, t: nv.Type, v: nV})
	return ""
}
func (sb StatementBody) Process() string {
	sts := sb.V
	for _, v := range sts {
		_ = v.(View).Process()
	}
	return ""
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
