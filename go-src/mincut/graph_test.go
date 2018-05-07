package graph

import (
	"log"
	"testing"
)

func TestNewGraph(t *testing.T) {
	graph := NewGraph()

	// insert nodes
	graph.insertNodeAdjacency(1, []uint64{2, 3, 4, 5})

	node1, _ := graph.GetNode(1)
	log.Printf("%v\n", node1)

	node2, _ := graph.GetNode(2)
	log.Printf("%v\n", node2)

}
