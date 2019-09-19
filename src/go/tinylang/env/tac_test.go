package env

import (
	"log"
	"testing"
	parsing "tiny-compiler/src/go/tinylang"
)

func Test_CommonTest(t *testing.T){
	//path := `C:\projects\tiny-compiler\src\ab\single_func.ab`
	path := `C:\projects\tiny-compiler\src\ab\calculator.ab`

	file := parsing.CompileToInterLang(path)

	functions := Parse(file.Name())

	for _, f := range functions.Item {
		log.Printf("- %+v\n",f)
	}
}