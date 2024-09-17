func productExceptSelf(nums []int) []int {
	final_result := make([]int, len(nums))

	for i, _ := range nums {
		//value,ok := get_product(nums[i+1:],nums[:i])
		if len(nums[i+1:]) == 0 && len(nums[:i]) == 0 {
			break
		}
		final_product := 1
		for _, v := range nums[i+1:] {
			final_product = final_product * v
		}

		for _, v := range nums[:i] {
			final_product = final_product * v
		}

		final_result[i] = final_product
	}

	return final_result

}

func get_product(first_array, second_array []int) (int, bool) {
	if len(first_array) == 0 && len(second_array) == 0 {
		return -1, false
	}
	final_product := 1
	for _, v := range first_array {
		final_product = final_product * v
	}

	for _, v := range second_array {
		final_product = final_product * v
	}

	return final_product, true
}