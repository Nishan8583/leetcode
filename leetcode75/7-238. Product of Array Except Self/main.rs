pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
    let mut product_array = vec![1; nums.len()];
    let mut prefix_product = 1;

    for (i, &v) in nums.iter().enumerate() {
        if i == 0 {
            continue;
        }
        product_array[i] = prefix_product * nums[i - 1];
        prefix_product = product_array[i];
    }

    let mut postfix_product = 1;

    for (i, v) in nums.iter().enumerate().rev() {
        product_array[i] = postfix_product * product_array[i];
        postfix_product = postfix_product * nums[i];
    }

    return product_array;
}

fn main() {
    println!("{:?}", product_except_self(vec![1, 2, 3, 4]));
}
