package parsing

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "tiny-compiler/src/antlr"
)

type App struct {
	functions []FuncDefinition
}

type FuncDefinition struct {
	line int
	name string
}

var app = App{functions: make([]FuncDefinition, 0)}

func addFunc(f FuncDefinition) error {
	app.functions = append(app.functions, f)
	return nil
}

func Parse(path string) error {

	inp, _ := antlr.NewFileStream(path)
	lexer := parser.NewTinyLangLexer(inp)
	p := parser.NewTinyLangParser(antlr.NewCommonTokenStream(lexer, 0))

	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	file := p.File()

	antlr.ParseTreeWalkerDefault.Walk(NewPrintListener(), file)

	return nil
}
