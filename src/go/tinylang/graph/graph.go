package graph

import "log"

type Graph struct {
	root *Vertex
}

type Vertex struct {
	Id         int
	Value      Value
	Neighbours []*Vertex
}

type Value interface {
	Get() interface{}
}

func NewGraph(root *Vertex) *Graph {
	return &Graph{root: root}
}

func NewVertex(id int, value Value) *Vertex {
	return &Vertex{Neighbours: make([]*Vertex, 0), Id: id, Value: value}
}

func Relation(left *Vertex, right *Vertex) error {
	err := left.add(right)
	err = right.add(left)

	log.Printf("relation between: [%d] - [%d]", left.Id, right.Id)

	return err
}

func (g *Graph) Relation(left *Vertex, right *Vertex) bool {
	left, bL := g.FindById(left.Id)
	right, bR := g.FindById(right.Id)

	if !(bL && bR) {
		return false
	}

	err := Relation(left, right)
	if err != nil {
		log.Println("error while adding to gragh: ", err)
		return false
	}
	return true
}

func (g *Graph) FindById(id int) (*Vertex, bool) {
	vertex := g.root

	if vertex.Id == id{
		return vertex,true
	}
	queue := NewQueue()

	for


	return nil, false
}

func (q *Queue) PushV(vm *Vertex, mark bool) {
	vmark := VertexMark{Marked: mark, Vertex: *vm}
	q.Push(&vmark)
}

func (q *Queue) Push(vm *VertexMark) {
	*q = append(*q, *vm)
}
func (q *Queue) Len() int {
	return len(*q)
}
func (q *Queue) Pop() *VertexMark {
	ln := len(*q)
	last := (*q)[ln-1]

	*q = (*q)[:ln-1]

	return &last
}


type Queue []VertexMark


func NewQueue() *Queue {
	queues := make(Queue, 0)
	return &queues
}

type VertexMark struct {
	Vertex Vertex
	Marked bool
}

func (v *Vertex) checkId(id int) bool {
	for _, el := range v.Neighbours {
		if el.Id == id {
			return true
		}
	}
	return false
}
func (v *Vertex) add(n *Vertex) error {
	if v.checkId(n.Id) {
		return nil
	}

	v.Neighbours = append(v.Neighbours, n)
	return nil
}

type Error struct {
	msg string
}

func (e Error) Error() string {
	return e.msg
}
