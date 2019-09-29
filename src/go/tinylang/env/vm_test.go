package env

import "testing"

func TestStart(t *testing.T) {
	path := `C:\projects\tiny-compiler\src\ab\calculator.ab`
	Start(path)
}