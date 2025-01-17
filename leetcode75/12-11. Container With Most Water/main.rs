use std::cmp::{max, min};

pub fn max_area(height: Vec<i32>) -> i32 {
    let mut left = 0;
    let mut right = height.len() - 1;
    let mut area_max = 0;

    while left < right {
        let min_height = min(height[left], height[right]) as usize;
        let width = right - left;

        area_max = max(area_max, min_height * width);

        if height[left] < height[right] {
            left += 1;
        } else {
            right -= 1;
        }
    }

    return area_max as i32;
}

fn main() {
    println!("{}", max_area(vec![1, 8, 6, 2, 5, 4, 8, 3, 7]));
}
