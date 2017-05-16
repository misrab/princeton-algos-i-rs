/*
    Queue implementation with doubly linked list
*/

use std::rc::Rc;
use std::marker::PhantomData;


// implementing a queue using a linked list just to practice
struct Node<T> {
    item: T,
    next: Link<T>,
    prev: Link<T>,
}

pub struct Queue<'a, T: 'a> {
    head: Link<T>,
    tail: Link<T>,

    phantom: PhantomData<&'a T>,
}

type Link<T> = Option<Rc<Node<T>>>;


impl <'a, T> Queue<'a, T> {
    fn new() -> Self {
        Queue {
            head: None,
            tail: None,

            phantom: PhantomData,
        }
    }

    fn add(&mut self, el: T) {
        let item = Rc::new(Node {
            item: el,
            next: self.head.take(),
        });

        // for use in head later
        let _y = item.clone();

        // if list was empty, point last to this element
        match self.tail {
            None =>  { self.tail = Some(_y); },
            Some(_) => {},
        }

        self.head = Some(item);
    }

    // remove the last element
    fn remove(&mut self) -> Option<T> {
        // match self.head {
        //     None => { return None; },
        //     Some(_) => {},
        // }
        //
        // match self.tail.take() {
        //     None => {},
        //     Some(last) => {
        //         match last.next {
        //             // it was the last element
        //             None => {
        //                 self.tail = None;
        //
        //             },
        //             Some(_) => {
        //
        //             },
        //         }
        //     }
        // }


        unimplemented!();
    }
}



#[test]
fn test_new() {
    println!("testing queue");

    let q: Queue<u32> = Queue::new();
}
