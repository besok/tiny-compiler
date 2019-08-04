package parsing


type App struct {
	functions []FuncDefinition
}

type FuncDefinition struct {
	line int
	name string
}


var app = App{functions: make([]FuncDefinition,0)}

func addFunc(f FuncDefinition) error {
	app.functions = append(app.functions,f)
	return nil
}