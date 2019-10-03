package memory

import (
	"log"
	"testing"
)

func Test_commonMemory(t *testing.T) {

	InitMemoryKb(1)
	pB := PutBool(true)
	var intBasic int64 = 1024 * 1024 * 1024 * 1024
	pI := PutInt(intBasic)
	strBasic := "Boris, Hello World!"
	pS := PutString(strBasic)

	if GetString(pS) != strBasic {
		t.Fatalf("error for string : %s", strBasic)
	}
	if GetInt(pI) != intBasic {
		t.Fatalf("error for int : %d", intBasic)
	}
	if !GetBool(pB) {
		t.Fatalf("error for bool", )
	}

	aB := []bool{true, true, true, false, false}
	aI := []int64{1, 2, 3, 4, 5, 6}
	aS := []string{"a", "boris", "!!!"}

	pAB := PutArrayBool(aB)
	pAI := PutArrayInt(aI)
	pAS := PutArrayString(aS)

	zB := GetArrayBool(pAB)
	zI := GetArrayInt(pAI)
	zS := GetArrayString(pAS)

	if !arrCmpB(aB, zB) || !arrCmpI(aI, zI) || !arrCmpS(aS, zS) {
		t.Fatalf("error for array", )
	}
}

func Test_addAndRemPointers(t *testing.T) {
	addPointer(&Pointer{len: 1, offset: 1, tp: Bool})
	addPointer(&Pointer{len: 1, offset: 2, tp: Bool})
	addPointer(&Pointer{len: 1, offset: 3, tp: Bool})

	if len(pointers) < 3 {
		t.Fatalf("error for poitners")
	}

	ok := remPointer(&Pointer{len: 1, offset: 1, tp: Bool})

	if !ok {
		t.Fatalf("error")
	}

	if len(pointers) > 2 {
		t.Fatalf("error for pointers")
	}

	ok = remPointer(&Pointer{len: 1, offset: 10, tp: Bool})

	if ok {
		t.Fatalf("error")
	}

}

func Test_FreeArea(t *testing.T) {
	initMemory(100)
	p1 := PutString("Boris")
	_ = PutString("Boris is going ")
	_ = PutString("Boris is going home")

	remPointer(p1)

	p := nextFreeArea()

	if p.offset != 0 {
		t.Fatalf("error!")
	}

	p2 := PutString("Boris is going home!")

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

	p1 := PutString("Boris")
	p2 := PutString("Boris is going ")
	p3 := PutString(s)
	p4 := PutString(s)
	p5 := PutString(s)

	if offset != 77 {
		t.Fatalf("wrong offswet:%d", offset)
	}

	remPointer(p2)

	if offset != 77 {
		t.Fatalf("error, offset:%d", offset)
	}

	defragmentation()

	sRes3 := GetString(p3)
	sRes4 := GetString(p4)
	sRes5 := GetString(p5)

	if sRes3 != s || s != sRes4 || s != sRes5 {
		log.Fatalf("wrong shift for offset")
	}

	// the last one
	if offset != 62 {
		log.Fatalf("wrong offset %d", offset)
	}

	remPointer(p5)

	if offset != 43 {
		log.Fatalf("wrong offset %d", offset)
	}

	if sRes3 != s || s != sRes4 {
		log.Fatalf("wrong shift for offset")
	}

	remPointer(p1)
	defragmentation()

	if offset != 38 {
		log.Fatalf("wrong offset %d", offset)
	}

	if sRes3 != s || s != sRes4 {
		log.Fatalf("wrong shift for offset")
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
