package digraph

import (
	//"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDijkstra(t *testing.T) {
	graph := NewDiGraph()

	graph.AddEdge(1, 2, 1)
	graph.AddEdge(1, 3, 5)
	graph.AddEdge(2, 4, 7)
	graph.AddEdge(3, 4, 1)

	lenghths_from_one := GetDistancesFromVertex(graph, 1)

	assert.True(t, lenghths_from_one[4] == 6)
}

func TestDijkstraMedium(t *testing.T) {
	graph := NewDiGraph()

	graph.AddEdge(1, 2, 1)
	graph.AddEdge(1, 3, 4)
	graph.AddEdge(2, 4, 6)
	graph.AddEdge(3, 4, 3)

	graph.AddEdge(2, 3, 2)

	lenghths_from_one := GetDistancesFromVertex(graph, 1)

	//fmt.Printf("%v", lenghths_from_one)

	assert.True(t, lenghths_from_one[4] == 6)
}
