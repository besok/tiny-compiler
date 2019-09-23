package env

import (
	"log"
	"testing"
	parsing "tiny-compiler/src/go/tinylang"
)

func Test_CommonTest(t *testing.T) {
	//path := `C:\projects\tiny-compiler\src\ab\single_func.ab`
	path := `C:\projects\tiny-compiler\src\ab\calculator.ab`

	file := parsing.IR(path)

	functions := ParseIR(file.Name())
	i := 0
	for _, f := range functions.List {
		log.Printf("- %s %+v %+v\n", f.Name, f.ReturnType, f.Args)
		for _, el := range f.Body {
			log.Printf("-- %#v\n", el)
			i++
		}
	}

	log.Println("total count : ",i)
}
