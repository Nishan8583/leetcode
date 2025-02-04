use std::collections::HashMap;

pub fn find_difference(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<Vec<i32>> {
    let mut nums1_hashmap: HashMap<i32, i32> = HashMap::new();
    let mut nums2_hashmap: HashMap<i32, i32> = HashMap::new();

    for (&v) in &nums1 {
        nums1_hashmap.insert(v, 1);
    }
    for (&v) in &nums2 {
        nums2_hashmap.insert(v, 1);
    }

    let mut answer0: Vec<i32> = vec![];
    let mut answer2: Vec<i32> = vec![];

    for &key in nums1_hashmap.keys() {
        if !nums2_hashmap.contains_key(&key) {
            answer0.push(key);
        }
    }

    for &key in nums2_hashmap.keys() {
        if !nums1_hashmap.contains_key(&key) {
            answer2.push(key);
        }
    }

    return vec![answer0, answer2];
}

fn main() {
    let result = find_difference(vec![1, 2, 3], vec![2, 4, 6]);
    println!(" {:?} {:?}", result[0], result[1]);
}
