/*
 *
 *  Recall: xy = 10^n ac + 10^(n/2) (ad + bc) + bd
 *  Step 1: recursively compute ac
 *  Step 2: " " bd
 *  Step 3: compute (a+b)(c+d) = ac + ad + bc + bd
 *  Gauss' trick: (3) - (1) - (2) = ad + bc
 *  Upshot: only need 3 recursive multiplications! (and some additions)
 *
 *
 *
 *
 *  Notes:
 *    - I suspect more efficient ways of dealing with digits and slice conversions
*/


//use std::u64::pow;




fn slice_to_int(x: &[u64]) -> u64 {
  let mut result: u64 = 0;

  let n = x.len();

  for power in 0..n {
    result += x[power] * 10u64.pow(n as u64 - 1 -power as u64);
  }

  result
}


fn int_to_vec(num: u64) -> Vec<u64> {
  num.to_string().chars().map(|d as u64| d.to_digit(10).unwrap()).collect()
}


// recursive Karatsuba multiplication
// takes vectors of only positive digits
// we use immutable read-only slices
// we return a new vector with results
// assumes |x| = |y|
fn multiply(x: &[u64], y: &[u64]) -> u64 {

  let n = x.len();
  let n32: u64 = x.len() as u64;
  //let n_y = y.len();


  //println!("multiplying {:?} and {:?}", x, y);

  // base case
  // second condition is where we're adding in step 3 below
  // and thus lengths may not match
  if n32 == 1 || x.len() != y.len() {
    return slice_to_int(x)*slice_to_int(y);
  }



  let a = &x[0..n/2];
  let b = &x[n/2..n];
  let c = &y[0..n/2];
  let d = &y[n/2..n];


  let ac = multiply(a, c);

  let bd = multiply(b, d);


  let step3 = multiply(&int_to_vec(slice_to_int(a)+slice_to_int(b)), &int_to_vec(slice_to_int(c)+slice_to_int(d)));
  let ad_plus_bc = step3 - bd - ac;



  10u64.pow(n32) * ac + 10u64.pow(n32/2) * ad_plus_bc + bd


}


// helper function to play around with manually
// in code without reading from command line
fn string_to_vec(s: &str) -> Vec<u64> {
  s.chars().map(|d as u64| d.to_digit(10).unwrap()).collect()
}



#[test]
fn test_int_to_vec() {
  let slice = &vec![3,4,5];
  let int = 345;
  let convert = &int_to_vec(int);

  assert_eq!(convert, slice, "slice {:?} and int {:?}", slice, int);
}


#[test]
fn test_slice_to_int() {
  let slice = &vec![1,2,3];
  let int = 123;
  let convert = slice_to_int(slice);

  assert_eq!(convert, int, "slice {:?} and int {:?}", slice, int);

}

#[test]
fn test_karatsuba() {
  println!("testing karatsuba!");

  let x = &[5,6,7,8];
  let y = &[1,2,3,4];
  let result = multiply(x, y);

  // answer should be 7006652;
  assert_eq!(result, 7006652, "{:?} times {:?} equals {:?}.", x, y, result);


  let p_str = "3141592653589793238462643383279502884197169399375105820974944592";
  let q_str = "2718281828459045235360287471352662497757247093699959574966967627";
  let p = &string_to_vec(p_str);
  let q = &string_to_vec(q_str);
  let pqresult = multiply(p,q);
  println!("pqresult {:?}", pqresult);
}
