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
	lines  = make([]Line, 0)
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
func nextNumber() int {
	return idx+1
}


func addLine(line string) int {
	i := number()
	lines = append(lines, Line{v:fmt.Sprintf("%3d:	%s", i, line),n:i})
	return i
}

func changeLine(number int, src string, trg string){
	for i:=0;i<len(lines);i++{

		line := &lines[i]
		if line.n == number{
			line.v = strings.ReplaceAll(line.v,src,trg)
			return
		}
	}
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
		_, _ = dw.WriteString(data.v + "\n")
	}

	defer f.Close()
	return dw.Flush()
}

type Line struct {
	v string
	n int
}