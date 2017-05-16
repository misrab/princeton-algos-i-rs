


pub struct UnionFind {
    // number of items
    n: i32,

    // parents[i] is parent of i
    parents: Vec<i32>,

    // number of objects rooted at i
    sizes: Vec<i32>,

}

impl UnionFind {
    pub fn new(n: i32) -> Self {
        let mut parents: Vec<i32> = Vec::with_capacity(n as usize);
        let mut sizes: Vec<i32> = Vec::with_capacity(n as usize);

        for i in 0..n {
            // parents[i as usize] = i;
            // sizes[i as usize] = 1;
            parents.push(i);
            sizes.push(1);
        }

        // println!("{:?}", parents);

        UnionFind {
            n: n,
            parents: parents,
            sizes: sizes,
        }
    }


    // find the roots of each node
    // for the smaller tree, make all nodes point to root of bigger tree
    pub fn union(&mut self, a: i32, b: i32) {
        let root_a = self.root(a);
        let root_b = self.root(b);
        if root_a == root_b { return; }

        let size_root_a = self.sizes[root_a as usize];
        let size_root_b = self.sizes[root_b as usize];
        if size_root_a < size_root_b {
            // self.repoint_tree(a, root_b);
            self.parents[root_a as usize] = root_b;
            self.sizes[root_b as usize] += self.sizes[root_a as usize];

        } else {
            // self.repoint_tree(b, root_a);
            self.parents[root_b as usize] = root_a;
            self.sizes[root_a as usize] += self.sizes[root_b as usize];
        }
    }


    pub fn connected(&self, a: i32, b: i32) -> bool {
        self.root(a) == self.root(b)
    }


    // take the node and all its parents, have them
    // point to root as their new parent.
    // update tree sizes!
    // fn repoint_tree(&self, node: i32, root: i32) {
    //     let mut parent = self.parents[node as usize];
    //     let mut curr = node;
    //     self.parents[curr as usize] = root;
    //
    //     let mut height = 0;
    //     while curr != parent {
    //         height += 1;
    //
    //         curr = parent;
    //         parent = self.parents[curr as usize];
    //
    //         self.parents[curr as usize] = root;
    //     }
    //
    //     // update sizes
    //     sizes[root as usize] = height;
    // }

    // get root of node
    fn root(&self, a: i32) -> i32 {
        let mut parent = self.parents[a as usize];
        let mut curr = a;
        while curr != parent {
            curr = parent;
            parent = self.parents[curr as usize];
        }

        parent
    }


}


pub fn union() {
    println!("called `my::function()`");
}




#[test]
#[ignore]
fn union_find() {
    let mut uf = UnionFind::new(36);

    assert!(uf.connected(1,2) == false);

    uf.union(1,2);
    assert!(uf.connected(1,2) == true);
}
