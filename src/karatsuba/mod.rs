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




use std::cmp;


fn slice_to_int(x: &[u64]) -> u64 {
  let mut result: u64 = 0;

  let n = x.len();

  for power in 0..n {
    result += x[power] * 10u64.pow(n as u32 - 1 - power as u32);
  }

  result
}


fn int_to_vec(num: u64) -> Vec<u64> {
  num.to_string().chars().map(|d| d.to_digit(10).unwrap() as u64).collect()
}



// helper function to play around with manually
// in code without reading from command line
fn string_to_vec(s: &str) -> Vec<u64> {
  s.chars().map(|d| d.to_digit(10).unwrap() as u64).collect()
}


// recursive Karatsuba multiplication
// takes vectors of only positive digits
// we use immutable read-only slices
// we return a new vector with results
// assumes |x| = |y|
fn multiply(x: &[u64], y: &[u64]) -> u64 {

  let n = x.len();

  let n32: u32 = x.len() as u32;
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


  println!("doing {:?} minus {:?} minus {:?}", step3, bd, ac);

  let ad_plus_bc = step3 - bd - ac;



  10u64.pow(n32) * ac + 10u64.pow(n32/2) * ad_plus_bc + bd


}



// trying out a string version instead
fn multiply_string(x: &str, y: &str) -> u64 {
  let n_x = x.len();
  let n_y = y.len();
  let n = cmp::max(n_x, n_y);
  let n2 = n/2; // ! this should be a floor


  // base case
  if n_x == 1 || n_y == 1 {
    return parse_int(x) * parse_int(y);
  }

  // get the pieces
  let a = &x[0..n_x - n2];
  let b = &x[n_x - n2..n_x];
  let c = &y[0..n_y - n2];
  let d = &y[n_y - n2..n_y];

  let a_plus_b = parse_int(a) + parse_int(b);
  let c_plus_d = parse_int(c) + parse_int(d);

  //let a_plus_b_string = a_plus_b.to_string();
  //let c_plus_d


  let ac = multiply_string(a, c);
  let step3 = multiply_string(&a_plus_b.to_string(), &c_plus_d.to_string());
  let bd = multiply_string(b, d);

  let ad_plus_bc = step3 - ac - bd;

  // ! consider adding trailing zeros to string instead
  let term1 = parse_int(&add_trailing_zeros(&ac.to_string(), n2 as u64 * 2));
  let term2 = parse_int(&add_trailing_zeros(&ad_plus_bc.to_string(), n2 as u64));
  let term3 = bd;

  term1 + term2 + term3
}

fn add_trailing_zeros(string: &str, numzeros: u64) -> String {
  let mut owned_string: String = string.to_owned();
  for i in 0..numzeros {
    owned_string.push_str("0");
  }

  owned_string
}

fn parse_int(input: &str) -> u64 {
  input.chars()
    .find(|a| a.is_digit(10))
    .and_then(|a| a.to_digit(10))
    .unwrap() as u64
}

#[test]
fn test_add_trailing_zeros() {
  let x = "101";
  let x_big = &add_trailing_zeros(x, 3);
  assert_eq!(x_big, "101000");
}


#[test]
fn test_karatsuba_string() {
  let x = "5678";
  let y = "1234";
  let result = multiply_string(x, y);
  assert_eq!(result, 7006652, "{:?} times {:?} equals {:?}.", x, y, result);
}


#[test]
#[ignore]
fn test_karatsuba() {
  println!("testing karatsuba!");




  let x = &[5,6,7,8];
  let y = &[1,2,3,4];
  let result = multiply(x, y);

  // answer should be 7006652;
  assert_eq!(result, 7006652, "{:?} times {:?} equals {:?}.", x, y, result);

  return;

  //let p_str = "3141592653589793238462643383279502884197169399375105820974944592";
  //let q_str = "2718281828459045235360287471352662497757247093699959574966967627";
  let p_str = "3141592653589";//79323846264338";//3279502884197169399375105820974944592";
  let q_str = "2718281828459";//04523536028747";//1352662497757247093699959574966967627";
  let p = &string_to_vec(p_str);
  let q = &string_to_vec(q_str);



  let pqresult = multiply(p,q);
  println!("pqresult {:?}", pqresult);
}

#[test]
fn test_int_to_vec() {
  let slice = &vec![3,4,5];
  let int = 345;
  let convert = &int_to_vec(int);
  assert_eq!(convert, slice, "slice {:?} and int {:?}", slice, int);

  let slice2 = &vec![4];
  let int2 = 4;
  let convert2 = &int_to_vec(int2);
  assert_eq!(convert2, slice2, "slice {:?} and int {:?}", slice2, int2);
}


#[test]
fn test_slice_to_int() {
  let slice = &vec![1,2,3];
  let int = 123;
  let convert = slice_to_int(slice);

  assert_eq!(convert, int, "slice {:?} and int {:?}", slice, int);

  let slice2 = &vec![4];
  let int2 = 4;
  let convert2 = slice_to_int(slice2);
  assert_eq!(convert2, int2, "slice {:?} and int {:?}", slice2, int2);

  let slice3 = &vec![0,0,7];
  let int3 = 7;
  let convert3 = slice_to_int(slice3);
  assert_eq!(convert3, int3, "slice {:?} and int {:?}", slice3, int3);


}

