#[cfg(test)]

use union_find::UnionFind;


#[test]
fn union_find() {
    let mut uf = UnionFind::new(36);

    assert!(uf.connected(1,2) == false);

    uf.union(1,2);
    assert!(uf.connected(1,2) == true);
}



//
//
// mod tests {
//     #[test]
//     fn it_works() {
//     }
// }
