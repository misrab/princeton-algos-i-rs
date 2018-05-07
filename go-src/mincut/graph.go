package graph

import (
	"fmt"
	"time"

	"math/rand"
)

type Vertex struct {
	id          uint64
	connections []*Edge
}

func (v Vertex) String() string {
	return fmt.Sprintf("{\n\tid: %d\n\tconnections: %v\n}", v.id, v.connections)
}

type Edge struct {
	from *Vertex
	to   *Vertex
}

func (e *Edge) String() string {
	return fmt.Sprintf("(%v,%v)", e.from.id, e.to.id)
}

func (g *graph) String() string {
	result := ""

	//result += fmt.Sprintf("vertices: %v\n", g.vertices)
	//result += fmt.Sprintf("edges: %v\n", g.edges)

	for _, node := range g.GetNodes() {
		result += fmt.Sprintf("%v\n", node)
	}

	return result
}

type graph struct {
	vertices map[uint64]*Vertex
	edges    []*Edge
}

type Graph interface {
	String() string
	NumNodes() uint64
	NumEdges() uint64

	GetNodes() map[uint64]*Vertex
	GetEdges() []*Edge

	//insertNode(id uint64)

	insertNodeAdjacency(id uint64, connections []uint64)

	//GetNodeAdjacency(id uint64) []*Edge

	GetNode(uint64) (*Vertex, bool)

	ContractEdge(e *Edge)

	RemoveNode(id uint64)

	ContractionAlgorithm() uint64
}

func NewGraph() Graph {
	graph := new(graph)

	// allocate memory
	graph.vertices = make(map[uint64]*Vertex)
	graph.edges = make([]*Edge, 0)

	return graph
}

func (g *graph) NumNodes() uint64 {
	return uint64(len(g.vertices))
}
func (g *graph) NumEdges() uint64 {
	return uint64(len(g.edges))
}

func (g *graph) ContractionAlgorithm() uint64 {
	rand.Seed(time.Now().Unix())

	fmt.Printf("graph starting as %v\n", g)

	for g.NumNodes() > 2 {
		num_edges := len(g.edges)
		println(num_edges)

		// choose an edge at random
		edge := g.edges[rand.Intn(num_edges)]
		fmt.Printf("contracting %v\n", edge)
		g.ContractEdge(edge)
		fmt.Printf("graph is now: %v\n", g)
	}

	return uint64(len(g.edges))
}

// remove all edges pointing to a node, then remove the node from the graph
func (g *graph) RemoveNode(id uint64) {
	//node_indices := make(map[int]bool)
	edge_indices := make(map[int]bool)

	for i, e := range g.edges {
		if e.to.id == id || e.from.id == id {
			edge_indices[i] = true
		}
	}

	//for i, n := range g.vertices {
	//	if n.id == id {
	//		node_indices[i] = true
	//	}
	//}

	delete(g.vertices, id)

	//g.vertices = removeFromNodes(g.vertices, node_indices)
	g.edges = removeFromEdges(g.edges, edge_indices)
}

func (g *graph) GetNodes() map[uint64]*Vertex {
	return g.vertices
}

func (g *graph) GetEdges() []*Edge {
	return g.edges
}

func (g *graph) ContractEdge(e *Edge) {
	from := e.from
	to := e.to

	//self_loop_indices := make(map[int]bool)

	// condense to into from
	// to's edges will now point to from
	for _, to_edge := range to.connections {
		if to_edge.from.id == to.id {
			to_edge.from = from // point to new location
		}
		if to_edge.to.id == to.id {
			to_edge.to = from
		}

		// remove self loop if any
		//if to_edge.from.id == to_edge.to.id {
		//	self_loop_indices[i] = true
		//	fmt.Printf("self loop %v\n", to_edge)
		//}
	}

	// now combine the edges into "from", after having removed self loops
	//new_connections := removeFromEdges(to.connections, self_loop_indices)
	from.connections = append(from.connections, to.connections...)

	// remove self loops on "from"
	from.connections = removeSelfLoops(from.connections)
	g.edges = removeSelfLoops(g.edges)

	// delete to in the graph
	g.RemoveNode(to.id)
}

func removeSelfLoops(edges []*Edge) []*Edge {
	var result []*Edge
	for _, e := range edges {
		if e.from.id == e.to.id {
			continue
		}
		result = append(result, e)
	}

	return result
}
func removeFromNodes(nodes []*Vertex, indices map[int]bool) []*Vertex {
	var result []*Vertex

	for i := 0; i < len(nodes); i++ {
		//_, ok := indices[i]
		if indices[i] {
			//fmt.Printf("ignoring edge index %d\n", i)
			continue
		}

		result = append(result, nodes[i])
	}

	return result
}

func removeFromEdges(edges []*Edge, indices map[int]bool) []*Edge {
	var result []*Edge

	for i := 0; i < len(edges); i++ {
		//_, ok := indices[i]
		if indices[i] {
			//fmt.Printf("ignoring edge index %d\n", i)
			continue
		}

		result = append(result, edges[i])
	}

	return result
}

func (g *graph) GetNode(id uint64) (*Vertex, bool) {
	v, ok := g.vertices[id]
	return v, ok
	/*for i := 0; i < len(g.vertices); i++ {
		if id == g.vertices[i].id {
			return g.vertices[i], true
		}
	}

	return nil, false
	*/
}

// !! this is the issue, got it
// ! ignoring duplicate edges in adjacency list for now
// (i, j) then (j, i)
// ! allow parallel edges
func (g *graph) insertNodeAdjacency(id uint64, connections []uint64) {
	//num_connections := len(connections)

	// if the node is in the graph, get it
	// otherwise create a new one
	node, ok := g.GetNode(id)
	if !ok {
		node = &Vertex{
			id:          id,
			connections: make([]*Edge, 0),
		}
	}

	// add connections
	// makes sure all nodes exist in the graph
	for _, v := range connections {
		new_vertex, ok := g.GetNode(v)
		if !ok {
			// create the node and add it
			new_vertex = new(Vertex)
			new_vertex.id = v
			new_vertex.connections = make([]*Edge, 0)
			g.vertices[v] = new_vertex
		}

		new_edge := new(Edge)
		new_edge.from = node
		new_edge.to = new_vertex

		//fmt.Printf("adding edge %v to connections of %v: %v\n", new_edge, id, node.connections)

		// add to both nodes
		node.connections = append(node.connections, new_edge)
		new_vertex.connections = append(new_vertex.connections, new_edge)
		// add to graph
		g.edges = append(g.edges, new_edge)
	}

	g.vertices[id] = node

	//g.vertices = append(g.vertices, node)

}
