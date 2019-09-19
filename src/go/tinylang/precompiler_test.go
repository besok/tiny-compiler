package parsing

import (
	"testing"
)

func Test_Intermediate(t *testing.T) {

	//path := `C:\projects\tiny-compiler\src\ab\single_func.ab`
	patch := `C:\projects\tiny-compiler\src\ab\calculator.ab`

	file := CompileToInterLang(patch)

	println("file name",file.Name())

}
