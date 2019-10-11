package env

import (
	"fmt"
	"log"
	"os"
	parsing "tiny-compiler/src/go/tinylang"
	"tiny-compiler/src/go/tinylang/memory"
)

var functions Functions

func Start(file string) {
	irFile := parsing.IR(file).Name()
	functions = ParseIR(irFile)
	memory.InitMemoryKb(1024)
	_ = CallFunc("main")
	_ = os.Remove(irFile)
}

func CallFunc(funcName string) interface{} {

	frame := PushFrame()

	f := findFuncByName(funcName)
	var ret interface{}
	var fin bool

	for _, el := range f.Body {
		if ret, fin = el.handle(frame); fin {
			break
		}
		log.Printf(" st : %#v",ret)
	}

	CleanTopFrame()
	return ret
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
	return st, false
}
func (st GotoSt) handle(frame *RecordTable) (interface{}, bool) {
	return st, false
}
func (st IfFalseSt) handle(frame *RecordTable) (interface{}, bool) {
	return st, false
}
