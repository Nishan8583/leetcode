/*
39. Combination Sum
Medium
17.5K
355
Companies

Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the
frequency
of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.



Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.

Example 2:

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]

Example 3:

Input: candidates = [2], target = 1
Output: []


Man im gonna be honest the explaination is gonna be really long
basically,  we do not want the duplicate look at this video bro
https://www.youtube.com/watch?v=GBKI9VSKdGg&ab_channel=NeetCode
*/

package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	final := [][]int{}
	var dfs func(i, total int, cur []int)
	dfs = func(i, total int, cur []int) {
		if total == target {
			dst := make([]int, len(cur))
			copy(dst, cur)
			final = append(final, dst)
			return
		}

		if i >= len(candidates) || total > target {
			return
		}

		// appending the value for now
		cur = append(cur, candidates[i])
		dfs(i, total+candidates[i], cur)

		// backtrack, we have once already used it, so we do not want to use it anymore
		cur = cur[:len(cur)-1]
		// skipping i with +1
		dfs(i+1, total, cur)

	}
	dfs(0, 0, []int{})
	return final
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
