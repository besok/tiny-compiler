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

	log.Printf("relation between: [%d] - [%d]",left.Id,right.Id)

	return err
}



func (v *Vertex) add(n *Vertex) error {
	v.Neighbours = append(v.Neighbours, n)
	return nil
}


type Error struct {
	msg string
}


func (e Error) Error() string {
	return e.msg
}
