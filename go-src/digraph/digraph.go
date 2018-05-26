package digraph

import (
	"fmt"

	"github.com/misrab/go-utils"
)

/*
  structs and interfaces
*/

type Vertex struct {
	id       uint64
	incoming []*Edge
	outgoing []*Edge
}

type Edge struct {
	from *Vertex
	to   *Vertex
}

type digraph struct {
	vertices map[uint64]*Vertex
	edges    []*Edge
}

type DiGraph interface {
	AddEdge(from, to uint64) *Edge
	AddVertex(id uint64) *Vertex

	RemoveEdge(e *Edge)

	GetVertex(id uint64) (*Vertex, bool)

	GetVertexIds() []uint64

	Connected(from, to uint64) bool

	Copy() DiGraph
	Reverse() DiGraph
}

/*
  constructors
*/

func NewDiGraph() *digraph {
	digraph := new(digraph)

	digraph.vertices = make(map[uint64]*Vertex)

	return digraph
}

// returns a graph with new memory allocated
func (d *digraph) Copy() DiGraph {
	new_graph := NewDiGraph()

	for _, edge := range d.edges {
		new_graph.AddEdge(edge.from.id, edge.to.id)
	}

	return new_graph
}

func (d *digraph) Reverse() DiGraph {

	return d
}

/*
  string formatting
*/

func (e *Edge) String() string {
	return fmt.Sprintf("(%d,%d)", e.from.id, e.to.id)
}

func (d *digraph) String() string {
	result := "edges:\n"
	for _, e := range d.edges {
		result += fmt.Sprintf("%v ", e)
	}
	result += "\nvertices:\n"

	result += fmt.Sprintf("%v\n", d.GetVertexIds())

	return result
}

/*
  public methods
*/

func (d *digraph) GetVertexIds() []uint64 {
	n := len(d.vertices)

	ids := make([]uint64, n)

	i := 0
	for id, _ := range d.vertices {

		ids[i] = id
		i += 1
	}

	return ids
}

func (d *digraph) Connected(from, to uint64) bool {
	from_vertex, from_found := d.GetVertex(from)
	if !from_found {
		return false
	}

	for _, edge := range from_vertex.outgoing {
		if edge.to.id == to {
			return true
		}
	}

	return false
}

func (d *digraph) GetVertex(id uint64) (*Vertex, bool) {
	v, ok := d.vertices[id]

	return v, ok
}

func (d *digraph) AddVertex(id uint64) *Vertex {
	// ignore if vertex already in graph
	if found, ok := d.vertices[id]; ok {
		return found
	}

	vertex := new(Vertex)
	vertex.id = id
	vertex.incoming = make([]*Edge, 0)
	vertex.outgoing = make([]*Edge, 0)

	d.vertices[id] = vertex

	return vertex
}

func (d *digraph) AddEdge(from, to uint64) *Edge {
	// ensure both vertices exist in graph
	var from_vertex, to_vertex *Vertex
	var ok bool

	if from_vertex, ok = d.vertices[from]; !ok {
		from_vertex = d.AddVertex(from)
	}

	if to_vertex, ok = d.vertices[to]; !ok {
		to_vertex = d.AddVertex(to)
	}

	// add to adjacency list
	// duplicate edges are possible
	edge := new(Edge)
	edge.from = from_vertex
	edge.to = to_vertex

	d.edges = append(d.edges, edge)
	from_vertex.outgoing = append(from_vertex.outgoing, edge)
	to_vertex.incoming = append(to_vertex.incoming, edge)

	return edge
}

func (d *digraph) RemoveEdge(edge *Edge) {
	// remove from graph list
	index := utils.Find(len(d.edges), func(i int) bool { return d.edges[i] == edge })
	if index != -1 {
		d.edges = append(d.edges[:index], d.edges[index+1:]...)
	}

	// remove from relevant vertices
	from, _ := d.GetVertex(edge.from.id)
	to, _ := d.GetVertex(edge.to.id)

	index = utils.Find(len(from.outgoing), func(i int) bool { return from.outgoing[i] == edge })
	if index != -1 {
		from.outgoing = append(from.outgoing[:index], from.outgoing[index+1:]...)
	}

	index = utils.Find(len(to.incoming), func(i int) bool { return to.incoming[i] == edge })
	if index != -1 {
		to.incoming = append(to.incoming[:index], to.incoming[index+1:]...)
	}

	// GC should clear edge from memory now
}

/*
	private methods
*/
