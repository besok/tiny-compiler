package env

import (
	"testing"
	"tiny-compiler/src/go/tinylang/memory"
)

func TestStart(t *testing.T) {
	path := `C:\projects\tiny-compiler\src\ab\calculator.ab`
	Start(path,true)
}


func TestFrame(t *testing.T) {
	tbl := &RecordTable{}

	tbl.init("1",&memory.Pointer{})
	tbl.init("2",&memory.Pointer{})
	tbl.init("3",&memory.Pointer{})
	tbl.init("4",&memory.Pointer{})

	tbl.putByKey("1",true,"_1")
	tbl.putByKey("1",true,"_2")
	tbl.putByKey("1",true,"_3")

	key := tbl.findByKey("1")

	if len(key) != 1{
		t.Fatalf(" should be 1")
	}

	if len(key[0].relations) != 3{
		t.Fatalf(" should be 3")
	}
	tbl.putByKey("2",true,"_1")

	rels := tbl.findByRel("_1")

	if len(rels) != 2{
		t.Fatalf(" should be 2")
	}
	ln := len(rels[0].relations) + len(rels[1].relations)

	if ln !=4 {
		t.Fatalf(" should be 5")
	}




}

