package digraph

import (
	//"fmt"

	"github.com/misrab/go-utils"
)

func TopologicallyOrder(graph DiGraph) (old_to_new map[uint64]uint64, new_to_old map[uint64]uint64) {

	explored := make(map[uint64]bool)

	// pick a random start vertex - just first for now
	ids := graph.GetVertexIds()
	n := len(ids)
	if n == 0 {
		return nil
	}

	old_to_new := make(map[uint64]uint64)
	current_label := uint64(n)
	for _, id := range ids {
		if explored[id] {
			continue
		}

		vertex, _ := graph.GetVertex(id)
		depthFirstSearch(graph, vertex, explored, old_to_new, &current_label)
	}

	new_to_old := utils.FlipMap(old_to_new)

	return old_to_new, new_to_old
}

func depthFirstSearch(graph DiGraph, vertex *Vertex, explored map[uint64]bool, labels map[uint64]uint64, current_label *uint64) {

	explored[vertex.id] = true

	for _, outgoing_edge := range vertex.outgoing {
		outgoing_vertex := outgoing_edge.to
		if explored[outgoing_vertex.id] {
			continue
		}

		depthFirstSearch(graph, outgoing_vertex, explored, labels, current_label)
	}

	labels[vertex.id] = *current_label
	*current_label -= 1
}
