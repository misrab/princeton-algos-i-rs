package digraph

import (
	"testing"

	//"fmt"

	"github.com/stretchr/testify/assert"
)

func TestTopolicallyOrder(t *testing.T) {
	graph := NewDiGraph()

	// the nodes are labelled in reverse order
	// 1 <- 2 <- 4
	//   <- 3 <- 4
	// we will want (1,2,3,4) -> (4,3,2,1) or (4,2,3,1)
	graph.AddEdge(2, 1)
	graph.AddEdge(3, 1)
	graph.AddEdge(4, 2)
	graph.AddEdge(4, 3)

	labels, _ := TopologicallyOrder(graph)
	//fmt.Printf("labels:\n%v", labels)

	assert.True(t, labels[4] == 1)
	assert.True(t, labels[1] == 4)
	assert.True(t, labels[3] == 3 || labels[3] == 2)
	assert.True(t, labels[2] == 2 || labels[2] == 3)

}
