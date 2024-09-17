// [-2,1,-3,4,-1,2,1,-5,4]
// My inital logic was, loop through the array
// if the current value is greater than the previous sum, and sum of current plus previous is greater than max_sum
// replace the max_sum with the current one
// damn it had some issues
func maxSubArray(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	local_sum := 0
	max_sum := 0

	first := true
	for i, v := range nums {
		if first {
			first = false
			local_sum = v
			max_sum = v
			continue
		}

		local_sum = local_sum + v

		fmt.Printf("local sum %d and max_sum %d at index %d value %d\n", local_sum, max_sum, i, v)
		if v > max_sum {
			fmt.Println("Found a value thats bigger than the sum", v)
			if local_sum > v {
				max_sum = local_sum
				continue
			}

			fmt.Println("NOW BOTH LOCAL AND mAX SUKM WILL BE", v)
			max_sum = v
			local_sum = v
			continue
		} else if local_sum > max_sum {
			fmt.Println("local_sum is greater than the max_sum", local_sum, max_sum)
			max_sum = local_sum
		}

	}

	return max_sum

}