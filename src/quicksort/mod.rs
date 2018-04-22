

pub enum PivotMethod {
  First,
  Last,
  MedianOfThree,
}




// return (index, value) of first element
fn choose_pivot_median_of_three(arr: &mut Vec<u64>, n: usize) -> (usize, u64) {
  // TEMP
  (0, arr[0])
}

// assumes first element is the pivot
fn partition(arr: &mut Vec<u64>, beginning: usize, end: usize) -> usize {
  let mut i = 0; // keep track of partition cutoff
  //let mut j = 0; // keep track of elements compared to pivot

  for j in beginning..end {

  }
}

fn swap(arr: &mut Vec<u64>, i: usize, j: usize) {
  let n = arr.len();

  if i == j { return; }
  if i < 0 || j < 0 { return; }
  if i > n-1 || j > n-1 { return; }

  let temp = arr[i];
  arr[i] = arr[j];
  arr[j] = temp;
}


// start and end are inclusive
pub fn quicksort(arr: &mut Vec<u64>, start: usize, end: usize, pivot_method: PivotMethod) {
  let n = end - start + 1;
  if n == 1 { return; }

  let (pivot_index, pivot_value) = match pivot_method {
    First => (0, arr[0]),
    Last => (n-1, arr[n-1]),
    MedianOfThree => choose_pivot_median_of_three(arr, n),
  };

  // swap pivot with first element
  swap(arr, 0, pivot_index);

  // partition the current array then recurse
  let partition_index = partition(arr, 0, n-1);

  // recurse
  //quicksort(arr, 0, partition_index-1, pivot_method);
  //quicksort(arr, partition_index+1, pivot_method);

}



#[test]
fn test_swap() {
  let mut vec = vec![1,2,3,4];
  swap(&mut vec, 0,2); // 3 2 1 4
  swap(&mut vec, 1,3); // 3 4 1 2
  swap(&mut vec, 1,2); // 3 1 4 2

  let answer = vec![3,1,4,2];
  assert_eq!(vec, answer, "{:?} should equal {:?}", vec, answer);
}


#[test]
fn test_quicksort() {
  println!("testing quicksort");

  let mut arr = vec![3,1,4,2];
  let n = arr.len();
  quicksort(&mut arr, 0, n-1, PivotMethod::First);
  println!("Sorted array is {:?}", arr);
}
