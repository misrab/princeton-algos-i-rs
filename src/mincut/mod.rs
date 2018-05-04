use std::collections::HashMap;


// think I should've used a heap here


struct Vertex<'a> {
  id: u64,
  edges: Vec<Edge<'a>>,
}

struct Edge<'a> {
  first: &'a Vertex<'a>,
  second: &'a Vertex<'a>,
}

struct AdjacencyList<'a> {
  vertices: HashMap<u64, &'a Vertex<'a>>, //Vec<Vertex<'a>>,
  edges: Vec<Edge<'a>>,
}



impl<'a> AdjacencyList<'a> {
  pub fn new() -> Self {
    AdjacencyList{
      vertices: HashMap::new(),
      edges: Vec::new(),
    }
  }

  pub fn node_present(&mut self, id: u64) -> bool {
    self.vertices.contains_key(&id)
  }


  pub fn insert_node(&mut self, id: u64) -> &Vertex {
    if self.vertices.contains_key(&id) { return &self.vertices[&id]; }

    let vertex = Vertex{
      id: id,
      edges: Vec::new(),
    };

    self.vertices.insert(id, &vertex);

    &vertex
  }

  // takes a vector where the first item is the node id
  // the remaining items are the other nodes it's connected to
  // ! assume we haven't seen this node's list before (for now)
  pub fn insert_adj_list(&mut self, list: Vec<u64>) {
    let n = list.len();

    if n == 0 { return; }

    let node = Vertex{
        id: list[0],
        edges: Vec::new(),
    };

    let connections: Vec<Edge> = Vec::new();

    for i in 1..n {
      // if the node is not in the graph, add it
      let second = self.insert_node(list[i]);

      // add the edge to this node's list
      let edge = Edge{
        first: &node,
        second: &second,
      };
    }

    // add note
  }

  //pub fn moo(&mut self) { self.moo = ... }
}


#[test]
fn test_new_adjacency_list() {
  let mut graph = AdjacencyList::new();

  let list = vec![1,2,3,4];
  graph.insert_adj_list(list);
}


#[test]
fn test_insert_node() {
  let mut graph = AdjacencyList::new();

  graph.insert_node(1);
  graph.insert_node(2);
  graph.insert_node(5);

  assert!(graph.node_present(1));
  assert!(graph.node_present(2));
  assert!(graph.node_present(5));
  assert!(!graph.node_present(4));


}

