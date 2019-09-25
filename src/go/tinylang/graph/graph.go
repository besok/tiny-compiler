package graph

import (
	"fmt"
	"log"
	"strings"
)

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
	var returnVertex *Vertex

	bfsSearch := func(v *Vertex) bool {
		if v.Id == id {
			returnVertex = v
			return false
		}
		return true
	}

	g.TraverseBfs(bfsSearch)

	if returnVertex != nil{
		return returnVertex, true
	}

	return nil, false
}

func PrintGraph(g *Graph) {

	printVertex := func(v *Vertex) bool {
		ids := make([]string, len(v.Neighbours))

		for i, el := range v.Neighbours {
			ids[i] = fmt.Sprintf("%d", el.Id)
		}

		idsStr := strings.Join(ids, ",")
		log.Printf(" vertex: id:%d, value:%#v, neighbours:[%s]", v.Id, v.Value.Get(), idsStr)
		return true
	}

	g.TraverseBfs(printVertex)
}

// Traverse vertexes for this graph.
// Action is a function taking every vertex and returning a flag indicating about further step.
// If action return is false the traversal process is broken.
func (g *Graph) TraverseBfs(action func(v *Vertex) bool) {
	history := make(map[int]bool, 0)
	q := NewQueue()

	q.Push(g.root)

	for ; q.Len() > 0; {
		v := q.Pop()
		_, ok := history[v.Id]
		if ok {
			continue
		}
		history[v.Id] = true

		if !action(v) {
			return
		}
		for _, n := range v.Neighbours {
			q.Push(n)
		}
	}
}

func (q *Queue) Push(vm *Vertex) {
	*q = append(*q, *vm)
}
func (q *Queue) Len() int {
	return len(*q)
}
func (q *Queue) Pop() *Vertex {
	ln := len(*q)
	last := (*q)[ln-1]

	*q = (*q)[:ln-1]

	return &last
}

type Queue []Vertex

func NewQueue() *Queue {
	queues := make(Queue, 0)
	return &queues
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
