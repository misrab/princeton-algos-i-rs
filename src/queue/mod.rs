use std::marker::PhantomData;

// implementing a queue using a linked list just to practice
struct Item<T> {
    item: T,
}


pub struct Queue<T> {
    num_items: u32,
    start: Item<T>,

    // moo: T,
}

impl <T> Queue<T> {
    fn new(&self) -> Queue<T> {
        Queue {
            num_items: 32,
            start: PhantomData<Item<T>>,
        }
    }
}



#[test]
fn test_new() {
    println!("testing queue");
}
