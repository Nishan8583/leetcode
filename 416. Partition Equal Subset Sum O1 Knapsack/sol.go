package main

/*
Given an integer array nums, return true if you can partition the array into two subsets such that the sum of the elements in both subsets is equal or false otherwise.



Example 1:

Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].
Example 2:

Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.

Soln
SO basically, we need to sum the array, and loop and check if sum/2 is found by summing any vallues
also if sum is odd, then not possible
[1,5,11,5]
loop through back
set = {0}, so on last 5, sum can either be 0 or 5 set= {0,5} on 11, add both so set= {0,5,11,16} on 5 go through all of them set = {0,5,11,16,5,10,16,21}

this basically, is adding each element to all possible sums
*/

import "fmt"

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	//fmt.Println("Sum", sum)
	if sum%2 == 1 {
		fmt.Println("Not even", sum)
		return false
	}
	target := sum / 2
	//fmt.Println("target", target)
	set := map[int]struct{}{}
	set[0] = struct{}{}
	for i := len(nums) - 1; i >= 0; i-- {
		v := nums[i]
		l := map[int]struct{}{}
		for t, _ := range set {
			//fmt.Println("looping", t)
			s := t + v
			if s == target {
				//fmt.Println("Foun target", s)
				return true
			}
			l[s] = struct{}{}
		}
		for v := range l {
			set[v] = struct{}{}
		}
		//fmt.Println("New set", set)
	}

	if _, ok := set[target]; ok {
		return true
	}

	return false
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}
