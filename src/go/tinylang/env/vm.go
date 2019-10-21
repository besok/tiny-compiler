package env

import (
	"fmt"
	"log"
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
	log.Printf("call func %s", funcName)
	frame := PushFrame()
	defer CleanTopFrame()

	f := findFuncByName(funcName)

	sts := f.Body
	for i := 0; i < len(sts); i++ {
		st := sts[i]
		if fn, ret, fin := handle(st, frame); fin {
			return ret
		} else {
			i = fn(i, &sts)
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

func (rt *RecordTable) putByKey(key string, isLeft bool, newRel string) bool {
	if r := rt.findByKey(key); len(r) > 0 {
		for i := range r {
			r[i].relations = append(r[i].relations, Relation{isLeft: isLeft, val: newRel})
			return true
		}
	}
	return false
}
func (rt *RecordTable) put(relOrKey string, isLeft bool, newRel string) bool {
	if done := rt.putByKey(relOrKey, isLeft, newRel); !done {
		return rt.putByRel(relOrKey, isLeft, newRel)
	}
	return true
}
func (rt *RecordTable) putByRel(rel string, isLeft bool, newRel string) bool {
	if records := rt.findByRel(rel); len(records) > 0 {
		for i := range records {
			records[i].relations = append(records[i].relations, Relation{isLeft: isLeft, val: newRel})
			return true
		}
	}
	return false
}

func (rt *RecordTable) find(key string) (*memory.Pointer, bool) {
	keyList := rt.findByKey(key)
	if len(keyList) > 0 {
		return keyList[0].pointer, true
	}
	relList := rt.findByRel(key)
	if len(relList) > 0 {
		return relList[0].pointer, true
	}
	return nil, false
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

func handle(st BodyStatement, frame *RecordTable) (func(int, *[]BodyStatement) int, interface{}, bool) {
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
	leftVar := st.Left.makeName()
	rightVar := st.Right.makeName()
	res := frame.putByRel(rightVar, true, leftVar)
	log.Printf(" put %s=%s is %t", leftVar, rightVar, res)
	return st, false
}
func (st InitPrimSt) handle(frame *RecordTable) (interface{}, bool) {
	pt := st.T
	vl := st.Val
	iv := st.Left.makeName()
	var pointer *memory.Pointer
	switch pt {
	case S:
		pointer = memory.PutString(vl.(string))
	case N:
		pointer = memory.PutInt(vl.(int64))
	case B:
		bValue := false
		if vl == "true" {
			bValue = true
		}
		pointer = memory.PutBool(bValue)
	}

	frame.init(iv, pointer)
	return nil, false
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
	iv := st.Var
	vr := iv.makeName()
	if pointer, ok := frame.find(vr); ok {
		return memory.GetGeneric(pointer), true
	}
	return nil, true
}
func (st GotoSt) handle(frame *RecordTable) (interface{}, bool) {
	return st.Goto, false
}
func (st IfFalseSt) handle(frame *RecordTable) (interface{}, bool) {
	name := st.Var.makeName()
	if pointer, ok := frame.find(name); ok {
		return CondGoto{st.Line, memory.GetBool(pointer)}, false
	}

	return CondGoto{st.Line, true}, false
}

func ident(i int, _bs *[]BodyStatement) int {
	return i
}
func gotoLine(line int) func(int, *[]BodyStatement) int {
	return func(i int, bs *[]BodyStatement) int {
		if el := findIdx(bs, line); el > 0 {
			return el
		}
		return line
	}
}

func findIdx(bs *[]BodyStatement, idx int) int {
	for i, st := range *bs {
		if st.line() == idx {
			return i
		}
	}
	return -1
}

type CondGoto struct {
	line int
	cond bool
}

func (st NewVarSt) line() int {
	return st.Line
}
func (st UpdVarSt) line() int {
	return st.Line
}
func (st InitItemSt) line() int {
	return st.Line
}
func (st InitInternalVarSt) line() int {
	return st.Line
}
func (st InitPrimSt) line() int {
	return st.Line
}
func (st ParamSt) line() int {
	return st.Line
}
func (st CallSt) line() int {
	return st.Line
}
func (st InitNumBoolOpSt) line() int {
	return st.Line
}
func (st InitArrElemSt) line() int {
	return st.Line
}
func (st ExtArrElemSt) line() int {
	return st.Line
}
func (st ReturnSt) line() int {
	return st.Line
}
func (st GotoSt) line() int {
	return st.Line
}
func (st IfFalseSt) line() int {
	return st.Line
}
