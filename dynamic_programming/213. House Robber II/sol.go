package main

/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.



Example 1:

Input: nums = [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.

Example 2:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.

Example 3:

Input: nums = [1,2,3]
Output: 3
*/

/*
Same as robber 1, do not include first and last values in the same time
So a helper function that does the same as robber 1, pass witout 0 and pass witouth len(-1) and see which is max
*/
import "fmt"

func rob(nums []int) int {

	a := helper(nums[1:])
	b := helper(nums[:len(nums)-1])
	c := nums[0]
	//fmt.Println(a, b, c)
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	} else {
		return c
	}
}
func helper(nums []int) int {
	//fmt.Println(nums)
	rob1, rob2 := 0, 0

	for _, v := range nums {
		temp := max(rob1+v, rob2)
		fmt.Println("max", temp)
		rob1 = rob2
		rob2 = temp
	}
	//fmt.Println("Returnhing", rob2)
	return rob2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))

	fmt.Println(rob([]int{2, 3, 2}))
}
