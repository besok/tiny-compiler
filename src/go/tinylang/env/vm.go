package env

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	parsing "tiny-compiler/src/go/tinylang"
	"tiny-compiler/src/go/tinylang/memory"
)

var functions Functions

func Start(file string, irRem bool, args []string) {
	irFile := parsing.IR(file).Name()
	functions = ParseIR(irFile)
	memory.InitMemoryKb(1024)
	argsGen := make([]interface{}, 1)
	argsGen[0] = args
	_ = CallFunc("main", argsGen)
	if !irRem {
		_ = os.Remove(irFile)
	}
}

func CallFunc(funcName string, args []interface{}) interface{} {
	log.Printf("call func %s", funcName)
	frame := PushFrame()
	defer CleanTopFrame()

	f := findFuncByName(funcName)

	argNames := f.Args

	expLn := len(f.Args)
	gotLn := len(args)
	if gotLn != expLn {
		panic(fmt.Sprintf("the function %s expects to get %d args, but got %d args", funcName, expLn, gotLn))
	}

	for i, arg := range argNames {
		vl := args[i]
		nm := arg.Name
		if arg.Type.IsArray {
			pt := arg.Type.T
			switch pt {
			case S:
				vlArr := vl.([]string)
				arrPtr := memory.PutArrayString(vlArr)
				frame.initArr(nm, arrPtr)
			case B:
				vlArr := vl.([]bool)
				arrPtr := memory.PutArrayBool(vlArr)
				frame.initArr(nm, arrPtr)
			case N:
				switch vl.(type) {
				case []int64:
					vlArr := vl.([]int64)
					arrPtr := memory.PutArrayInt(vlArr)
					frame.initArr(nm, arrPtr)
				case []interface{}:
					arr := vl.([]interface{})
					if len(arr) > 0 {
						intArr := make([]int64, 0)
						for _, el := range arr {
							intArr = append(intArr, el.(int64))
						}
						arrayInt := memory.PutArrayInt(intArr)
						frame.initArr(nm,arrayInt)
					} else {
						frame.initArr(nm, make([]*memory.Pointer, 0))
					}
				}

			}
		} else {
			p := memory.PutGeneric(vl)
			frame.init(nm, p)
		}
	}

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
	frames = frames[:l-1]
	return &ret
}

func CleanTopFrame() {
	frame := PopFrame()
	for _, el := range *frame {
		if !el.value.isArray {
			memory.RemovePointer(el.value.pointer)
		} else {
			array := el.value.pointersArray
			for _, p := range array {
				memory.RemovePointer(p)
			}
		}
	}

	memory.Defragmentation()
	cleanVarTbl(frame)
}
func checkFnReturnType(name string) ReturnType {
	return findFuncByName(name).ReturnType
}
func findFuncByName(name string) Function {
	for _, f := range functions.List {
		if f.Name == name {
			return f
		}
	}

	panic(fmt.Sprintf("the function %s does not exist", name))
}

type RecordPointer struct {
	isArray       bool
	pointersArray []*memory.Pointer
	pointer       *memory.Pointer
}

func (rp *RecordPointer) cloneInto(p *RecordPointer) {
	if rp.isArray {
		memory.RemovePointers(rp.pointersArray)
		rp.isArray = p.isArray
		rp.pointersArray = p.pointersArray
	} else {
		oldPointer := rp.pointer
		memory.RemovePointer(oldPointer)
		rp.pointer = p.pointer
	}
}
func (rp *RecordPointer) put(p *memory.Pointer) {
	rp.pointersArray = append(rp.pointersArray, p)
}

type Record struct {
	key       string
	value     RecordPointer
	relations []Relation
}

type Relation struct {
	val    string
	isLeft bool
}
type RecordTable []*Record
type Frames []RecordTable

var frames Frames = make([]RecordTable, 0)

type VarType struct {
	name  string
	t     Type
	frame *RecordTable
}

type VarTable []*VarType

var varTbl = make(VarTable, 0)

func cleanVarTbl(fr *RecordTable) bool {
	tempArr := make(VarTable, 0)
	for _, el := range varTbl {
		if el.frame != fr {
			tempArr = append(tempArr, el)
		} else {
			log.Printf(" the var %s will be deleted", el.name)
		}
	}
	varTbl = tempArr
	return true
}
func addVar(name string, t Type, fr *RecordTable) bool {
	if _, ext := findVar(name, fr); ext {
		removeVar(name, fr)
	}
	varTbl = append(varTbl, &VarType{name: name, frame: fr, t: t})
	return true
}
func removeVar(name string, fr *RecordTable) bool {
	flag := false
	for i := range varTbl {
		vt := varTbl[i]
		if vt.frame == fr && vt.name == name {
			vt.name = ""
			flag = true
		}
	}
	return flag
}
func findVar(name string, fr *RecordTable) (*VarType, bool) {
	for i := range varTbl {
		vt := varTbl[i]
		if vt.frame == fr && vt.name == name {
			return vt, true
		}
	}

	return nil, false
}

func (rt *RecordTable) init(key string, pointer *memory.Pointer) {
	if rp, ext := rt.find(key); ext {
		oldPointer := rp.pointer
		memory.RemovePointer(oldPointer)
		rp.cloneInto(&RecordPointer{pointer: pointer})
	} else {
		*rt = append(*rt, &Record{relations: []Relation{}, key: key, value: RecordPointer{pointer: pointer}})
	}
}
func (rt *RecordTable) initArr(key string, pointers []*memory.Pointer) {
	if rp, ext := rt.find(key); ext {
		oldPointerArr := rp.pointersArray
		memory.RemovePointers(oldPointerArr)
		rp.cloneInto(&RecordPointer{pointersArray: pointers, isArray: true})
	} else {
		*rt = append(*rt, &Record{relations: []Relation{}, key: key, value: RecordPointer{pointersArray: pointers, isArray: true}})
	}
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
	if isLeft {
		rt.remove(newRel)
	}

	if done := rt.putByKey(relOrKey, isLeft, newRel); !done {
		return rt.putByRel(relOrKey, isLeft, newRel)
	}
	return true
}

func (rt *RecordTable) remove(newRel string) {
	if remExt := rt.remKey(newRel); !remExt {
		if remRel := rt.remRel(newRel); remRel {
			log.Printf("remove rel %s ", newRel)
		} else {
			log.Printf("remove key %s ", newRel, )
		}
	}
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

func (rt *RecordTable) printVal(key string) {
	if pointer, ext := rt.find(key); ext && !pointer.isArray {
		g := memory.GetGeneric(pointer.pointer)
		log.Printf("value for %s = %v ", key, g)
	}
}
func (rt *RecordTable) find(key string) (*RecordPointer, bool) {
	keyList := rt.findByKey(key)
	if len(keyList) > 0 {
		p := &keyList[0].value
		return p, true
	}
	relList := rt.findByRel(key)
	if len(relList) > 0 {
		p := &relList[0].value
		return p, true
	}
	return &RecordPointer{}, false
}

func (rt *RecordTable) remRel(rel string) bool {
	res := false
	for i := range *rt {
		rc := (*rt)[i]
		for i2 := range rc.relations {
			if rc.relations[i2].val == rel {
				rc.relations[i2].val = ""
				res = true
				log.Printf("rel %s has been deleted from %s", rel, rc.key)
			}
		}
	}
	return res
}

func (rt *RecordTable) remKey(key string) bool {
	for i := range *rt {
		if (*rt)[i].key == key {
			*rt = append((*rt)[:i], (*rt)[i+1:]...)
			log.Printf("key %s has been deleted ", key)
			return true
		}
	}
	return false
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
		if !cg.cond {
			return gotoLine(cg.line), res, fin
		}
		return ident, res, fin
	default:
		return ident, res, fin
	}
}
func (st NewVarSt) handle(frame *RecordTable) (interface{}, bool) {
	args := st.ArrArgs
	if st.T.IsArray {
		name := st.Name
		if _, ok := frame.find(name); ok {
			panic(fmt.Sprintf(" the var %s can not be used before init", name))
		}

		if len(args) > 0 {
			tp := st.T.T
			var pointers []*memory.Pointer
			switch tp {
			case N:
				vals := make([]int64, 0)
				for _, arg := range args {
					vals = append(vals, arg.Value.(int64))
				}
				pointers = memory.PutArrayInt(vals)
			case B:
				vals := make([]bool, 0)
				for _, arg := range args {
					vals = append(vals, arg.Value.(bool))
				}
				pointers = memory.PutArrayBool(vals)
			case S:
				vals := make([]string, 0)
				for _, arg := range args {
					vals = append(vals, arg.Value.(string))
				}
				pointers = memory.PutArrayString(vals)
			}
			log.Printf("add var %s as an array %#v", name, pointers)
			frame.initArr(name, pointers)
		} else {
			frame.initArr(name, []*memory.Pointer{})
		}
	} else {
		addVar(st.Name, st.T, frame)
	}

	return nil, false
}

func (st UpdVarSt) handle(frame *RecordTable) (interface{}, bool) {
	vr := st.Var.makeName()
	nm := st.Name
	res := frame.put(vr, true, nm)

	if vt, ok := findVar(nm, frame); ok {
		pt := vt.t
		if !pt.IsArray {
			ptype := pt.T
			p, _ := frame.find(vr)
			pointerTp := p.pointer.Type()

			switch ptype {
			case N:
				if pointerTp == memory.String {
					str := memory.GetString(p.pointer)
					if res, err := strconv.Atoi(str); err == nil {
						pInt := memory.PutInt(int64(res))

						p.cloneInto(&RecordPointer{pointer: pInt})
					} else {
						panic(fmt.Sprintf(" wrong type for %s", nm))
					}
				}
			}
		}

	}

	log.Printf("put %s=%s is %t", nm, vr, res)
	return st, false
}

func (st InitItemSt) handle(frame *RecordTable) (interface{}, bool) {
	iv := st.Var.makeName()
	item := st.Item

	ok := frame.put(item, true, iv)
	frame.printVal(iv)
	frame.printVal(item)
	log.Printf("put %s=%s is %t", iv, item, ok)
	return st, false
}
func (st InitInternalVarSt) handle(frame *RecordTable) (interface{}, bool) {
	leftVar := st.Left.makeName()
	rightVar := st.Right.makeName()
	res := frame.put(rightVar, true, leftVar)
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
		var el = int64(vl.(int))
		pointer = memory.PutInt(el)
	case B:
		bValue := false
		if vl == true {
			bValue = true
		}
		pointer = memory.PutBool(bValue)
	}
	log.Printf("init value %s and %#v", iv, pointer)
	frame.init(iv, pointer)
	return nil, false
}
func (st ParamSt) handle(frame *RecordTable) (interface{}, bool) {
	PushParam(st)
	return nil, false
}
func (st CallSt) handle(frame *RecordTable) (interface{}, bool) {
	nm := st.Count
	funNm := st.Func
	if st.IsSys {
		switch funNm {
		case "Output":
			params := TakeParams(nm)
			for _, p := range params {
				iv := p.Right
				ivNm := iv.makeName()
				if pointer, ext := frame.find(ivNm); ext {
					str := memory.GetString(pointer.pointer)
					fmt.Printf("%s \n", str)
				} else {
					panic(fmt.Sprintf("should be present %s", ivNm))
				}
			}
		case "Input":
			params := TakeParams(nm)
			if len(params) != 1 {
				panic(fmt.Sprintf("should be 1 for len present in line %d", st.Line))
			}
			ivNm := params[0].Right.makeName()
			if pointer, ext := frame.find(ivNm); ext {
				str := memory.GetString(pointer.pointer)
				iv := st.Var.makeName()
				fmt.Printf("%s", str)
				reader := bufio.NewReader(os.Stdin)
				txt, _ := reader.ReadString('\n')
				txt = strings.TrimSuffix(txt, "\n")
				pointer := memory.PutString(txt)
				frame.init(iv, pointer)
			} else {
				panic(fmt.Sprintf("should be present %s", ivNm))
			}
		case "Len":
			params := TakeParams(nm)
			if nm != 1 {
				panic(fmt.Sprintf("should be 1 for len present in line %d", st.Line))
			}
			ivParNm := params[0].Right.makeName()
			if pointer, ext := frame.find(ivParNm); ext {
				var lenOfArr = int64(len(pointer.pointersArray))
				ivNm := st.Var.makeName()
				p := memory.PutInt(lenOfArr)
				frame.init(ivNm, p)
			} else {
				panic(fmt.Sprintf("should be present %s", ivParNm))
			}
		case "Add":
			params := TakeParams(nm)
			if nm != 2 {
				panic(fmt.Sprintf("should be 2 for add present in line %d", st.Line))
			}
			arr := params[0].Right.makeName()
			arg := params[1].Right.makeName()

			if pArr, ext := frame.find(arr); ext {
				if pArg, ext := frame.find(arg); ext {
					p := pArg.pointer
					pArr.put(p)
				} else {
					panic(fmt.Sprintf("should be present %d", st.Line))
				}
			} else {
				panic(fmt.Sprintf("should be present %d", st.Line))
			}
		}

	} else {
		params := TakeParams(nm)
		args := make([]interface{}, 0)
		for _, pr := range params {
			nmIv := pr.Right.makeName()
			if p, ext := frame.find(nmIv); ext {
				if p.isArray {
					vals := make([]interface{}, 0)
					pArrs := p.pointersArray
					for _, p := range pArrs {
						gen := memory.GetGeneric(p)
						vals = append(vals, gen)
					}
					args = append(args, vals)
				} else {
					res := memory.GetGeneric(p.pointer)
					args = append(args, res)
				}
			} else {
				panic(fmt.Sprintf("the param %s for func %s does not exist, line: %d", nmIv, funNm, st.Line))
			}
		}

		res := CallFunc(funNm, args)
		iv := st.Var.makeName()
		rt := checkFnReturnType(funNm)

		if rt.IsVoid {
			frame.init(iv, nil)
		} else if rt.Type.IsArray {
			genArr := res.([]interface{})
			if len(genArr) > 0 {
				el := genArr[0]
				temp := make([]*memory.Pointer, 0)
				switch el.(type) {
				case int64:
					for _, e := range genArr {
						p := memory.PutInt(e.(int64))
						temp = append(temp, p)
					}
				case string:
					for _, e := range genArr {
						p := memory.PutString(e.(string))
						temp = append(temp, p)
					}
				case bool:
					for _, e := range genArr {
						p := memory.PutBool(e.(bool))
						temp = append(temp, p)
					}
				}

				frame.initArr(iv, temp)
			} else {
				frame.initArr(iv, make([]*memory.Pointer, 0))
			}

		} else {
			genP := memory.PutGeneric(res)
			frame.init(iv, genP)
		}

	}

	return nil, false
}
func (st InitNumBoolOpSt) handle(frame *RecordTable) (interface{}, bool) {
	toInit := st.Left.makeName()
	leftOp := st.RightFirst.makeName()
	rightOp := st.RightSecond.makeName()

	leftP, okLeft := frame.find(leftOp)
	rightP, okRight := frame.find(rightOp)

	if !okLeft || !okRight {
		panic(fmt.Sprintf("should be present %s ", toInit))
	}

	if st.IsBool {
		sign := st.Sign
		switch sign {
		case "||":
			bLeft := memory.GetBool(leftP.pointer)
			bRight := memory.GetBool(rightP.pointer)
			newPointer := memory.PutBool(bLeft || bRight)
			frame.init(toInit, newPointer)
		case "&&":
			bLeft := memory.GetBool(leftP.pointer)
			bRight := memory.GetBool(rightP.pointer)
			newPointer := memory.PutBool(bLeft && bRight)
			frame.init(toInit, newPointer)
		case "==":
			bLeft := memory.GetGeneric(leftP.pointer)
			bRight := memory.GetGeneric(rightP.pointer)
			newPointer := memory.PutBool(bLeft == bRight)
			frame.init(toInit, newPointer)
		case "!=":
			bLeft := memory.GetGeneric(leftP.pointer)
			bRight := memory.GetGeneric(rightP.pointer)
			newPointer := memory.PutBool(bLeft != bRight)
			frame.init(toInit, newPointer)
		case ">", "<":
			bLeft := memory.GetGeneric(leftP.pointer)
			bRight := memory.GetGeneric(rightP.pointer)
			switch bLeft.(type) {
			case string, bool:
				log.Printf("can not possible to compare with > or < , %s", toInit)
				panic(fmt.Sprintf("can not possible to compare with > or < , line %d", st.Line))
			case int64:
				l := bLeft.(int64)
				r := bRight.(int64)
				if st.Sign == ">" {
					newPointer := memory.PutBool(l > r)
					frame.init(toInit, newPointer)
				} else {
					newPointer := memory.PutBool(l < r)
					frame.init(toInit, newPointer)
				}
			}
		}
	} else {
		bLeft := memory.GetGeneric(leftP.pointer)
		bRight := memory.GetGeneric(rightP.pointer)
		var p *memory.Pointer
		switch st.Sign {
		case "+":
			switch bLeft.(type) {
			case string:
				s := ""
				switch bRight.(type) {
				case string:
					s = bRight.(string)
				case int64:
					s = strconv.FormatInt(bRight.(int64), 10)
				case bool:
					s = strconv.FormatBool(bRight.(bool))
				}
				str := fmt.Sprintf("%s%s", bLeft.(string), s)
				p = memory.PutString(str)
				frame.init(toInit, p)
			case bool:
				panic(fmt.Sprintf("should be or string or int %d ", st.Line))
			case int64:
				res := bLeft.(int64) + bRight.(int64)
				p = memory.PutInt(res)
			}
		case "-":
			switch bLeft.(type) {
			case string, bool:
				panic(fmt.Sprintf("should be or string or int %d ", st.Line))
			case int64:
				res := bLeft.(int64) - bRight.(int64)
				p = memory.PutInt(res)
			}
		case "*":
			switch bLeft.(type) {
			case string, bool:
				panic(fmt.Sprintf("should be or string or int %d ", st.Line))
			case int64:
				res := bLeft.(int64) * bRight.(int64)
				p = memory.PutInt(res)
			}
		case "/":
			switch bLeft.(type) {
			case string, bool:
				panic(fmt.Sprintf("should be or string or int %d ", st.Line))
			case int64:
				res := bLeft.(int64) / bRight.(int64)
				p = memory.PutInt(res)
			}
		case "%":
			switch bLeft.(type) {
			case string, bool:
				panic(fmt.Sprintf("should be  int %d ", st.Line))
			case int64:
				res := bLeft.(int64) % bRight.(int64)
				p = memory.PutInt(res)
			}
		}
		frame.init(toInit, p)
		frame.printVal(toInit)
	}
	return st, false
}
func (st InitArrElemSt) handle(frame *RecordTable) (interface{}, bool) {
	toInit := st.Left.makeName()
	toFind := st.Right.makeName()
	arr := st.Name
	if p, ext := frame.find(toFind); ext {
		if pArr, extArr := frame.find(arr); extArr {
			nm := memory.GetInt(p.pointer)
			p := pArr.pointersArray[nm]
			generic := memory.GetGeneric(p)
			newPointer := memory.PutGeneric(generic)
			frame.init(toInit, newPointer)
		} else {
			panic(fmt.Sprintf("should exist %s", toInit))
		}

	} else {
		panic(fmt.Sprintf("should exist %s", toInit))
	}

	return st, false
}
func (st ExtArrElemSt) handle(frame *RecordTable) (interface{}, bool) {
	arr := st.Name
	nm := 0
	if st.IsNum {
		nm = st.Num
	} else {
		name := st.Left.makeName()
		if pointer, ext := frame.find(name); ext {
			nm = int(memory.GetInt(pointer.pointer))
		} else {
			panic(fmt.Sprintf("should exist array el %s", arr))
		}
	}

	pointer, ext := frame.find(arr)
	if !ext {
		panic(fmt.Sprintf("should exist array el %s", arr))
	}

	if len(pointer.pointersArray) < nm+1 {
		panic(fmt.Sprintf("should exist array el %s", arr))
	}

	p := pointer.pointersArray[nm]

	rightP, ext := frame.find(st.Right.makeName())
	if !ext {
		panic(fmt.Sprintf("should exist array el %s", arr))
	}

	p.Replace(rightP.pointer)

	return st, false
}
func (st ReturnSt) handle(frame *RecordTable) (interface{}, bool) {
	iv := st.Var
	vr := iv.makeName()
	if pointer, ok := frame.find(vr); ok {
		log.Printf("return %#v", pointer)
		if pointer.isArray {
			res := make([]interface{}, 0)
			for _, p := range pointer.pointersArray {
				v := memory.GetGeneric(p)
				res = append(res, v)
				return res, true
			}
		} else {
			return memory.GetGeneric(pointer.pointer), true
		}
	}
	log.Printf("nothing to return")
	return nil, true
}
func (st GotoSt) handle(frame *RecordTable) (interface{}, bool) {
	return st.Goto, false
}
func (st IfFalseSt) handle(frame *RecordTable) (interface{}, bool) {
	name := st.Var.makeName()
	if pointer, ok := frame.find(name); ok {
		return CondGoto{st.Goto, memory.GetBool(pointer.pointer)}, false
	}

	return CondGoto{st.Line, true}, false
}

func ident(i int, _bs *[]BodyStatement) int {
	return i
}
func gotoLine(line int) func(int, *[]BodyStatement) int {
	return func(i int, bs *[]BodyStatement) int {
		if el := findIdx(bs, line); el > 0 {
			log.Printf("goto to %d", el)
			return el - 1
		}
		log.Printf("goto to %d", line)
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

var paramStack = make([]ParamSt, 0)

func PushParam(param ParamSt) {
	paramStack = append(paramStack, param)
}
func TakeParams(nm int) []ParamSt {
	if len(paramStack) < nm {
		panic("a wrong signature")
	}

	params := paramStack[len(paramStack)-nm:]

	paramStack = paramStack[:len(paramStack)-nm]

	return params
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
