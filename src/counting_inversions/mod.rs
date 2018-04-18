










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
  //println!("arry length {:} for {:?}", n, arr);


  if n <= 1 { return (arr.to_vec(), 0); }

  let mid = n/2; //(n as f64 / 2).floor() as usize;
  //println!("splitting at index {:?} out of {:?}", mid, n);


  let (left_sorted, x) = sort_and_count_inversions(&arr[0..mid]);
  let (right_sorted, y) = sort_and_count_inversions(&arr[mid..n]);

  let (full_sorted, z) = count_split_inversions(&left_sorted, &right_sorted);

  //let mut sorted = arr.to_vec();
  //sorted.sort();

  (full_sorted, x + y + z)
}


fn count_split_inversions(left: &[u64], right: &[u64]) -> (Vec<u64>, u64) {
  let n = left.len();

  let mut count = 0;
  let mut left_index = 0;
  let mut right_index = 0;

  let mut full_sorted = Vec::new();

  while left_index < n && right_index < n {
    if left[left_index] < right[right_index] {
      full_sorted.push(left[left_index]);
      left_index += 1;
      continue;
    }

    // we have an inversion, and so all remaining elements
    // in the left will also be inversions
    full_sorted.push(right[right_index]);
    right_index += 1;
    count += n as u64 - left_index as u64 + 1;
  }


  (full_sorted, count)
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
