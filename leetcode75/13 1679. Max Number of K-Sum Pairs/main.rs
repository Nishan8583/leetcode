pub fn max_operations(nums: Vec<i32>, k: i32) -> i32 {
    let mut c = nums.clone();
    c.sort();

    let mut left = 0;
    let mut right = c.len() - 1;
    while left >= right {
        let mid = (left + right) / 2;
        if c[mid] > k {
            c = (&c[..mid]).to_vec();
            right = mid;
        } else {
            left = mid + 1;
        }
    }
    left = 0;
    right = c.len() - 1;
    let mut total = 0;
    while left < right {
        let sum = c[left] + c[right];
        if sum == k {
            total += 1;
            left += 1;
            right -= 1;
        } else if sum < k {
            left += 1;
        } else if sum > k {
            right -= 1;
        }
    }
    return total;
}

fn main() {
    println!("{}", max_operations(vec![1, 2, 3, 4], 4));
}
