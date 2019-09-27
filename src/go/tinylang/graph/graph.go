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
	_, bL := g.FindById(left.Id, DFS)
	_, bR := g.FindById(right.Id, BFS)

	if !(bL || bR) {
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

	Task(g, t, compare)

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

func DefaultTask(g *Graph, f func(v *Vertex) bool) {
	Task(g, BFS, f)
}
func Task(g *Graph, t TypeOfSearch, f func(v *Vertex) bool) {
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
	Task(g, t, printVertex)
}

// graph len. Takes (O) to get a value so avoid invoking this function twice
func (g *Graph) Len() int {
	ln := 0
	DefaultTask(g, func(v *Vertex) bool { ln++; return true })
	return ln
}
func (g *Graph) ShortestPath(vStart int, vFinish int) []int {

	var fin *Vertex
	var str *Vertex
	dist := map[int]int{}
	prev := map[int]int{}
	queue := NewPriorityQueue(Lowest)

	ln := g.Len()

	g.TraverseBfs(func(v *Vertex) bool {
		id := v.Id
		val := ln
		if id == vStart {
			val = 0
			str = v
		}
		if id == vFinish {
			fin = v
		}

		dist[id] = val
		prev[id] = -1
		queue.pushWithPriority(v, val)
		return true
	})

	if str == nil || fin == nil{
		return []int{}
	}

	for ; queue.len() > 0; {
		vertex := queue.pop()
		dst := dist[vertex.Id]

		for i := 0; i < len(vertex.Neighbours); i++ {
			n := vertex.Neighbours[i]
			alt := dst + 1
			nId := n.Id
			if alt < dist[nId] {
				dist[nId] = alt
				prev[nId] = vertex.Id
				queue.pushWithPriority(n, alt)
			}
		}

		if vertex.Id == vFinish {
			break
		}
	}

	res := make([]int, 0)

	v := prev[fin.Id]

	for ; v != -1; {
		res = append(res, v)
		next, _ := g.FindById(v, BFS)
		if v == str.Id {
			break
		}
		v = prev[next.Id]

	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	if len(res) > 0 {
		res = append(res, fin.Id)
	}
	return res
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

type PrQueue struct {
	order []*Vertex
	elems map[*Vertex]int
	cmp   func(left int, right int) bool
}

func (p *PrQueue) pushWithPriority(vm *Vertex, priority int) {
	p.push(vm)
	p.elems[vm] = priority

}
func (p *PrQueue) push(vm *Vertex) {
	isNew := true
	for _, v := range p.order {
		if v == vm {
			isNew = false
			break
		}
	}
	if isNew {
		p.order = append(p.order, vm)
		p.elems[vm] = 0
	}
}
func (p *PrQueue) len() int {
	return len(p.order)
}
func (p *PrQueue) pickPriority(v *Vertex) int {
	pr, ok := p.elems[v]
	if !ok {
		return -1
	}
	return pr
}
func (p *PrQueue) popWithPriority() (*Vertex, int) {
	var curr *Vertex
	var currPr int
	idx := -1
	for i, v := range p.order {
		pr := p.elems[v]
		if curr == nil {
			curr = v
			currPr = pr
			idx = i
		} else {
			if p.cmp(currPr, pr) {
				curr = v
				currPr = pr
				idx = i
			}
		}
	}
	if idx >= 0 {
		delete(p.elems, curr)
		p.order = append(p.order[:idx], p.order[idx+1:]...)
	}
	return curr, currPr
}
func (p *PrQueue) pop() *Vertex {
	vertex, _ := p.popWithPriority()
	return vertex
}

func Highest(srs int, inc int) bool {
	if srs < inc {
		return true
	}
	return false
}
func Lowest(srs int, inc int) bool {
	if srs >= inc {
		return true
	}
	return false
}

func NewPriorityQueue(cmp func(srs int, inc int) bool) *PrQueue {
	return &PrQueue{elems: map[*Vertex]int{}, cmp: cmp, order: []*Vertex{}}
}
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
