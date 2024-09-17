/*
Given an integer array nums of unique elements, return all possible
subsets
(the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

Example 1:

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:

Input: nums = [0]
Output: [[],[0]]

Basically use decision tree
[1,2,3]
First decision to include 1, then not include 1,
and on top of that decision, incude 2, not innclude 2
and so on, and return final array.
https://neetcode.io/courses/advanced-algorithms/11
*/
package main

import "fmt"

func subsets(nums []int) [][]int {
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
		fmt.Println("Appending", dst, *curset)
		*subset = append(*subset, dst)
		return
	}

	// When we decide to  include the ith element
	*curset = append(*curset, nums[i])
	helper(i+1, nums, curset, subset)

	// now when we decide do not include the ith element
	*curset = (*curset)[:len(*curset)-1]

	helper(i+1, nums, curset, subset)

}
func main() {
	values := []int{3, 2, 4, 1}
	fmt.Println(subsets(values))
}
