package main

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	curset := []int{}
	subset := [][]int{}
	helper(0, nums, &curset, &subset)
	return subset
}
func helper(i int, nums []int, curset *[]int, subset *[][]int) {
	// if i is greater than length then return
	if i >= len(nums) {
		dst := []int{}
		dst = append(dst, *curset...)
		*subset = append(*subset, dst)
		return
	}

	// When we decide to  include the ith element
	*curset = append(*curset, nums[i])

	helper(i+1, nums, curset, subset)

	// now when we decide do not include the ith element
	*curset = (*curset)[:len(*curset)-1]

	for {

		if (i + 1) >= len(nums) {
			break
		}
		if nums[i] != nums[i+1] {
			break
		}
		i += 1

	}
	helper(i+1, nums, curset, subset)

}
