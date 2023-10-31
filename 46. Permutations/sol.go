package main

import (
	"fmt"
	"slices"
)

func permute(nums []int) [][]int {
	return helper(0, nums)
}

func helper(n int, nums []int) [][]int {
	if n >= len(nums) {
		return [][]int{}
	}

	resPerms := [][]int{}
	perms := helper(n+1, nums)
	if len(perms) == 0 {
		return [][]int{
			{nums[len(nums)-1]},
		}
	}
	fmt.Println("received permutation", perms)
	// we are going to loop through the all the permutations and
	// put the current value in each position
	for _, p := range perms {
		fmt.Println("loop")
		for i := 0; i < (len(p) + 1); i++ {
			dst := make([]int, len(perms))
			copy(dst, p)
			lPerm := slices.Insert(dst, i, nums[n])
			resPerms = append(resPerms, lPerm)
		}
	}

	return resPerms
}

func main() {
	fmt.Println(permute([]int{5, 4, 6, 2}))
}
