package parsing

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
	parser "tiny-compiler/src/antlr"
)

func Test_EnterEveryRule(t *testing.T){

	inp,_ := antlr.NewFileStream(	`C:\projects\tiny-compiler\src\ab\calculator.ab`)
	lexer := parser.NewTinyLangLexer(inp)
	p := parser.NewTinyLangParser(antlr.NewCommonTokenStream(lexer, 0))

	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	file := p.File()

	antlr.ParseTreeWalkerDefault.Walk(NewPrintListener(), file)
	fmt.Println("",app)
}
