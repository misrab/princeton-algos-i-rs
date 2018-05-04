package graph

import ()

type Vertex struct {
	id          uint64
	connections []*Edge
}

type Edge struct {
	from *Vertex
	to   *Vertex
}

type graph struct {
	vertices map[uint64]*Vertex
	edges    []*Edge
}

type Graph interface {
	//insertNode(id uint64)

	insertNodeAdjacency(id uint64, connections []uint64)

	getNodeAdjacency(id uint64) []*Edge

	getNode(uint64) (*Vertex, bool)
}

func NewGraph() Graph {
	graph := new(graph)

	// allocate memory
	graph.vertices = make(map[uint64]*Vertex)
	//graph.edges = make([]*Edge)

	return graph
}

func (g *graph) getNode(id uint64) (*Vertex, bool) {
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

func (g *graph) getNodeAdjacency(id uint64) []*Edge {
	node, ok := g.getNode(id)

	if !ok {
		return nil
	}
	return node.connections
}

//func (g *graph) insertNode(id uint64) {
//}

func (g *graph) insertNodeAdjacency(id uint64, connections []uint64) {
	num_connections := len(connections)

	// if the node is in the graph, get it
	// otherwise create a new one
	node, ok := g.getNode(id)
	if !ok {
		node = &Vertex{
			id:          id,
			connections: make([]*Edge, num_connections),
		}
	}

	// TODO add connections
	// makes sure all nodes exist in the graph
	for _, v := range connections {
		n, ok := g.getNode(v)
		if !ok {
			// create the node and add it
			g.vertices[v] = &Vertex{
				id:          v,
				connections: make([]*Edge, 0),
			}
		}

		new_edge := &Edge{
			from: node,
			to:   n,
		}

		node.connections = append(node.connections, new_edge)
	}

	g.vertices[id] = node

	//g.vertices = append(g.vertices, node)

}
