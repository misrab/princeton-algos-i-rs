

pub enum PivotMethod {
  First,
  Last,
}




// return first element as pivot
fn choose_pivot_first() -> {

}




pub fn quicksort(arr: &mut Vec<u64>, n: usize, pivot_method: PivotMethod) {
  if n == 1 { return; }

  let p = match pivot_method {
    First => choose_pivot_first(),
    Last => choose_pivot_last(),
  }
}






#[test]
fn test_quicksort() {
  println!("testing quicksort");
}
