use std::cmp::max;
impl Solution {
    pub fn kids_with_candies(candies: Vec<i32>, extra_candies: i32) -> Vec<bool> {
        let greatest = candies.iter().max().unwrap();

        let mut result = vec![false; candies.len()];

        for (index, value) in candies.iter().enumerate() {
            if *value + extra_candies > *greatest {
                result[index] = true;
            }
        }
        return result;
    }
}

pub fn optimized(candies: Vec<i32>, extra_candies: i32) -> Vec<bool> {
    if let Some(&greatest) = candies.iter().max() {
        candies
            .iter()
            .map(|&value| value + extra_candies >= greatest)
    } else {
        Vec::new()
    }
}
