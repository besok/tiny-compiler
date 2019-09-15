package memory

import (
	"encoding/binary"
)

var (
	memory []byte
	offset int
)

const (
	intSize        = 8
	boolSize       = 1
	String   PType = "STRING"
	Int      PType = "INT"
	Bool     PType = "BOOL"
)

func initMemory(sizeKb int) {
	memory = make([]byte, sizeKb*1024)
	offset = 0
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
func putArrayInt(arr []int) []Pointer {
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
	case int:
		return putInt(el.(int))
	case bool:
		return putBool(el.(bool))
	default:
		return Pointer{len: 0, offset: -1}
	}

}

func putInt(v int) Pointer {
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
	offset = offset + len(bt)

	return p
}
func hasMemory(lenIn int) bool {
	if lenIn < len(memory)-offset {
		return true
	}

	return false

}

func increaseMemory() {
	oldLen := len(memory)
	newLen := oldLen << 1

	temp := memory
	memory = make([]byte, newLen)

	for i, el := range temp {
		memory[i] = el
	}

}

type OutOfMemoryError string

func (er *OutOfMemoryError) Error() string {
	return "OutOfMEmoryError"
}

type PType string

type Pointer struct {
	len    int
	offset int
	tp     PType
}
