package graph

import
(
	"testing"
)

func constructGraph1()*Graph {
	gr := NewGraph()
	gr.AddNode("first")
	gr.AddNode("two")
	gr.AddNode("bang")
	gr.AddEdge("first", "two")
	gr.AddDirectedEdge("first", "bang")
	return gr
}

func TestHasNode(t *testing.T) {
	gr := constructGraph1()
	nodename1 := "first"
	nodename2 := "two"
	if !gr.HasNode(nodename1) {
		t.Errorf("node %s is not found", nodename1)
	}

	if !gr.HasNode(nodename2) {
		t.Errorf("node %s is not found", nodename2)
	}
}

func TestHasEdge(t *testing.T) {
	gr := constructGraph1()
	nodename1 := "first"
	nodename2 := "two"
	if !gr.HasEdge(nodename1, nodename2) {
		t.Errorf("Edge between %s and %s is not found", nodename1, nodename2)
	}

	if !gr.HasEdge(nodename2, nodename1) {
		t.Errorf("Edge between %s and %s is not found", nodename2, nodename1)
	}
}

func TestHasDirectedEdge(t *testing.T) {
	gr := constructGraph1()
	nodename1 := "first"
	nodename2 := "bang"
	if !gr.HasEdge(nodename1, nodename2) {
		t.Errorf("Edge between %s and %s is not found", nodename1, nodename2)
	}

	if gr.HasEdge(nodename2, nodename1) {
		t.Errorf("Edge between %s and %s is found", nodename2, nodename1)
	}
}