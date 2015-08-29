package btree

import
(
	//"fmt"
)

type Node struct {
	left, right, root *Node
	item int
}

type BinaryTree struct {
	node *Node
	height uint
}

func New()*BinaryTree {
	tree := new(BinaryTree)
	return tree
}

func (bt *BinaryTree) Insert(item int) {
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

func (bt *BinaryTree) createNode(item int)*Node {
	return &Node{nil, nil, nil, item}
}