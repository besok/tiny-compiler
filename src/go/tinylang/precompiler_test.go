package main

import (
	"testing"
)

func Test_Intermediate(t *testing.T) {

	//path := `C:\projects\tiny-compiler\src\ab\single_func.ab`
	path := `C:\projects\tiny-compiler\src\ab\calculator.ab`

	file := IR(path)

	println("file name",file.Name())

}
