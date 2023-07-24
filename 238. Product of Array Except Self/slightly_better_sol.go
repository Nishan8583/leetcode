// Basically create two arrays
// one is prefix array, that holds product of all the numbers before it
// one is postfix array, that holds product of all numbers after it
// Then loop through the nums array, and get product of before and after index in prefix and postfix array respectively
// EX; nums = [1,2,3,4]
// prefix becomes = [1,2,6,24] , (current value in nums) * (prev_value in prefix arrau)
// first 1 has no prefix, so its one, then 2*1 = 2, then 3*2 = 6, then  4*6 = 2
// postfix becomes [24,24,12,4]
// in reverse, (current value in nums i.e. 4) * (prev value in potfix)
// 4 * 1 = 4m 3*4 = 12, 12*2 = 24 and 24 * 1 =1
// Finally loop through sums
// [24,12,8,6]
// for 1, no prefix, postfix has 24, so 24*1 = 24
// for 2, prefix is 1, postfix is 12,so 12
// for 3, prefix is 2, postfix is 4, so 8
// for 4, prefix is 6 and postfix is none, so 1 = 6
func productExceptSelf(nums []int) []int {
	final_result := make([]int, len(nums))

	prefix_product_arrays := make([]int, len(nums))
	postfix_product_arrays := make([]int, len(nums))

	prev_prefix := 1
	for i, v := range nums {
		prev_prefix = prev_prefix * v
		prefix_product_arrays[i] = prev_prefix
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		postfix = postfix * nums[i]
		postfix_product_arrays[i] = postfix
	}

	fmt.Println(prefix_product_arrays, postfix_product_arrays)
	for i, _ := range nums {
		prefix_product_position := i - 1
		postfix_product_position := i + 1
		prefix_product_value := 1
		postfix_product_value := 1
		if prefix_product_position >= 0 {
			prefix_product_value = prefix_product_arrays[prefix_product_position]
		}
		if postfix_product_position >= len(nums) {
			postfix_product_value = 1
		} else {
			postfix_product_value = postfix_product_arrays[postfix_product_position]
		}
		product := prefix_product_value * postfix_product_value
		final_result[i] = product
	}
	return final_result

}