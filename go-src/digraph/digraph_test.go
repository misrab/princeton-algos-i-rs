package digraph

import (
	//"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	d := NewDiGraph()

	d.AddEdge(1, 2, 1)
	d.AddEdge(1, 3, 1)
	d.AddEdge(2, 4, 1)
	d.AddEdge(3, 4, 1)

	new_graph := d.Copy()

	// assert vertices are there
	_, one_found := new_graph.GetVertex(1)
	_, two_found := new_graph.GetVertex(2)
	_, three_found := new_graph.GetVertex(3)
	_, four_found := new_graph.GetVertex(4)

	assert.True(t, one_found && two_found && three_found && four_found)

	// randomly check connections of two nodes
	// ! could be made exhaustive
	new_graph.Connected(1, 3)
	new_graph.Connected(3, 4)
	new_graph.Connected(4, 3)
}

func TestBasic(t *testing.T) {
	d := NewDiGraph()

	edge := d.AddEdge(1, 2, 1)

	from, from_found := d.GetVertex(1)
	to, to_found := d.GetVertex(2)

	assert.True(t, from_found)
	assert.True(t, to_found)

	assert.Empty(t, from.incoming)
	assert.Empty(t, to.outgoing)

	assert.Len(t, from.outgoing, 1)
	assert.Len(t, to.incoming, 1)

	assert.Equal(t, from.outgoing[0], edge)
	assert.Equal(t, to.incoming[0], edge)

	d.RemoveEdge(edge)

	assert.Empty(t, d.edges)

	assert.Empty(t, from.incoming)
	assert.Empty(t, from.outgoing)
	assert.Empty(t, to.incoming)
	assert.Empty(t, to.outgoing)

	//fmt.Printf("%v", d)
}
