package graph

import (
	"testing"
)

func Test_commonGraph(t *testing.T){
	v1 := NewVertex(1, StringValue("Test1"))
	v2 := NewVertex(2, StringValue("Test2"))
	v3 := NewVertex(3, StringValue("Test3"))
	v4 := NewVertex(4, StringValue("Test4"))
	v5 := NewVertex(5, StringValue("Test5"))

	err := Relation(v1, v2)
	err = Relation(v2, v3)
	err = Relation(v3, v4)
	err = Relation(v4, v5)
	if err != nil {
		t.Fatalf("bad dec")
	}


}

func Test_MarkQueue(t *testing.T){
	v1 := NewVertex(1, StringValue("Test1"))
	v2 := NewVertex(2, StringValue("Test2"))
	v3 := NewVertex(3, StringValue("Test3"))

	queue := NewQueue()
	queue.push(v1)
	queue.push(v2)
	queue.push(v3)

	if queue.len() < 3{
		t.Fatalf("error - needs at least 3")
	}

	_ = queue.pop()
	if queue.len() > 2{
		t.Fatalf("error - needs at least 3")
	}

}

func Test_PrintGraph(t *testing.T){
	v1 := NewVertex(1, StringValue("Test1"))
	v2 := NewVertex(2, StringValue("Test2"))
	v3 := NewVertex(3, StringValue("Test3"))
	v4 := NewVertex(4, StringValue("Test4"))
	v5 := NewVertex(5, StringValue("Test5"))

	err := Relation(v1, v2)
	err = Relation(v2, v3)
	err = Relation(v3, v4)
	err = Relation(v4, v5)

	if err != nil {
		t.Fatalf("bad dec")
	}

	graph := NewGraph(v1)

	PrintGraph(graph)

}
func Test_GraphRelation(t *testing.T){
	v1 := NewVertex(1, StringValue("Test1"))
	v2 := NewVertex(2, StringValue("Test2"))
	v3 := NewVertex(3, StringValue("Test3"))
	v4 := NewVertex(4, StringValue("Test4"))
	v5 := NewVertex(5, StringValue("Test5"))

	err := Relation(v1, v2)
	err = Relation(v2, v3)



	if err != nil {
		t.Fatalf("bad dec")
	}

	graph := NewGraph(v1)
	relationFalse := graph.Relation(v4, v5)
	relationTrue := graph.Relation(v1, v3)

	if relationFalse {
		t.Fatalf("bad dec")
	}
	if !relationTrue {
		t.Fatalf("bad dec")
	}


}

type StringValue string

func (v StringValue) Get() interface{} {
	return string(v)
}
