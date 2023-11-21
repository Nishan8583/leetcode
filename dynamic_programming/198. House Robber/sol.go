func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) <= 2 {
		if nums[0] < nums[1] {
			return nums[1]
		} else {
			return nums[0]

		}
	}
	firstSum := getSum(nums[0], nums[2:])
	secondSum := getSum(nums[1], nums[3:])

	if firstSum > secondSum {
		return firstSum
	}

	return secondSum
}

func getSum(prevSum int, values []int) int {
	if len(values) < 2 {
		if len(values) == 1 {
			return prevSum + values[0]
		}
		return prevSum
	}

	sum := prevSum + values[0]
	firstSum := getSum(sum, values[2:])

	sum = sum - values[0] + values[1]
	secondSum := sum
	if len(values) >= 4 {
		secondSum = getSum(sum, values[3:])
	}

	if secondSum < firstSum {
		return firstSum
	}

	return secondSum
}