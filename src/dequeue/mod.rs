// public class Deque<Item> implements Iterable<Item> {
//    public Deque()                           // construct an empty deque
//    public boolean isEmpty()                 // is the deque empty?
//    public int size()                        // return the number of items on the deque
//    public void addFirst(Item item)          // add the item to the front
//    public void addLast(Item item)           // add the item to the end
//    public Item removeFirst()                // remove and return the item from the front
//    public Item removeLast()                 // remove and return the item from the end
//    public Iterator<Item> iterator()         // return an iterator over items in order from front to end
//    public static void main(String[] args)   // unit testing (optional)
// }

struct Item<T> {
    item: T,
    next: *mut Item<T>,
}

pub struct Deque<T> {
  n: u32,
  start: Item<T>,
  // moo: T,
}

impl <T> Deque<T> {
  pub fn new(&self) -> Deque<T> {
    Deque<T> {
        n: 0,
        // moo: T,
    }
  }

  pub fn addFirst(&self, item: T) {

  }

    // pub fn isEmpty(&self) -> bool {
    //
    // }
    //
    // pub fn size(&self) -> i32 {
    //
    // }
    //
    // pub fn addFirst(&self, item: T) {
    //
    // }
    //
    // pub fn addLast(&self, item: T) {
    //
    // }
    //
    // pub fn removeFirst(&self) -> T {
    //
    // }
    //
    // pub fn removeLast(&self) -> T {
    //
    // }

    // Iterator TODO
}



#[test]
fn test_dequeue_new() {
    println!("{:?}", "testing new dequeue");
    // let x = Deque { n: 3 };
}
