package memory

import (
	"log"
	"testing"
)

func Test_commonMemory(t *testing.T) {

	initMemoryKb(1)
	pB := putBool(true)
	var intBasic int64 = 1024 * 1024 * 1024 * 1024
	pI := putInt(intBasic)
	strBasic := "Boris, Hello World!"
	pS := putString(strBasic)

	if getString(pS) != strBasic {
		t.Fatalf("error for string : %s", strBasic)
	}
	if getInt(pI) != intBasic {
		t.Fatalf("error for int : %d", intBasic)
	}
	if !getBool(pB) {
		t.Fatalf("error for bool", )
	}

	aB := []bool{true, true, true, false, false}
	aI := []int64{1, 2, 3, 4, 5, 6}
	aS := []string{"a", "boris", "!!!"}

	pAB := putArrayBool(aB)
	pAI := putArrayInt(aI)
	pAS := putArrayString(aS)

	zB := getArrayBool(pAB)
	zI := getArrayInt(pAI)
	zS := getArrayString(pAS)

	if !arrCmpB(aB, zB) || !arrCmpI(aI, zI) || !arrCmpS(aS, zS) {
		t.Fatalf("error for array", )
	}
}

func Test_addAndRemPointers(t *testing.T) {
	addPointer(Pointer{len: 1, offset: 1, tp: Bool})
	addPointer(Pointer{len: 1, offset: 2, tp: Bool})
	addPointer(Pointer{len: 1, offset: 3, tp: Bool})

	if len(pointers) < 3 {
		t.Fatalf("error for poitners")
	}

	ok := remPointer(Pointer{len: 1, offset: 1, tp: Bool})

	if !ok {
		t.Fatalf("error")
	}

	if len(pointers) > 2 {
		t.Fatalf("error for pointers")
	}

	ok = remPointer(Pointer{len: 1, offset: 10, tp: Bool})

	if ok {
		t.Fatalf("error")
	}

}

func Test_FreeArea(t *testing.T) {
	initMemory(100)
	p1 := putString("Boris")
	_ = putString("Boris is going ")
	_ = putString("Boris is going home")

	remPointer(p1)

	p := nextFreeArea()

	if p.offset != 0 {
		t.Fatalf("error!")
	}

	p2 := putString("Boris is going home!")

	if offset != 59 {
		t.Fatalf("error, offset:%d", offset)
	}

	remPointer(p2)

	if offset != 39 {
		t.Fatalf("error, offset:%d", offset)
	}

}

func Test_Defragmentation(t *testing.T) {
	initMemory(100)
	s := "Boris is going home"

	_ = putString("Boris")
	p2 := putString("Boris is going ")
	p3 := putString(s)
	p4 := putString(s)
	_ = putString(s)

	if offset != 77 {
		t.Fatalf("error, offset:%d", offset)
	}

	remPointer(p2)

	if offset != 77 {
		t.Fatalf("error, offset:%d", offset)
	}

	defragmentation()

	sRes3 := getString(p3)
	sRes4 := getString(p4)

	if sRes3 == s || s == sRes4{
		log.Fatalf(" error , offset:%d",offset)
	}
}

func arrCmpB(left []bool, right []bool) bool {
	if len(left) == len(right) {
		for i := range left {
			if left[i] != right[i] {
				return false
			}
		}
		return true
	}
	return false
}
func arrCmpS(left []string, right []string) bool {
	if len(left) == len(right) {
		for i := range left {
			if left[i] != right[i] {
				return false
			}
		}
		return true
	}
	return false
}
func arrCmpI(left []int64, right []int64) bool {
	if len(left) == len(right) {
		for i := range left {
			if left[i] != right[i] {
				return false
			}
		}
		return true
	}
	return false
}
