// https://leetcode.com/problems/maximum-product-subarray/
// https://youtu.be/lXVy6YWFcRM
// get max value first, initially current min and max is 1,
// if 0 both is equals to 1, hold the most minimum and maximim and respective variable
// then out of max and result, update result with the highest
package main

func maxProduct(nums []int) int {
	result := max(nums...)
	for _, v := range nums {
		if v > result {
			result = v
		}
	}
	curr_min := 1
	curr_max := 1

	for _, n := range nums {
		if n == 0 {
			curr_min = 1
			curr_max = 1
			continue
		}

		tmp := curr_max * n
		curr_max = max(tmp, curr_min*n, n)
		curr_min = min(tmp, curr_min*n, n)
		if curr_max > result {
			result = curr_max
		}
	}

	return result
}

func max(values ...int) int {
	max_value := values[0]
	for _, v := range values {
		if v > max_value {
			max_value = v
		}
	}
	return max_value
}

func min(values ...int) int {
	min_value := values[0]
	for _, v := range values {
		if v < min_value {
			min_value = v
		}
	}
	return min_value
}
