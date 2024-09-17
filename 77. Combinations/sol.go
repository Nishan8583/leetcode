/*
https://leetcode.com/problems/combinations/description/
Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].

You may return the answer in any order.

Example 1:

Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
Explanation: There are 4 choose 2 = 6 total combinations.
Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.

Similar to decision tree
lets say n=4, and k=2
so we need possible combinations of slices, with length 2 ranging from 1 to 4, we could do again like decision tree
but we might also need to go through not having value, so we can do
1			2			3		4
2,3,4		3,4			4

As u can see we, do not calculate the set where the possibility of not having is exclued
and at each leevel we reduce the already calcualted value to avoid duplicates
https://neetcode.io/courses/advanced-algorithms/12
*/
package main

import "fmt"

func combine(n int, k int) [][]int {
	combinations := [][]int{}
	helper(1, []int{}, &combinations, n, k)
	return combinations
}

func helper(i int, currentComb []int, combs *[][]int, n, k int) {
	if len(currentComb) == k {
		dst := make([]int, len(currentComb))
		copy(dst, currentComb)
		fmt.Println("length matched", dst, currentComb)
		*combs = append(*combs, dst)
		return
	}

	if i > n {
		return
	}

	for j := i; j < (n + 1); j++ {
		currentComb = append(currentComb, j)

		// decision to include
		helper(j+1, currentComb, combs, n, k)
		currentComb = currentComb[:len(currentComb)-1]

	}
}

func main() {
	fmt.Println(combine(4, 2))
}
