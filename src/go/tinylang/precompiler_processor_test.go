package parsing

import (
	"testing"
)

func Test_Intermediate(t *testing.T) {

	path := `C:\projects\tiny-compiler\src\ab\single_func.ab`
	//path := `C:\projects\tiny-compiler\src\ab\calculator.ab`
	s := Parse(path)

	for _,f := range s.Functions{
		f.Process()
	}

	file, _ := CreateFile(path)
	_ = WriteFile(file)

}
