package parsing

import (
	"fmt"
	"testing"
)

func Test_EnterEveryRule(t *testing.T) {

	_ = Parse(`C:\projects\tiny-compiler\src\ab\hello_world.ab`)
	fmt.Println("", app)
}
