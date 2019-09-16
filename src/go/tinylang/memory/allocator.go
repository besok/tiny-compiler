package memory

import (
	"encoding/binary"
	"log"
)

var (
	memory   []byte
	offset   int
	pointers = make([]Pointer, 0)
)

const (
	intSize  = 8
	boolSize = 1
)

func initMemoryKb(sizeKb int) {
	initMemory(sizeKb * 1024)
}

func initMemory(sizeBt int) {
	memory = make([]byte, sizeBt)
	offset = 0
	log.Printf("init memory with %d bytes", sizeBt)
}

func getBytes(p Pointer) []byte {
	st := p.offset
	fn := st + p.len
	return memory[st:fn]
}

func getBool(p Pointer) bool {
	bytes := getBytes(p)
	el := bytes[0]
	if el == 1 {
		return true
	}
	return false
}

func getInt(p Pointer) int64 {
	return int64(binary.BigEndian.Uint64(getBytes(p)))
}

func getString(p Pointer) string {
	return string(getBytes(p))
}

func getArrayString(strArray []Pointer) []string {
	checkTypeOfASrray(strArray, String)
	arr := make([]string, len(strArray))
	for i, p := range strArray {
		arr[i] = getGeneric(p).(string)
	}
	return arr
}

func getArrayBool(boolArray []Pointer) []bool {
	checkTypeOfASrray(boolArray, Bool)
	arr := make([]bool, len(boolArray))
	for i, p := range boolArray {
		arr[i] = getGeneric(p).(bool)
	}
	return arr
}

func getArrayInt(intArray []Pointer) []int64 {
	checkTypeOfASrray(intArray, Int)

	arr := make([]int64, len(intArray))
	for i, p := range intArray {
		arr[i] = getGeneric(p).(int64)
	}
	return arr

}

func checkTypeOfASrray(arr []Pointer, t PType) {
	p := arr[0]
	if p.tp != t {
		panic("wrong type")
	}
}

func getArray(pArr []Pointer) []interface{} {
	arr := make([]interface{}, len(pArr))
	for i, p := range pArr {
		arr[i] = getGeneric(p)
	}
	return arr
}

func getGeneric(p Pointer) interface{} {
	switch p.tp {
	case Int:
		return getInt(p)
	case String:
		return getString(p)
	case Bool:
		return getBool(p)
	}
	return nil
}

func putArrayBool(arr []bool) []Pointer {
	ln := len(arr)
	pointers := make([]Pointer, ln)
	for i, el := range arr {
		pointers[i] = putGeneric(el)
	}

	return pointers
}
func putArrayInt(arr []int64) []Pointer {
	ln := len(arr)
	pointers := make([]Pointer, ln)
	for i, el := range arr {
		pointers[i] = putGeneric(el)
	}

	return pointers
}
func putArrayString(arr []string) []Pointer {
	ln := len(arr)
	pointers := make([]Pointer, ln)
	for i, el := range arr {
		pointers[i] = putGeneric(el)
	}

	return pointers
}

func putGeneric(el interface{}) Pointer {
	switch el.(type) {
	case string:
		return putString(el.(string))
	case int64:
		return putInt(el.(int64))
	case bool:
		return putBool(el.(bool))
	default:
		return Pointer{len: 0, offset: -1}
	}

}

func putInt(v int64) Pointer {
	b := make([]byte, intSize)

	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	return putBytes(b, Int)
}

func putString(s string) Pointer {
	return putBytes([]byte(s), String)
}

func putBool(el bool) Pointer {
	b := make([]byte, boolSize)
	if el {
		b[0] = byte(1)
	} else {
		b[0] = byte(0)
	}
	return putBytes(b, Bool)
}

func putBytes(bt []byte, pType PType) Pointer {
	for ; !hasMemory(len(bt)); increaseMemory() {
	}

	for i, el := range bt {
		memory[i+offset] = el
	}
	p := Pointer{offset: offset, len: len(bt), tp: pType}
	addPointer(p)
	offset = offset + len(bt)

	return p
}
func hasMemory(lenIn int) bool {
	return lenIn < len(memory)-offset
}

func increaseMemory() {
	oldLen := len(memory)
	newLen := oldLen << 1

	temp := memory
	memory = make([]byte, newLen)

	for i, el := range temp {
		memory[i] = el
	}

	log.Printf("increase memory from %d to %d ", oldLen, newLen)
}

type OutOfMemoryError string

func (er *OutOfMemoryError) Error() string {
	return "OutOfMEmoryError"
}

type PType string

const (
	String PType = "STRING"
	Int    PType = "INT"
	Bool   PType = "BOOL"
)

type Pointer struct {
	len    int
	offset int
	tp     PType
}

func addPointer(p Pointer) bool {
	pointers = append(pointers, p)
	return true
}
func remPointer(p Pointer) bool {
	for i, el := range pointers {
		if el == p {
			if i == len(pointers)-1 {
				offset = offset - el.len
			}
			pointers = append(pointers[:i], pointers[i+1:]...)
			return true
		}
	}
	return false
}

func defragmentation() error {
	return nil
}

func nextFreeArea() Pointer {
	prevOffset := 0

	for _, el := range pointers {
		if el.offset-prevOffset > 1 {
			return Pointer{offset: prevOffset, len: el.offset - prevOffset}
		}
		prevOffset = countShift(el)
	}
	return Pointer{offset: -1, len: 0}
}

func countShift(p Pointer) int {
	return p.offset + p.len
}
