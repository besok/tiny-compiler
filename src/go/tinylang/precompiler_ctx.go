package parsing

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)


type VT string

const (
	Str VT = "_s"
	Num VT = "_n"
	Bool VT = "_b"
	Ptr VT = "_p"
)


var (
	varIdx = map[VT]int{Str:0,Num:0,Bool:0,Ptr:0}
	idx    = 0
	lines  = make([]Line, 0)
	tbl    = make(map[string]Variable, 0)
	patch = "____"
)

func cleanTable() {
	tbl = make(map[string]Variable, 0)
}
func addLocal(v Variable) Variable {
	tbl[v.name] = v
	return v
}

func createVar(vt VT) string {
	varIdx[vt]++
	el := varIdx[vt]
	return fmt.Sprintf("%s%d",string(vt), el)
}
func removeVar(vt VT) {
	varIdx[vt]--
}

func number() int {
	idx++
	return idx
}
func nextNumber() int {
	return idx + 1
}

func addLine(line string) int {
	i := number()
	lines = append(lines, Line{v: fmt.Sprintf("%4d:  %s", i, line), n: i})
	return i
}

func changeLine(number int, trg string) {
	for i := 0; i < len(lines); i++ {

		line := &lines[i]
		if line.n == number {
			line.v = strings.ReplaceAll(line.v, patch, trg)
			return
		}
	}
}

func CreateFile(f string) (*os.File, error) {
	if hasSuffix := strings.HasSuffix(f, ".ab"); hasSuffix {
		file := fmt.Sprintf("%s.cab", strings.TrimSuffix(f, ".ab"))
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

var gotoStack = make([]gotoCtx, 0)

func addJump(line int, toStart bool) {
	ctx := &gotoStack[len(gotoStack)-1]
	*ctx = append(*ctx, gotoJump{line: line, toStart: toStart})
}

func initGotoCtx() {
	gotoStack = append(gotoStack, make(gotoCtx, 0))
}

func releaseGotoCtx(startLine string, endLine string) {
	ctx := gotoStack[len(gotoStack)-1]
	for _, j := range ctx {
		if j.toStart {
			changeLine(j.line, startLine)
		} else {
			changeLine(j.line, endLine)
		}
	}
	gotoStack = gotoStack[:len(gotoStack)-1]
}

type gotoCtx []gotoJump

type gotoJump struct {
	line    int
	toStart bool
}
