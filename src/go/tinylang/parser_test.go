package parsing

import (
	"fmt"
	"testing"
)

func Test_EnterEveryRule(t *testing.T) {

	_ = Parse(`C:\projects\tiny-compiler\src\ab\calculator.ab`)
	s := script
	fmt.Printf("%+v /n", s)

}
