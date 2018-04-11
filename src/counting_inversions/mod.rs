










fn count_brute_inversions(arr: &[u64]) -> u64 {
  let n = arr.len();

  let mut count = 0;
  for i in 0..n {
    for j in (i+1)..n {
      if arr[i] > arr[j] { count += 1; }
    }
  }

  count
}


//use std::cmp::max;

fn sort_and_count_inversions(arr: &[u64]) -> (Vec<u64>, u64) {
  let n = arr.len();
  println!("arry length {:} for {:?}", n, arr);


  if n <= 1 { return (arr.to_vec(), 0); }

  let mid = (n as f64 * 0.5).ceil() as usize;

  let (B, x) = sort_and_count_inversions(&arr[0..mid]);
  let (C, y) = sort_and_count_inversions(&arr[mid..n-1]);

  let z = count_split_inversions(&arr);

  let mut sorted = arr.to_vec();
  sorted.sort();

  (sorted, x + y)
}


fn count_split_inversions(arr: &[u64]) -> u64 {
  0
}


#[test]
fn test_brute() {
  println!("testing counting inversions");

  let arr = vec![3,6,5,1];

  let brute = count_brute_inversions(&arr);
  assert_eq!(brute, 4, "array {:?} has {:?} inversions", arr, brute);
}


#[test]
fn test_efficient() {
  let arr = vec![32,4324,11,22,3454,21,563,987,5634,2321];

  let brute = count_brute_inversions(&arr);
  let (_, efficient) = sort_and_count_inversions(&arr);

  assert_eq!(brute, efficient);
}
