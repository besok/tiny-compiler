package parsing

type Script struct {
	Functions []FuncDefinition
}

type FuncDefinition struct {
	Line       int
	Name       string
	Args       []Var
	ReturnType Type
	Body       StatementBody
	Return     interface{}
}

func NewFuncDefinition(line int, name string) *FuncDefinition {
	return &FuncDefinition{Line: line, Name: name, Args: make([]Var, 0)}
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
func (f *FuncDefinitionCtx) PutItem(ctx Ctx) {

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
		f.F.Body = variable.(StatementBody)
	case End:
		f.F.Return = ctx.get()
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

func (c NewVariable) getState() CtxState {
	return Start
}

func (c NewVariable) setState(st CtxState) {
}

func (c *NewVariable) PutItem(ctx Ctx) {
	c.Right = ctx.get()
}

func (c NewVariable) Close() {
	parent, _ := Pop()
	(*parent).PutItem(&c)
	Push(*parent)
}
func (c NewVariable) get() interface{} {
	return c
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
	return v
}

type ArrayInit struct {
	T   Type
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
	f.Args = append(f.Args, ctx.get())
}

func (f FuncInvoc) Close() {
}

func (f FuncInvoc) get() interface{} {
	return f
}

type ArrayElem struct {
	Name   string
	Pos    int
	HasPos bool
	Calc   interface{}
}

func (a ArrayElem) getState() CtxState {
	return Start
}

func (a ArrayElem) setState(st CtxState) {

}

func (a *ArrayElem) PutItem(ctx Ctx) {
	a.Calc = ctx.get()
}

func (a ArrayElem) Close() {

}

func (a ArrayElem) get() interface{} {
	return a
}

type UpdVar struct {
	Line int
	Left struct {
		IsArrEl bool
		ArrElem ArrayElem
	}
	Right interface{}
}

type UpdVarCtx struct {
	V UpdVar
	S CtxState
}

func (v UpdVarCtx) getState() CtxState {
	return v.S
}

func (v UpdVarCtx) setState(st CtxState) {
	v.S = st
}

func (v *UpdVarCtx) PutItem(ctx Ctx) {
	switch v.S {
	case Start:
		v.V.Left.ArrElem = ctx.get().(ArrayElem)
	case End:
		v.V.Right = ctx.get()
	}
	v.S = End
}

func (v UpdVarCtx) Close() {
}

func (v UpdVarCtx) get() interface{} {
	return v.V
}

type Expr struct {
	Line  int
	Left  interface{}
	Right interface{}
	Sign  string
}
type ExprCtx struct {
	S CtxState
	E Expr
}

func (e *ExprCtx) getState() CtxState {
	return e.S
}

func (e *ExprCtx) setState(st CtxState) {
	e.S = st
}

func (e *ExprCtx) PutItem(ctx Ctx) {
	switch e.S {
	case Start:
		e.E.Left = ctx.get()
	case End:
		e.E.Right = ctx.get()
	}

	e.S = End
}

func (e *ExprCtx) Close() {
}

func (e *ExprCtx) get() interface{} {
	return e.E
}

type ExprOperand struct {
	T string
	V interface{}
}

func (op *ExprOperand) getState() CtxState {
	return Start
}

func (op *ExprOperand) setState(st CtxState) {
}

func (op *ExprOperand) PutItem(ctx Ctx) {
	switch ctx.(type) {
	case *FuncInvoc:
		op.V = ctx.get()
		op.T = "f"
	case *ArrayElem:
		op.V = ctx.get()
		op.T = "a"
	}
}

func (op *ExprOperand) Close() {
}

func (op *ExprOperand) get() interface{} {
	return *op
}

type BoolExpr struct {
	Line     int
	BoolExpr []BoolExprSingleExtra
}

func (b *BoolExpr) getState() CtxState {
	return Start
}

func (b *BoolExpr) setState(st CtxState) {
}

func (b *BoolExpr) PutItem(ctx Ctx) {
	switch ctx.(type) {
	case *BoolExprSingleCtx:
		boolExpr := (ctx.get()).(BoolExprSingle)
		b.BoolExpr = append(b.BoolExpr, BoolExprSingleExtra{Op: "", V: boolExpr})
	case *BoolExprSingleExtra:
		b.BoolExpr = append(b.BoolExpr, (ctx.get()).(BoolExprSingleExtra))

	}
}

func (b *BoolExpr) Close() {
}

func (b *BoolExpr) get() interface{} {
	return *b
}

type BoolExprSingleExtra struct {
	Op string
	V  BoolExprSingle
}

func (o *BoolExprSingleExtra) getState() CtxState {
	return Start
}

func (o *BoolExprSingleExtra) setState(st CtxState) {
}

func (o *BoolExprSingleExtra) PutItem(ctx Ctx) {
	o.V = ctx.get().(BoolExprSingle)
}

func (o *BoolExprSingleExtra) Close() {
}

func (o *BoolExprSingleExtra) get() interface{} {
	return *o
}

type BoolExprSingle struct {
	Sign  string
	Left  interface{}
	Right interface{}
}

type BoolExprSingleCtx struct {
	V BoolExprSingle
	S CtxState
}

func (b *BoolExprSingleCtx) getState() CtxState {
	return b.S
}

func (b *BoolExprSingleCtx) setState(st CtxState) {
	b.S = st
}

func (b *BoolExprSingleCtx) PutItem(ctx Ctx) {
	switch b.S {
	case Start:
		b.V.Left = ctx.get()
	case End:
		b.V.Right = ctx.get()
	}
	b.S = End
}

func (b *BoolExprSingleCtx) Close() {
}

func (b *BoolExprSingleCtx) get() interface{} {
	return b.V
}

type BoolExprOperand struct {
	IsPrim bool
	V      interface{}
}

func (b *BoolExprOperand) getState() CtxState {
	return Start
}

func (b *BoolExprOperand) setState(st CtxState) {
}

func (b *BoolExprOperand) PutItem(ctx Ctx) {
	b.V = ctx.get()
}

func (b *BoolExprOperand) Close() {
}

func (b *BoolExprOperand) get() interface{} {
	return *b
}

type StatementBody struct {
	V []interface{}
}

func (b *StatementBody) getState() CtxState {
	return Start
}

func (b *StatementBody) setState(st CtxState) {
}

func (b *StatementBody) PutItem(ctx Ctx) {
	b.V = append(b.V, ctx.get())
}

func (b *StatementBody) Close() {
}

func (b *StatementBody) get() interface{} {
	return *b
}

type BreakOrContinue struct {
	IsBreak bool
}

func (BreakOrContinue) getState() CtxState {
	return Start
}

func (BreakOrContinue) setState(st CtxState) {
}

func (BreakOrContinue) PutItem(ctx Ctx) {
}

func (BreakOrContinue) Close() {
}

func (c BreakOrContinue) get() interface{} {
	return c
}

type WhileSt struct {
	Line     int
	BoolExpr interface{}
	BoolExprT string
	Body     StatementBody
}
type WhileStCtx struct {
	W WhileSt
	S CtxState
}

func (w *WhileStCtx) getState() CtxState {
	return w.S
}

func (w *WhileStCtx) setState(st CtxState) {
	w.S = st
}

func (w *WhileStCtx) PutItem(ctx Ctx) {
	switch w.S {
	case Start:
		w.W.BoolExpr = ctx.get()
	case End:
		w.W.Body = ctx.get().(StatementBody)
	}
	w.setState(End)
}

func (w *WhileStCtx) Close() {
}

func (w *WhileStCtx) get() interface{} {
	return w.W
}

type ForSt struct {
	Line    int
	InitVar UpdVar
	UpdVar  UpdVar
	Cond    BoolExpr
	Body    StatementBody
}

type ForStCtx struct {
	F ForSt
	S CtxState
}

func (f *ForStCtx) getState() CtxState {
	return f.S
}

func (f *ForStCtx) setState(st CtxState) {
	f.S = st
}

func (f *ForStCtx) PutItem(ctx Ctx) {
	switch f.S {
	case Start:
		f.F.InitVar = ctx.get().(UpdVar)
		f.S = Next
	case Next:
		f.F.Cond = ctx.get().(BoolExpr)
		f.S = Middle
	case Middle:
		f.F.UpdVar = ctx.get().(UpdVar)
		f.S = End
	case End:
		f.F.Body = ctx.get().(StatementBody)
	}
}

func (f *ForStCtx) Close() {
}

func (f *ForStCtx) get() interface{} {
	return f.F
}

type IfSt struct {
	Cond interface{}
	Body StatementBody
}
type IfStCtx struct {
	V IfSt
	S CtxState
}

func (i *IfStCtx) getState() CtxState {
	return i.S
}

func (i *IfStCtx) setState(st CtxState) {
	i.S = st
}

func (i *IfStCtx) PutItem(ctx Ctx) {
	switch i.S {
	case Start:
		i.V.Cond = ctx.get()
	case End:
		i.V.Body = ctx.get().(StatementBody)
	}
	(*i).setState(End)
}

func (i *IfStCtx) Close() {
}

func (i *IfStCtx) get() interface{} {
	return i.V
}

type ElseSt struct {
	Body StatementBody
}

func (e *ElseSt) getState() CtxState {
	return Start
}

func (e *ElseSt) setState(st CtxState) {
}

func (e *ElseSt) PutItem(ctx Ctx) {
	e.Body = ctx.get().(StatementBody)
}

func (e *ElseSt) Close() {
}

func (e *ElseSt) get() interface{} {
	return *e
}

type IfElseIfSt struct {
	Line   int ``
	If     IfSt
	ElseIf []IfSt
	Else   ElseSt
}
type IfElseIfStCtx struct {
	S CtxState
	V IfElseIfSt
}

func (e *IfElseIfStCtx) getState() CtxState {
	return e.S
}

func (e *IfElseIfStCtx) setState(st CtxState) {
	e.S = st
}

func (e *IfElseIfStCtx) PutItem(ctx Ctx) {
	switch e.S {
	case Start:
		e.V.If = ctx.get().(IfSt)
		e.setState(Middle)
	case Middle:
		e.V.ElseIf = append(e.V.ElseIf, ctx.get().(IfSt))
	case End:
		e.V.Else = ctx.get().(ElseSt)

	}
}

func (e *IfElseIfStCtx) Close() {
}

func (e *IfElseIfStCtx) get() interface{} {
	return e.V
}
