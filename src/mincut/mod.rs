



struct Vertex<'a> {
  id: u64,
  edges: Vec<Edge<'a>>,
}

struct Edge<'a> {
  first: &'a Vertex<'a>,
  second: &'a Vertex<'a>,
}

struct AdjacencyList<'a> {
  vertices: Vec<Vertex<'a>>,
  edges: Vec<Edge<'a>>,
}




