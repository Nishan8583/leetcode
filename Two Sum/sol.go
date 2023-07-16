func twoSum(nums []int, target int) []int {
	indices := []int{-1, -1}
	for i, v1 := range nums {
		for j, v2 := range nums {
			if i == j {
				continue
			}
			values := v1 + v2
			if values == target {
				indices[0] = i
				indices[1] = j
				return indices
			}
		}
	}

	return indices
}