package parsing

var ctx = make([]Ctx, 0)

func Push(c Ctx) {
	ctx = append(ctx, c)
}
func Pop() (*Ctx, bool) {
	l := len(ctx)
	if l == 0 {
		return nil, false
	}
	ret := ctx[l-1]
	ctx = ctx[:l-1]
	return &ret, true
}
type VarCtx struct {
	V Var
	S CtxState
}

func (v VarCtx) getState() CtxState {
	return v.S
}

func (v VarCtx) setState(st CtxState) {
}

func (v VarCtx) PutItem(ctx Ctx) {
}

func (v VarCtx) Close() {
}

type Var struct {
	Name string
	Type Type
}


type Type struct {
	IsArray bool
	T       string
}

