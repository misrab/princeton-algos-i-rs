package graph

import (
	"fmt"
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

type graph struct {
	vertices map[uint64]*Vertex
	edges    []*Edge
}

type Graph interface {
	//insertNode(id uint64)

	insertNodeAdjacency(id uint64, connections []uint64)

	//GetNodeAdjacency(id uint64) []*Edge

	GetNode(uint64) (*Vertex, bool)
}

func NewGraph() Graph {
	graph := new(graph)

	// allocate memory
	graph.vertices = make(map[uint64]*Vertex)
	graph.edges = make([]*Edge, 0)

	return graph
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
