/*

55. Jump Game
Solved
Medium
Topics
Companies

You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.



Example 1:

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:

Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

Solution:
For [2,3,1,1,4]
1. We start from the second last , our goal is index 4
2. value 1 index 3 can reach 4, because it can jump 1
3. new goal is changed to index 3
4. value 1 index 2 can reach goal 3
5. goal is updated to index 2
6. Repeat till index 0, if final goal is 0, it means we could reach the last index


Alternative solutions:
Start from last, do like decision tree, and cache whether each location could reach the final position or not
*/

package main

func canJump(nums []int) bool {
	goal := len(nums) - 1

	for i := len(nums) - 2; i >= 0; i-- {
		if (i + nums[i]) >= goal {
			goal = i
		}
	}

	if goal != 0 {
		return false
	}

	return true
}
