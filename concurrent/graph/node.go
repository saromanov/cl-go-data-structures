package graph


//Node is basic implementation of nodes for graphs
type Node struct {
	Title  string
	Value  interface{}
	Outnodes []*Node
	Innodes  []*Node
}

//NewNode provides construction of Node object
func NewNode(title string)*Node {
	node := new(Node)
	node.Title = title
	node.Outnodes = []*Node{}
	node.Innodes = []*Node{}
	return node
}

func (n *Node) AddOutnode(node *Node) {
	n.Outnodes = append(n.Outnodes, node)
}

func (n *Node) AddInnode(node *Node) {
	n.Innodes = append(n.Innodes, node)
}

func (n *Node) OutnodesLen() int {
	return len(n.Outnodes)
}

func (n *Node) InnodesLen() int {
	return len(n.Innodes)
}

//Find provides finding nodes from Outnodes
func (n *Node) Find(title string)*Node {
	for _, n := range n.Outnodes {
		if n.Title == title {
			return n
		}
	}

	return nil
}

//HasEdge provides finding node in outnodes
func (n *Node) HasEdge(title string) bool {
	result := n.Find(title)
	if result != nil {
		return true
	}

	return false
}

//Copy nodes returns copy of current node
func (node *Node) Copy()*Node {
	newnode := NewNode(node.Title)
	newnode = node
	return newnode
}
