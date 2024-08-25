use std::collections::HashMap;
impl Solution {
    pub fn contains_duplicate(nums: Vec<i32>) -> bool {
        let mut values: HashMap<i32,i32> = HashMap::new();

        for i in nums.iter() {
            if let Some(v) = values.get(&i) {
                return true;
            }
            values.insert(*i,1);
        }

        return false;
    }
}