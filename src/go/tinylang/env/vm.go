package env

import (
	"fmt"
	"os"
	parsing "tiny-compiler/src/go/tinylang"
	"tiny-compiler/src/go/tinylang/memory"
)

var functions Functions

func Start(file string, irRem bool) {
	irFile := parsing.IR(file).Name()
	functions = ParseIR(irFile)
	memory.InitMemoryKb(1024)
	_ = CallFunc("main")
	if !irRem {
		_ = os.Remove(irFile)
	}
}

func CallFunc(funcName string) interface{} {

	frame := PushFrame()
	defer CleanTopFrame()

	f := findFuncByName(funcName)

	for i := 0; i < len(f.Body); i++ {
		st := f.Body[i]
		if fn, ret, fin := handle(st, frame); fin {
			return ret
		} else {
			i = fn(i)
		}
	}

	return nil

}

func PushFrame() *RecordTable {
	tbl := RecordTable{}
	frames = append(frames, tbl)
	return &tbl
}
func PopFrame() *RecordTable {
	l := len(frames)
	if l == 0 {
		return nil
	}
	ret := frames[l-1]
	ctx = ctx[:l-1]
	return &ret
}

func CleanTopFrame() {
	frame := PopFrame()
	for _, el := range *frame {
		memory.RemovePointer(el.pointer)
	}

	memory.Defragmentation()

}

func findFuncByName(name string) Function {
	for _, f := range functions.List {
		if f.Name == name {
			return f
		}
	}

	panic(fmt.Sprintf("the function %s does not exist", name))
}

type Record struct {
	key       string
	pointer   *memory.Pointer
	relations []Relation
}
type Relation struct {
	val    string
	isLeft bool
}
type RecordTable []*Record
type Frames []RecordTable

var frames Frames = make([]RecordTable, 0)

func (rt *RecordTable) init(key string, pointer *memory.Pointer) {
	*rt = append(*rt, &Record{relations: []Relation{}, key: key, pointer: pointer})
}

func (rt *RecordTable) putByKey(key string, isLeft bool, newRel string) {
	if r := rt.findByKey(key); len(r) > 0 {
		for i := range r {
			r[i].relations = append(r[i].relations, Relation{isLeft: isLeft, val: newRel})
		}
	}
}
func (rt *RecordTable) putByRel(rel string, isLeft bool, newRel string) {
	if records := rt.findByRel(rel); len(records) > 0 {
		for i := range records {
			records[i].relations = append(records[i].relations, Relation{isLeft: isLeft, val: newRel})
		}
	}
}

func (rt *RecordTable) findByKey(key string) []*Record {

	var records []*Record

	for i := range *rt {
		if (*rt)[i].key == key {
			records = append(records, (*rt)[i])
		}
	}
	return records
}
func (rt *RecordTable) findByRel(rel string) []*Record {
	var records []*Record

	for i := range *rt {
		rc := (*rt)[i]
		for i2 := range rc.relations {
			if rc.relations[i2].val == rel {
				records = append(records, rc)
			}
		}
	}
	return records
}

func handle(st BodyStatement, frame *RecordTable) (func(int) int, interface{}, bool) {
	res, fin := st.handle(frame)
	switch st.(type) {
	case GotoSt:
		return gotoLine(res.(int)), res, fin
	case IfFalseSt:
		cg := res.(CondGoto)
		if cg.cond {
			return ident, res, fin
		}
		return gotoLine(cg.line), res, fin
	default:
		return ident, res, fin
	}
}

func (st NewVarSt) handle(frame *RecordTable) (interface{}, bool) {
	return st, false
}
func (st UpdVarSt) handle(frame *RecordTable) (interface{}, bool) {
	return st, false
}
func (st InitItemSt) handle(frame *RecordTable) (interface{}, bool) {
	return st, false
}
func (st InitInternalVarSt) handle(frame *RecordTable) (interface{}, bool) {
	return st, false
}
func (st InitPrimSt) handle(frame *RecordTable) (interface{}, bool) {
	return st, false
}
func (st ParamSt) handle(frame *RecordTable) (interface{}, bool) {
	return st, false
}
func (st CallSt) handle(frame *RecordTable) (interface{}, bool) {
	return st, false
}
func (st InitNumBoolOpSt) handle(frame *RecordTable) (interface{}, bool) {
	return st, false
}
func (st InitArrElemSt) handle(frame *RecordTable) (interface{}, bool) {
	return st, false
}
func (st ExtArrElemSt) handle(frame *RecordTable) (interface{}, bool) {
	return st, false
}
func (st ReturnSt) handle(frame *RecordTable) (interface{}, bool) {
	return st, true
}
func (st GotoSt) handle(frame *RecordTable) (interface{}, bool) {
	return st.Goto, false
}
func (st IfFalseSt) handle(frame *RecordTable) (interface{}, bool) {
	return CondGoto{1,false}, false
}

func ident(i int) int {
	return i
}
func gotoLine(line int) func(int) int {
	return func(i int) int {
		return line
	}
}

type CondGoto struct {
	line int
	cond bool
}
