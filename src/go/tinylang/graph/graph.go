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
	left, bL := g.FindById(left.Id, DFS)
	right, bR := g.FindById(right.Id, BFS)

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

func (g *Graph) FindById(id int, t TypeOfSearch) (*Vertex, bool) {
	var returnVertex *Vertex

	compare := func(v *Vertex) bool {
		if v.Id == id {
			returnVertex = v
			return false
		}
		return true
	}

	CommonTask(g, t, compare)

	if returnVertex != nil {
		return returnVertex, true
	}

	return nil, false
}

type TypeOfSearch string

var (
	BFS TypeOfSearch = "bfs"
	DFS TypeOfSearch = "dfs"
)

func (g *Graph) pickBy(t TypeOfSearch) func(func(v *Vertex) bool) {
	switch t {
	case BFS:
		return g.TraverseBfs
	case DFS:
		return g.TraverseDfs
	default:
		panic(" BFS or DFS for TypeOfSearch")
	}
}

func CommonTask(g *Graph, t TypeOfSearch, f func(v *Vertex) bool) {
	g.pickBy(t)(f)
}

func PrintGraph(g *Graph, t TypeOfSearch) {

	printVertex := func(v *Vertex) bool {
		ids := make([]string, len(v.Neighbours))

		for i, el := range v.Neighbours {
			ids[i] = fmt.Sprintf("%d", el.Id)
		}

		idsStr := strings.Join(ids, ",")
		log.Printf(" vertex: id:%d, value:%#v, neighbours:[%s]", v.Id, v.Value.Get(), idsStr)
		return true
	}
	CommonTask(g, t, printVertex)
}

func(g *Graph) ShortestPath(vStart *Vertex, vFinish *Vertex) int {

}

// Traverse vertexes for this graph.
// Action is a function taking every vertex and returning a flag indicating about further step.
// If action return is false the traversal process is broken.
func (g *Graph) TraverseBfs(action func(v *Vertex) bool) {
	history := make(map[int]bool, 0)
	q := NewQueue()

	root := g.root
	history[root.Id] = true
	q.push(root)

	for ; q.len() > 0; {
		v := q.pop()

		if !action(v) {
			return
		}

		for _, n := range v.Neighbours {
			id := n.Id
			_, ok := history[id]
			if !ok {
				history[id] = true
				q.push(n)
			}
		}
	}
}

// Traverse vertexes for this graph.
// Action is a function taking every vertex and returning a flag indicating about further step.
// If action return is false the traversal process is broken.
func (g *Graph) TraverseDfs(action func(v *Vertex) bool) {
	history := make(map[int]bool, 0)
	q := NewStack()

	q.push(g.root)

	for ; q.len() > 0; {
		v := q.pop()
		_, ok := history[v.Id]
		if ok {
			continue
		}
		history[v.Id] = true

		if !action(v) {
			return
		}
		for _, n := range v.Neighbours {
			q.push(n)
		}
	}
}

type History interface {
	push(vm *Vertex)
	len() int
	pop() *Vertex
}

func (q *Queue) push(vm *Vertex) {
	*q = append(*q, *vm)
}
func (q *Queue) len() int {
	return len(*q)
}
func (q *Queue) pop() *Vertex {
	first := (*q)[0]
	*q = (*q)[1:]

	return &first
}

func (q *Stack) push(vm *Vertex) {
	*q = append(*q, *vm)
}
func (q *Stack) len() int {
	return len(*q)
}
func (q *Stack) pop() *Vertex {
	ln := len(*q)
	last := (*q)[ln-1]
	*q = (*q)[:ln-1]
	return &last
}

type Queue []Vertex
type Stack []Vertex

func NewStack() *Stack {
	stack := make(Stack, 0)
	return &stack
}
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
