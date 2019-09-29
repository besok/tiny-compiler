package env

import (
	"fmt"
	"log"
	parsing "tiny-compiler/src/go/tinylang"
)

var functions Functions

func Start(file string) {
	functions = ParseIR(parsing.IR(file).Name())
}

func CallFunc(funcName string) interface{}{

	f := findFuncByName(funcName)
	var ret interface{}

	for _,el :=range f.Body{
		log.Printf(" %#v",el)
	}

	return ret
}




func findFuncByName(name string) Function {
	for _, f := range functions.List {
		if f.Name == name {
			return f
		}
	}

	panic(fmt.Sprintf("the function %s does not exist",name))
}
