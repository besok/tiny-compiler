package graph

import (
	"log"
	"testing"
)

func Test_commonGraph(t *testing.T) {
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

func Test_MarkQueue(t *testing.T) {
	v1 := NewVertex(1, StringValue("Test1"))
	v2 := NewVertex(2, StringValue("Test2"))
	v3 := NewVertex(3, StringValue("Test3"))

	queue := NewQueue()
	queue.push(v1)
	queue.push(v2)
	queue.push(v3)

	if queue.len() < 3 {
		t.Fatalf("error - needs at least 3")
	}

	_ = queue.pop()
	if queue.len() > 2 {
		t.Fatalf("error - needs at least 3")
	}

}

func Test_PrintGraph(t *testing.T) {
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

	PrintGraph(graph, BFS)

}
func Test_StackAndQueue(t *testing.T) {
	q := NewQueue()
	q.push(NewVertex(1, StringValue("Test1")))
	q.push(NewVertex(2, StringValue("Test2")))

	v1 := q.pop()

	if q.len() > 1 {
		t.Fatalf("fatal")
	}

	if v1.Id != 1 {
		t.Fatalf("fatal")
	}

	s := NewStack()
	s.push(NewVertex(1, StringValue("Test1")))
	s.push(NewVertex(2, StringValue("Test2")))

	v1 = q.pop()

	if q.len() > 1 {
		t.Fatalf("fatal")
	}

	if v1.Id != 2 {
		t.Fatalf("fatal")
	}
}
func Test_BfsDfs(t *testing.T) {
	v1 := NewVertex(1, StringValue("Test1"))
	v2 := NewVertex(2, StringValue("Test2"))
	v3 := NewVertex(3, StringValue("Test3"))
	v4 := NewVertex(4, StringValue("Test4"))
	v5 := NewVertex(5, StringValue("Test5"))

	_ = Relation(v1, v2)
	_ = Relation(v1, v3)
	_ = Relation(v1, v4)
	_ = Relation(v2, v3)
	_ = Relation(v3, v4)
	_ = Relation(v3, v5)

	graph := NewGraph(v2)

	PrintGraph(graph, BFS)
	log.Println(" ---- ")
	PrintGraph(graph, DFS)

}

func Test_PriorityQueue(t *testing.T) {
	v1 := NewVertex(1, StringValue("Test1"))
	v2 := NewVertex(2, StringValue("Test2"))
	v3 := NewVertex(3, StringValue("Test3"))
	v4 := NewVertex(4, StringValue("Test4"))
	v5 := NewVertex(5, StringValue("Test5"))

	minPrQ := NewPriorityQueue(Lowest)
	minPrQ.pushWithPriority(v4, 4)
	minPrQ.pushWithPriority(v3, 3)
	minPrQ.pushWithPriority(v5, 5)
	minPrQ.pushWithPriority(v1, 1)
	minPrQ.pushWithPriority(v2, 2)

	for i := 1; i < 6; i++ {
		v := minPrQ.pop()
		if v.Id != i {
			t.Fatalf(" error : it should be %d but it is %d", i, v.Id)
		}

	}

	maxPrQ := NewPriorityQueue(Highest)

	maxPrQ.pushWithPriority(v4, 4)
	maxPrQ.pushWithPriority(v3, 3)
	maxPrQ.pushWithPriority(v5, 5)
	maxPrQ.pushWithPriority(v1, 1)
	maxPrQ.pushWithPriority(v2, 2)

	for i := 5; i > 0; i-- {
		v := maxPrQ.pop()
		if v.Id != i {
			t.Fatalf(" error : it should be %d but it is %d", i, v.Id)
		}
	}
	maxPrQ.push(v4)
	maxPrQ.push(v3)
	maxPrQ.push(v2)
	maxPrQ.push(v1)

	v := maxPrQ.pop()

	if v.Id != v4.Id {
		t.Fatalf("error:%d", v.Id)
	}

	q := NewPriorityQueue(Highest)

	q.push(v1)
	q.push(v1)

	ln := q.len()

	if ln > 1 {
		t.Fatalf("it needs to have idenpotent")
	}

}
func Test_GraphShortestPath(t *testing.T) {

	v1 := NewVertex(1, StringValue("Test1"))
	v2 := NewVertex(2, StringValue("Test2"))
	v3 := NewVertex(3, StringValue("Test3"))
	v4 := NewVertex(4, StringValue("Test4"))
	v5 := NewVertex(5, StringValue("Test5"))

	graph := NewGraph(v1)
	graph.Relation(v1, v2)
	graph.Relation(v2, v3)
	graph.Relation(v3, v4)
	graph.Relation(v4, v5)

	path := graph.ShortestPath(1, 5)
	log.Println(path)
	for i := 0; i < 3; i++ {
		if path[i] != i + 2{
			t.Fatalf("error ")
		}
	}
}
func Test_GraphShortestPath2(t *testing.T) {

	v1 := NewVertex(1, StringValue("Test1"))
	v2 := NewVertex(2, StringValue("Test2"))
	v3 := NewVertex(3, StringValue("Test3"))
	v4 := NewVertex(4, StringValue("Test4"))
	v5 := NewVertex(5, StringValue("Test5"))

	graph := NewGraph(v1)
	graph.Relation(v1, v2)
	graph.Relation(v1, v3)
	graph.Relation(v1, v4)
	graph.Relation(v1, v5)

	path := graph.ShortestPath(1, 5)
	log.Println(path)

	if path[0] != 1 || path[1] !=5 {
		t.Fatalf("error")
	}
}
func Test_GraphShortestPath3(t *testing.T) {

	v1 := NewVertex(1, StringValue("Test1"))
	v2 := NewVertex(2, StringValue("Test2"))
	v3 := NewVertex(3, StringValue("Test3"))
	v4 := NewVertex(4, StringValue("Test4"))
	v5 := NewVertex(5, StringValue("Test5"))

	graph := NewGraph(v1)
	graph.Relation(v1, v2)
	graph.Relation(v1, v3)
	graph.Relation(v2, v4)
	graph.Relation(v3, v4)
	graph.Relation(v4, v5)

	path := graph.ShortestPath(1, 5)
	if len(path) != 4 {
		t.Fatalf("error")
	}

}
func Test_GraphRelation(t *testing.T) {
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

	ln := graph.Len()
	if ln != 3 {
		t.Fatalf("bad len:%d", ln)
	}
	PrintGraph(graph, BFS)

}

type StringValue string

func (v StringValue) Get() interface{} {
	return string(v)
}
