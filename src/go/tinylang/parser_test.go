package parsing

import (
	"fmt"
	"testing"
)

func Test_EnterEveryRule(t *testing.T) {

	path := `C:\projects\tiny-compiler\src\ab\calculator.ab`
	s := Parse(path)
	fmt.Printf("%+v /n", s)

}
