package digraph

import (
	"testing"

	//"fmt"

	"github.com/stretchr/testify/assert"
)

/*func TestSCC(t *testing.T) {
	graph := createMediumGraph2()
	sccs := FindSCCs(graph)

}*/

/* // can't expect a cyclic graph to have a topological ordering!
func TestTopolicallyOrderMedium(t *testing.T) {
	graph := createMediumGraph()

	println("doing medium")
	labels, _ := TopologicallyOrder(graph)

	// make sure scc labels are ordered
	assert.True(t, labels[1] < 4 && labels[2] < 4 && labels[3] < 4)
	assert.True(t, labels[4] < 7 && labels[5] < 7 && labels[6] < 7)

}
*/

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

// create (1,2,3) -> (4,5,6) -> (7,8,9)
func createMediumGraph() DiGraph {
	graph := NewDiGraph()

	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 1)

	graph.AddEdge(4, 5)
	graph.AddEdge(5, 6)
	graph.AddEdge(6, 4)

	graph.AddEdge(7, 8)
	graph.AddEdge(8, 9)
	graph.AddEdge(9, 7)

	graph.AddEdge(3, 4)
	graph.AddEdge(6, 7)

	return graph
}

// create (7..,1,4) <- (6..,..9,3) <- (5,2,..8)
func createMediumGraph2() DiGraph {
	graph := NewDiGraph()

	graph.AddEdge(5, 2)
	graph.AddEdge(2, 8)
	graph.AddEdge(8, 5)

	graph.AddEdge(6, 9)
	graph.AddEdge(9, 3)
	graph.AddEdge(3, 6)

	graph.AddEdge(7, 1)
	graph.AddEdge(1, 4)
	graph.AddEdge(4, 7)

	graph.AddEdge(8, 6)
	graph.AddEdge(9, 7)

	return graph
}
