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
	tree.node = createNode(0)
	return tree
}

//Insert provides append new item
func (bt *BinaryTree) Insert(item int) {
	bt.insert(bt.node.root, item)
}

//Remove provides removing item from tree
func(bt *BinaryTree) Remove(item int) {
	bt.remove(bt.node.root, item)
}

func (bt *BinaryTree) Find(item int) bool {
	return bt.find(bt.node, item)
}

//helful method for the insert
func (bt *BinaryTree) insert(node *Node, item int) {
	if bt.node == nil {
		node := createNode(item)
		bt.node = node
	} else {
		if item < bt.node.item {
			bt.node.left = createNode(item)
		} else if item > bt.node.item {
			bt.node.right = createNode(item)
		}
	}
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

//helful method for finding element
func (bt *BinaryTree) find(node *Node, item int) bool {
	if node == nil {
		return false
	} else {
		if item < bt.node.item {
			return bt.find(node.left, item)
		} else if item > bt.node.item {
			return bt.find(node.right, item)
		} else {
			return true
		}
	}
}

//helpful method for create node
func createNode(item int)*Node {
	return &Node{nil, nil, nil, item}
}