/*
Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.
https://leetcode.com/problems/contains-duplicate-ii/description/
*/
package main

func containsNearbyDuplicate(nums []int, k int) bool {
	i := 0
	j := 1
	win_size := 1
	for {
		if i >= len(nums) && j >= len(nums) {
			break
		}
		if j >= len(nums) {

			j = i + 2
			i++
			win_size = 1
			continue
		}

		if nums[i] == nums[j] && win_size <= k {
			return true
		}
		if win_size > k {

			j = i + 2
			i++
			win_size = 1
			continue
		}
		j++
		win_size++
	}
	return false
}
