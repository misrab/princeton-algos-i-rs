
pub struct Queue<T> {
    head: Link<T>,
}

type Link<T> = Option<Box<Node<T>>>;

struct Node<T> {
    elem: T,
    next: Link<T>,
}



impl<T> Queue<T> {
    pub fn new() -> Self {
        Queue { head: None }
    }

    pub fn enque(&mut self, elem: T) {
        let new_node = Box::new(Node {
            elem: elem,
            next: self.head.take(),
        });

        self.head = Some(new_node);
    }

    pub fn deque(&mut self) -> Option<T> {
        // self.head.take().map(|node| {
        //     let node = *node;
        //     self.head = node.next;
        //     node.elem
        // })

        let mut cur_link: Link<T> = self.head.take();
        while let Some(mut boxed_node) = cur_link {
            cur_link = boxed_node.next.take();
            // match boxed_node.next {
            //     None => {},
            //     Some(x) => {
            //         cur_link = boxed_node.next.take();
            //     },
            // }
        }

        cur_link
    }
}

impl<T> Drop for Queue<T> {
    fn drop(&mut self) {
        let mut cur_link = self.head.take();
        while let Some(mut boxed_node) = cur_link {
            cur_link = boxed_node.next.take();
        }
    }
}


#[test]
fn test_stack() {
    println!("testing queue");

    let mut queue: Queue<i32> = Queue::new();
    queue.enque(1);
    queue.enque(2);
    queue.enque(3);
    assert!(queue.deque().unwrap() == 1);
    assert!(queue.deque().unwrap() == 2);
    assert!(queue.deque().unwrap() == 3);
}
