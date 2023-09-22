/*
https://leetcode.com/problems/minimum-size-subarray-sum/description/
Given an array of positive integers nums and a positive integer target, return the minimal length of a
subarray
whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.



Example 1:

Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.

Example 2:

Input: target = 4, nums = [1,4,4]
Output: 1

Example 3:

Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0

sliding window
[2,3,1,2,4,3]
target = 7
total = 0
l = 2
r = 2
total = 2, r++, total = 5, r++ total = 6, r++, total = 7
now increase l,decrease total, to see if array in between satisifies condition
if so, continue increasing l, if not break out and continue with previous loop

*/

func minSubArrayLen(target int, nums []int) int {
	prev_sum := 0
	l := 0
	min_array := len(nums) + 1

	for r, v := range nums {
		prev_sum = prev_sum + v

		if prev_sum >= target {
			if ((r - l) + 1) < min_array {
				min_array = r - l + 1

			}
			for {
				prev_sum = prev_sum - nums[l]
				l++
				if prev_sum >= target {
					if ((r - l) + 1) < min_array {
						min_array = r - l + 1
					}
				} else {
					break
				}
			}
		}
	}

	if min_array > len(nums) {
		return 0
	}
	return min_array
}