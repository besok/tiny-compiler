package parsing

type Script struct {
	functions []FuncDefinition
}

type FuncDefinition struct {
	Line       int
	Name       string
	Args       []Var
	ReturnType Type
	Body       []interface{}
}

func NewFuncDefinition(line int, name string) *FuncDefinition {
	return &FuncDefinition{Line: line, Name: name, Args: make([]Var, 0), Body: make([]interface{}, 0)}
}

type FuncDefinitionCtx struct {
	F *FuncDefinition
	S CtxState
}

func (f FuncDefinitionCtx) getState() CtxState {
	return f.S
}

func (f *FuncDefinitionCtx) setState(st CtxState) {
	f.S = st
}

func (f FuncDefinitionCtx) Close() {
	AddFunc(f.F)
}
func (f FuncDefinitionCtx) PutItem(ctx Ctx) {

	switch f.S {
	case Start:
		switch ctx.(type) {
		case Var:
			f.F.Args = append(f.F.Args, ctx.(Var))
		}
	case Middle:
		switch ctx.(type) {
		case Var:
			f.F.ReturnType = ctx.(Var).Type
		}
	case Next:
			variable := ctx.get()
			f.F.Body = append(f.F.Body, variable)
	}
}

func (f FuncDefinitionCtx) get() interface{} {
	return *(f.F)
}

type NewVariable struct {
	Line  int
	Name  string
	Type  Type
	Right interface{}
}

type NewVariableCtx struct {
	V *NewVariable
	S CtxState
}

func (c NewVariableCtx) getState() CtxState {
	return c.S
}

func (c NewVariableCtx) setState(st CtxState) {
	c.S = st
}

func (c NewVariableCtx) PutItem(ctx Ctx) {
		c.V.Right = ctx.get()
}

func (c NewVariableCtx) Close() {
	parent, _ := Pop()
	(*parent).PutItem(c)
	Push(*parent)
}
func (c NewVariableCtx) get() interface{} {
	return *(c.V)
}


func (v Var) getState() CtxState {
	return Start
}

func (v Var) setState(st CtxState) {
}

func (v Var) PutItem(ctx Ctx) {
}

func (v Var) Close() {}
func (v Var) get() interface{} {
	return v
}


type Var struct {
	Name string
	Type Type
}


type Type struct {
	IsArray bool
	T       string
}

type Val struct {
	T string
	V interface{}
}

func (v Val) getState() CtxState {
	return Start
}

func (v Val) setState(st CtxState) {}

func (v Val) PutItem(ctx Ctx) {}

func (v Val) Close() {}

func (v Val) get() interface{} {
	return v.V
}


type ArrayInit struct {
	T Type
	Cap int
	Val []interface{}
}

func (l ArrayInit) getState() CtxState {
	return Start
}

func (l ArrayInit) setState(st CtxState) {
}

func (l ArrayInit) PutItem(ctx Ctx) {
}

func (l ArrayInit) Close() {
}

func (l ArrayInit) get() interface{} {
	return l
}

type FuncInvoc struct {
	Line int
	Name string
	Args []interface{}
}

func (f FuncInvoc) getState() CtxState {
	return Start
}

func (f FuncInvoc) setState(st CtxState) {
}

func (f *FuncInvoc) PutItem(ctx Ctx) {
	f.Args = append(f.Args,ctx.get())
}

func (f FuncInvoc) Close() {
}

func (f FuncInvoc) get() interface{} {
	return f
}

type ArrayElem struct {
	Name string
	Pos int
	HasPos bool
	Calc interface{}
}

func (a ArrayElem) getState() CtxState {
	return Start
}

func (a ArrayElem) setState(st CtxState) {

}

func (a ArrayElem) PutItem(ctx Ctx) {
	a.Calc = ctx.get()
}

func (a ArrayElem) Close() {

}

func (a ArrayElem) get() interface{} {
	return a
}


