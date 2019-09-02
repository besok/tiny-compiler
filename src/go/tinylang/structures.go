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

type CtxState string

var (
	Start  CtxState = "start"
	Middle CtxState = "middle"
	Next   CtxState = "next"
	End    CtxState = "end"
)

var script = Script{Functions: make([]FuncDefinition, 0)}

func AddFunc(f *FuncDefinition) {
	script.Functions = append(script.Functions, *f)
}

type Ctx interface {
	getState() CtxState
	setState(st CtxState)
	PutItem(ctx Ctx)
	Close()
	get() interface{}
	line() int
}

func Release(c Ctx){
	ctx, ok := Pop()
 	if ok {
		(*ctx).PutItem(c)
		Push(*ctx)
	}
}

