/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

I solved it in pretty simple way, just sort the slice
check if prev element and current element differnce is gte > 1, if it is check if local counter is greater than max counter and update accordingly

*/
import (
	"sort"
	"fmt"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)

	max_c := 0
	checked := map[int]struct{}{
		nums[0]: {},
	}
	c := 1
	for i := 1; i < len(nums); i++ {
		if _, ok := checked[nums[i]]; ok {
			continue
		}
		checked[nums[i]] = struct{}{}

		if (nums[i] - nums[i-1]) > 1 {
			if c > max_c {

				max_c = c
				c = 1
			} else {
				c = 1
			}
			continue
		}
		c += 1

	}

	if c > max_c {
		max_c = c
	}
	return max_c
}