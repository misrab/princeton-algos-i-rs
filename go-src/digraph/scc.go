package digraph

import (
	//"fmt"
	"sort"

	"github.com/misrab/go-utils"
)

// ! yes, this could be made more generic
func DFSSCC(graph DiGraph, vertex *Vertex, explored map[uint64]bool, labels map[uint64]uint64, leaders map[uint64]uint64, leader_id uint64, finishing_time *uint64) {

	explored[vertex.id] = true
	leaders[vertex.id] = leader_id

	for _, outgoing_edge := range vertex.outgoing {
		outgoing_vertex := outgoing_edge.to
		if explored[outgoing_vertex.id] {
			continue
		}

		DFSSCC(graph, outgoing_vertex, explored, labels, leaders, leader_id, finishing_time)
	}

	*finishing_time++
	labels[vertex.id] = *finishing_time

}
func FindSCCs(graph DiGraph) map[uint64][]*Vertex {
	//fmt.Println("find sccs")

	reversed_graph := graph.Reverse()
	//fmt.Printf("original: %v\nreverse: %v\n", graph, reversed_graph)

	var finishing_time uint64 = 0
	var leader *Vertex

	explored := make(map[uint64]bool)
	leaders := make(map[uint64]uint64)
	labels := make(map[uint64]uint64)

	ids := reversed_graph.GetVertexIds()
	num_vertices := len(ids)
	sort.Slice(ids, func(i int, j int) bool { return ids[i] < ids[j] })

	// run DFS-Loop on reverse graph
	for i := num_vertices - 1; i > 0; i-- {
		if explored[ids[i]] {
			continue
		}
		leader, _ = reversed_graph.GetVertex(ids[i])
		// care about labels, not leaders
		DFSSCC(reversed_graph, leader, explored, labels, leaders, ids[i], &finishing_time)
	}

	//fmt.Printf("finishint times:\n%v", labels)

	// run DFS-Loop on original graph
	// using reverse finishing times
	explored = make(map[uint64]bool) // reinitialise
	finishing_time_to_id := utils.FlipMapUint64(labels)
	for i := len(labels); i > 0; i-- {
		node_id := finishing_time_to_id[uint64(i)]
		if explored[node_id] {
			continue
		}
		// initiate dfs from that actual node
		leader, _ := graph.GetVertex(node_id)
		// care about leaders, not labels
		DFSSCC(graph, leader, explored, labels, leaders, node_id, &finishing_time)
	}

	// nodes with same leaders are the scc's
	sccs := make(map[uint64][]*Vertex)
	for node_id, leader := range leaders {
		if sccs[leader] == nil {
			sccs[leader] = make([]*Vertex, 0)
		}

		node, _ := graph.GetVertex(node_id)
		sccs[leader] = append(sccs[leader], node)
	}

	return sccs
}

func TopologicallyOrder(graph DiGraph) (old_to_new map[uint64]uint64, new_to_old map[uint64]uint64) {

	explored := make(map[uint64]bool)

	// pick a random start vertex - just first for now
	ids := graph.GetVertexIds()
	n := len(ids)
	if n == 0 {
		return nil, nil
	}

	old_to_new = make(map[uint64]uint64)
	current_label := uint64(n)

	for _, id := range ids {
		if explored[id] {
			continue
		}

		vertex, _ := graph.GetVertex(id)
		depthFirstSearch(graph, vertex, explored, old_to_new, &current_label)
	}

	new_to_old = utils.FlipMapUint64(old_to_new)

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
	*current_label--
}
