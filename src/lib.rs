pub mod karatsuba;


pub mod counting_inversions;

pub mod quicksort;


//pub mod union_find;


//mod stack;

// pub mod queue;
// pub mod dequeue;
//
//



use std::io::BufReader;
use std::io::BufRead;
use std::fs::File;
use std::path::Path;

// file reading helpers
fn read_file_to_vec(path: &str) -> Vec<u64> {
  let mut result = Vec::new();

  let f = File::open(path).unwrap();
  let mut file = BufReader::new(&f);
  for line in file.lines() {
    let l = line.unwrap();
    //println!("{}", l); 
    result.push(l.parse::<u64>().unwrap());
  }

  result
}



// use this to run programming assignments
#[test]
#[ignore]
fn test_quicksort() {
  let mut vec = read_file_to_vec("/home/misrab/code/src/github.com/misrab/stanford-algos-rs/data/quicksort.txt");

  let n = vec.len();
  quicksort::quicksort(&mut vec, 0, n, &quicksort::PivotMethod::First);

  println!("{:?}", vec);
}


#[test]
#[ignore]
fn test_split_inversions() {
  println!("moo");

  let vec = read_file_to_vec("/home/misrab/code/src/github.com/misrab/stanford-algos-rs/data/intarray.txt");

  //let arr = vec![1,2,3];
  //counting_inversions::count_brute_inversions(&arr);
  //
  let (_, count) = counting_inversions::sort_and_count_inversions(&vec);
  println!("Found {:} inversions", count);

}
