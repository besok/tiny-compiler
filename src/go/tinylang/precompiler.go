package parsing

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	varIdx = 0
	idx    = 0
	lines  = make([]string, 0)
	tbl    = make(map[string]Variable, 0)
)

func cleanTable() {
	tbl = make(map[string]Variable, 0)
}
func addLocal(v Variable) Variable {
	tbl[v.name] = v
	return v
}

func createVar() string {
	varIdx++
	return fmt.Sprintf("v%d", varIdx)
}
func removeVar() {
	varIdx--
}

func number() int {
	idx++
	return idx
}

func addLine(line string) {
	lines = append(lines, fmt.Sprintf("%3d:	%s", number(), line))
}

func CreateFile(f string) (*os.File, error) {
	if hasSuffix := strings.HasSuffix(f, "ab"); hasSuffix {
		file := fmt.Sprintf("%s.cab", strings.TrimSuffix(f, "ab"))
		p, e := filepath.Abs(file)
		if e == nil {
			return os.Create(p)
		} else {
			return nil, e
		}
	}
	return nil, filepath.ErrBadPattern
}

func WriteFile(f *os.File) error {
	dw := bufio.NewWriter(f)

	for _, data := range lines {
		_, _ = dw.WriteString(data + "\n")
	}

	defer f.Close()
	return dw.Flush()
}