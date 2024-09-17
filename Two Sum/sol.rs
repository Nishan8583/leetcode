impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        for (pos1,i) in nums.iter().enumerate() {
            println!("{} {}",pos1,i);
            for (pos2,j) in nums.iter().enumerate() {
                if pos1 == pos2 {
                    continue;
                }
                println!("{} {}",pos2,j);
                let sum = i+j;
                if (sum == target) {
                    return vec![pos1 as i32, pos2 as i32];
                }
            }
        }
        return vec![];
    }
}