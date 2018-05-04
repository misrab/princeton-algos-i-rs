package graph

import (
	"testing"
)

func TestNewGraph(t *testing.T) {
	graph := NewGraph()

	// insert nodes
	graph.insertNodeAdjacency(1, []uint64{2, 3, 4, 5})
}
