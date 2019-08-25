package parsing

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	idx   = 0
	lines = make([]string, 0)
)

func number() int {
	idx++
	return idx
}

func addLine(line string) {
	lines = append(lines, fmt.Sprintf("%d:%s", number(), line))
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

func Process(item interface{}) {
	switch item.(type) {
	case FuncDefinition:
		processFuncDefinition(item.(FuncDefinition))
	}
}

func processFuncDefinition(fd FuncDefinition) {
	f := fmt.Sprintf("function=%s", fd.Name)
	addLine(f)
	for _, vr := range fd.Args {
		arr := ""
		if vr.Type.IsArray {
			arr = "]"
		}
		f = fmt.Sprintf("arg %s=%s%s", vr.Name, arr, vr.Type.T)
		addLine(f)
	}
	rt := fd.ReturnType
	arr := ""
	if rt.IsArray {
		arr = "]"
	}
	f = fmt.Sprintf("return %s%s", arr, rt.T)
	addLine(f)
}


