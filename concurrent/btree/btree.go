package btree

import
(
	//"sync"
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

//Insert provides append new item
func (bt *BinaryTree) Insert(item int) {
	bt.insert(bt.node.root, item)
}

func (bt *BinaryTree) insert(node *Node, item int) {
	if bt.node == nil {
		node := bt.createNode(item)
		bt.node = node
	} else {
		if item < bt.node.item {
			bt.node.left = bt.createNode(item)
		} else if item > bt.node.item {
			bt.node.right = bt.createNode(item)
		}
	}
}

//Remove provides removing item from tree
func(bt *BinaryTree) Remove(item int) {
	bt.remove(bt.node.root, item)
}

//helful method for removing data
func (bt *BinaryTree) remove(node *Node, item int)(c *Node) {
	if bt.node.root == nil {
		return
	}

	if bt.node.left == nil && bt.node.right == nil && bt.node.item == item {
		c = bt.node.root
		bt.node.root = nil
		return
	}
	return
}

//helpful method for create node
func (bt *BinaryTree) createNode(item int)*Node {
	return &Node{nil, nil, nil, item}
}