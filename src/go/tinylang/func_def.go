package parsing

type CtxState string

var (
	Start  CtxState = "start"
	Middle CtxState = "middle"
	Next   CtxState = "next"
	End    CtxState = "end"
)

var script = Script{functions: make([]FuncDefinition, 0)}

func AddFunc(f *FuncDefinition) {
	script.functions = append(script.functions, *f)
}

type Ctx interface {
	getState() CtxState
	setState(st CtxState)
	PutItem(ctx Ctx)
	Close()
}

type Script struct {
	functions []FuncDefinition
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
	// func args
	case Start:
		switch ctx.(type) {
		case VarCtx:
			f.F.Args = append(f.F.Args, ctx.(VarCtx).V)
		}
	case Middle:
		switch ctx.(type) {
		case VarCtx:
			f.F.ReturnType = ctx.(VarCtx).V.Type
		}
	case End:
	}
}

type FuncDefinition struct {
	Line       int
	Name       string
	Args       []Var
	ReturnType Type
	Body       []interface{}
}



func NewFuncDefinition(line int, name string) *FuncDefinition {
	return &FuncDefinition{Line: line, Name: name, Args: make([]Var, 0),Body:make([]interface{},0)}
}
