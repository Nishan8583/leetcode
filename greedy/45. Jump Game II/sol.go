/*
45. Jump Game II
Medium
Topics
Companies

You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].

Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:

    0 <= j <= nums[i] and
    i + j < n

Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].

Solution:
Decision tree with cache works, O(n2)

Here we use greeedy,
So first [2,3,1,1,4]
count = 0, l =0, r = 0
Max is 2, least jump is 1, so our imagination is at [3 ,1], l = 1,r  = 2, count++ = 2

Again, l = r+1=3
for r for [3,1], max jump is 3, which is r = index + value = 1+3 = 4, so r = 4
count++ = 2
so new imagination is [3,4], we reached the final position
*/

package main

import "fmt"

func jump(nums []int) int {
	// count is the number of step it took to reach the final position
	count := 0

	l, r := 0, 0

	for r < (len(nums) - 1) {
		farthest := 0

		// find the max jump we can do, for [2,[3,1],1,4] when l points to 3 and r to 1
		// 3 can jump the farthest, i.e. index=1+value=3 = 4
		for i := l; i < (r + 1); i++ {
			cur_jump := i + nums[i]
			if cur_jump > farthest {
				farthest = cur_jump
			}
		}
		l = r + 1
		r = farthest
		count++
	}
	return count
}

func main() {
	type tests struct {
		values   []int
		expected int
	}

	testCases := []tests{
		{values: []int{2, 3, 1, 1, 4}, expected: 2},
		{values: []int{2, 3, 0, 1, 4}, expected: 2},
	}

	for _, test := range testCases {
		v := jump(test.values)
		if v != test.expected {
			fmt.Printf("failure, expected %d got %d for %v\n", test.expected, v, test)
		} else {
			fmt.Printf("success for %v\n", test)
		}
	}
}
