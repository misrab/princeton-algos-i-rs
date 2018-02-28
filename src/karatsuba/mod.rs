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


extern crate num_bigint;

use std::cmp;
//use self::bigint::BigUint;
use self::num_bigint::{BigUint,ToBigInt};



// trying out a string version instead
fn multiply_string(x: &str, y: &str) -> BigUint {
  let n_x = x.len();
  let n_y = y.len();
  let n = cmp::max(n_x, n_y);
  let n2 = n/2; // ! this should be a floor

  //println!("multiplying {} and {}", x, y);


  // base case
  if n_x == 1 || n_y == 1 {
    //println!("Returning base case {:?} * {:?} = {:?}", x, y, parse_int(x)*parse_int(y));
    return parse_int(x) * parse_int(y);
  }

  // get the pieces
  let a = &x[0..n_x - n2];
  let b = &x[n_x - n2..n_x];
  let c = &y[0..n_y - n2];
  let d = &y[n_y - n2..n_y];

  let a_plus_b = parse_int(a) + parse_int(b);
  let c_plus_d = parse_int(c) + parse_int(d);

  let ac = multiply_string(a, c);
  let step3 = multiply_string(&a_plus_b.to_string(), &c_plus_d.to_string());
  let bd = multiply_string(b, d);

  //println!("doing {:?} - {:?} - {:?}, for x {:?} and y {:?}", step3, ac, bd, x, y);

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

fn parse_int(input: &str) -> BigUint {
  //input.parse::<BigUint>().unwrap()
  /*let i = match input.parse::<BigUint>() {
    Ok(i) => i,
    Err(e) => {
      println!("error parsing {:?}: {:?}", input, e);
      0
    }
  };

  i*/

  //BigUint::from_dec_str(input).unwrap()
  BigUint::parse_bytes(input.as_bytes(), 10).unwrap()

}


#[test]
fn test_parse_int() {
  assert_eq!(parse_int("11"), ToBigInt::to_bigint(&11).unwrap());
  //assert_eq!(parse_int("3"), BigUint::from(3));
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


  let p = "3141592653589";//79323846264338";//3279502884197169399375105820974944592";
  let q = "2718281828459";//04523536028747";//1352662497757247093699959574966967627";
  let pqresult = multiply_string(p,q);
  println!("{:?}", pqresult);



}


/*
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
*/
