package graph

import (
	"fmt"
)

type GraphFuncs interface {
	//Store new node
	AddNode(title string) error

	//Connect two nodes
	AddEdge(innode, outnode string) error

	//GetNode(title string) *Node
}

type Graph struct {
	nodes map[string]*Node
}

func NewGraph() *Graph {
	graph := new(Graph)
	graph.nodes = map[string]*Node{}
	return graph
}

//AddNode provoides creating of new node object and storing by title
func (node *Graph) AddNode(title string) error {
	node.nodes[title] = NewNode(title)
	return nil
}

func (node *Graph) AddEdge(innode, outnode string) error {
	node1, ok := node.nodes[innode]
	if !ok {
		return fmt.Errorf("Node %s is not found", innode)
	}

	node2, ok2 := node.nodes[outnode]
	if !ok2 {
		return fmt.Errorf("Node %s is not found", outnode)
	}

	node.nodes[innode].AddOutnode(node1)
	node.nodes[outnode].AddInnode(node2)
	return nil
}
