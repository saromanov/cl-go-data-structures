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

//NewGraph provides construction of the graph object
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

//HasNode return true if node is registred and false otherwise
func (node *Graph) HasNode(title string) bool {
	_, ok := node.nodes[title]
	if !ok {
		return false
	}

	return true
}

//AddEdge provides connect tow nodes
func (node *Graph) AddEdge(innode, outnode string) error {
	node1, ok := node.nodes[innode]
	if !ok {
		return fmt.Errorf("Node %s is not found", innode)
	}

	node2, ok2 := node.nodes[outnode]
	if !ok2 {
		return fmt.Errorf("Node %s is not found", outnode)
	}

	node.nodes[innode].AddOutnode(node2)
	node.nodes[outnode].AddOutnode(node1)
	return nil
}

//HasEdge return true if innode and outnode is connected
func (node *Graph) HasEdge(innode, outnode string) bool {
	node1, ok := node.nodes[innode]
	if !ok {
		return false
	}
	return node1.HasEdge(outnode)
}
