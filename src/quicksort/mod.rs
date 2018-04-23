

pub enum PivotMethod {
  First,
  Last,
  MedianOfThree,
}




// return (index, value) of first element
fn choose_pivot_median_of_three(arr: &mut Vec<u64>, n: usize) -> (usize, u64) {
  // TEMP
  // bookkeeping
  (0, arr[0])
}

// assumes first element is the pivot
// i is partition cutoff
// j is element being compared to pivot
fn partition(mut arr: &mut Vec<u64>, beginning: usize, end: usize) -> usize {
  let p = arr[beginning];

  let mut i = beginning + 1;

  //println!("{:?}", arr);

  for j in (beginning + 1)..end {
    if arr[j] < p {
      swap(&mut arr, i, j);
      i += 1;
      //println!("{:?}", arr);
    }
  }

  // put array in its spot
  swap(&mut arr, beginning, i-1);
  //println!("{:?}", arr);

  // return pivot index
  i-1
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


// start <= . < end
// returns number of comparisons made
pub fn quicksort(arr: &mut Vec<u64>, start: usize, end: usize, pivot_method: &PivotMethod) -> u64 {
  //println!("running quicksort from {:?} <= ... < {:?}", start, end);
  if end == 0 { return 0; }
  let n = end - start;
  if n <= 1 { return 0; }

  let (pivot_index, pivot_value) = match pivot_method {
    ref First => (start, arr[start]),
    ref Last => (end-1, arr[end-1]),
    ref MedianOfThree => choose_pivot_median_of_three(arr, n),
  };

  // swap pivot with first element
  swap(arr, start, pivot_index);


  let mut comparisons = n as u64 - 1;

  // partition the current array then recurse
  let partition_index = partition(arr, start, end);

  //println!("partition index of {:?} is {:?}, we now have {:?}", pivot_value , partition_index, arr);

  // recurse
  comparisons += quicksort(arr, start, partition_index, pivot_method);
  comparisons += quicksort(arr, partition_index+1, end, pivot_method);

  comparisons
}


#[test]
#[ignore]
fn test_num_comparisons() {
  let mut arr = vec![3,4,1,2,0];
  let n = arr.len();
  let comparisons = quicksort(&mut arr, 0, n, &PivotMethod::First);


  assert_eq!(6, comparisons);
}


#[test]
#[ignore]
fn test_partition() {
  let mut arr = vec![22, 312, 322, 47, 87, 4321];
  let n = arr.len();
  let index = partition(&mut arr, 1, 5);

  //println!("partition index of 4 is {:?} in {:?}", index, arr);
  //assert_eq!(index, 4);
}



#[test]
#[ignore]
fn test_swap() {
  let mut vec = vec![1,2,3,4];
  swap(&mut vec, 0,2); // 3 2 1 4
  swap(&mut vec, 1,3); // 3 4 1 2
  swap(&mut vec, 1,2); // 3 1 4 2

  let answer = vec![3,1,4,2];
  assert_eq!(vec, answer, "{:?} should equal {:?}", vec, answer);
}


#[test]
#[ignore]
fn test_quicksort() {
  println!("testing quicksort");

  //let mut arr = vec![3,1,4,2]; 
  let mut arr = vec![4321,312,5435,6767,322,47,87,22,3435];
  //let mut arr = vec![4321,312,322,47,87,22];


  let n = arr.len();
  quicksort(&mut arr, 0, n, &PivotMethod::First);
  println!("Sorted array is {:?}", arr);
}
