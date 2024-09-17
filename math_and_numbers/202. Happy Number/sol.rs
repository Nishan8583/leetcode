/*
202. Happy Number
Solved
Easy
Topics
Companies

Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:

    Starting with any positive integer, replace the number by the sum of the squares of its digits.
    Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
    Those numbers for which this process ends in 1 are happy.

Return true if n is a happy number, and false if not.

 

Example 1:

Input: n = 19
Output: true
Explanation:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1

Example 2:

Input: n = 2
Output: false


Solution:
It is pretty straight forward.
Use a hashset to check if a number has been visited before. This prevents inifinite loops since the question predicts this is a normal occurence.
Do the required calculations and return true if we even find 1, or false if a number is visited before i.e. it is not a happy number. 
*/

use std::collections::HashSet;

impl Solution {
    

//impl Solution {
    pub fn is_happy(n: i32) -> bool {

        // holds set of visited numbers
        let mut visit: HashSet<i32> = HashSet::new(); 

        // shadow n as mutable
        let mut n = n;

        // loop till we reach same number 2x
        while !visit.contains(&n) {

            // insert into set
            visit.insert(n);

            n = Solution::sumOfSquares(n);
            //n = sumOfSquares(n);
            if n == 1 {
                return true;
            }
        }

        return false;
        
    }

    pub fn sumOfSquares(mut n: i32) -> i32 {
        let mut output = 0;

        while n > 0 {
            // get last digit
            let mut digit = n % 10;

            // square it
            digit = i32::pow(digit, 2);

            // add to output
            output += digit;
            
            // remove last digit
            n = n / 10;
        }
        return output;
    }
//}
}
pub fn main() {
    print!("{}", is_happy(19));
}