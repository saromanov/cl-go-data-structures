package aatree

import
(
	"sync"
)

//https://en.wikipedia.org/wiki/AA_tree


type Node struct {
	root, left, right *Node
	item int
}

type AATree struct {
	node *Node
	height uint
}

func New()*AATree{
	node := new(AATree)
}

//Insert provides append new item
func (aat *AATree) Insert(item int) {

	if bt.node == nil {
		node := bt.createNode(item)
		bt.node = node
	} else {
		if item < bt.node.item {
			bt.node.left = bt.createNode(item)
		} else {

		}
	}
}

func (aat *AATree) skew(node *Node)*Node {
	if node == nil {
		return nil
	} else if node.left == nil {
		return node
	}
}