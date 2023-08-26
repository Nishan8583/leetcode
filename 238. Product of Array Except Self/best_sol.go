// first loop through array, in an index of final array put the product of last value * prefix_product
// update prefix_product
// for first 1
// in the end, put the value current value * postfix_product
// We are basically first appending product of all prefix indices in my current position
// Then we are appending product of postfix product value and current value
// https://youtu.be/bNvIQI2wAjk
// https://leetcode.com/problems/product-of-array-except-self/submissions/
// [1,2,3,4], initially finay_result  =[], final_result[0]=1
// Second explaination, so basically when we loop first, we are trying to get product of all the prefixes
// on second loop we are trying to get product of all postfixes
//

func productExceptSelf(nums []int) []int {
	final_result := make([]int, len(nums))
	final_result[0] = 1
	prefix_product := 1
	// first loop to get the product of values before the current index
	for i := 1; i < len(nums); i++ {
		final_result[i] = nums[i-1] * prefix_product
		prefix_product = nums[i-1] * prefix_product
	}

	postfix_prodcut := 1
	for i := len(nums) - 1; i >= 0; i-- {

		final_result[i] = final_result[i] * postfix_prodcut
		postfix_prodcut = postfix_prodcut * nums[i]
	}

	return final_result

}
