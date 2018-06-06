package digraph

import (
	//"fmt"

	"math"
)

func GetDistancesFromVertex(graph DiGraph, source_id uint64) map[uint64]uint64 {
	distances := make(map[uint64]uint64)
	distances[source_id] = 0

	explored := make(map[uint64]bool)
	explored[source_id] = true

	all_ids := graph.GetVertexIds()
	//total_number_vertices := len(all_ids)

	unexplored := make(map[uint64]bool)
	for _, id := range all_ids {
		unexplored[id] = true
	}
	delete(unexplored, source_id)

	for len(unexplored) > 0 {
		chooseNextEdge(graph, distances, explored, unexplored)
	}

	//fmt.Printf("distances %v\n", distances)

	return distances
}

// ! naive for now, should be a heap
func chooseNextEdge(graph DiGraph, distances map[uint64]uint64, explored, unexplored map[uint64]bool) {
	var min uint64 = math.MaxUint64
	//var min_edge *Edge = nil
	var chosen_unexplored_id uint64

	// TODO check all edges, not all vertex combos; more efficient
	for explored_node_id, _ := range explored {
		for unexplored_node_id, _ := range unexplored {

			// check both "directions" of possible connectivity since we're using a digraph
			connected_to_edge, connected_to_ok := graph.Connected(explored_node_id, unexplored_node_id)
			connected_from_edge, connected_from_ok := graph.Connected(unexplored_node_id, explored_node_id)

			// TODO edge.weight is implementation coupling
			// if connected check the min length
			if connected_to_ok && (distances[explored_node_id]+connected_to_edge.weight) < min {
				min = distances[explored_node_id] + connected_to_edge.weight
				//min_edge = connected_to_edge
				chosen_unexplored_id = unexplored_node_id

				continue
			}

			if connected_from_ok && (distances[explored_node_id]+connected_from_edge.weight) < min {
				min = distances[explored_node_id] + connected_from_edge.weight
				//min_edge = connected_from_edge
				chosen_unexplored_id = unexplored_node_id
			}
		}
	}

	distances[chosen_unexplored_id] = min
	explored[chosen_unexplored_id] = true
	delete(unexplored, chosen_unexplored_id)

	//return min_edge
}
