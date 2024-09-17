/*
300. Longest Increasing Subsequence
Solved
Medium
Topics
Companies

Given an integer array nums, return the length of the longest strictly increasing
subsequence
.



Example 1:

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

Example 2:

Input: nums = [0,1,0,3,2,3]
Output: 4

Example 3:

Input: nums = [7,7,7,7,7,7,7]
Output: 1



Constraints:

    1 <= nums.length <= 2500
    -104 <= nums[i] <= 104



Follow up: Can you come up with an algorithm that runs in O(n log(n)) time complexity?


Solution:
You see I actually figured this one out on my own. THe only thing i missed is that we would actually have to loop through the array once again to see if the next element was greater than the current element.

So we start at the back, calculate the longest increasing subsequence from the back, for the last its always going to be 1, we cache it.
Then we go to second last, check if next element is > then current, if so get the LIS for next elelemtn from cahce, but we also have to check the element one after that and so on, we we have to loop till the end.
Cache it and so on

*/

package main

import "fmt"

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	max := 1
	for i := len(nums) - 1; i >= 0; i-- {
		count := 1 // current position will always have 1

		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				if dp[j] != 0 {
					//fmt.Println("Found cahce for", nums[j])
					count = findMax(count, dp[j]+1)
				}
			}
		}
		dp[i] = count
		if count > max {
			max = count
		}
	}
	//fmt.Printf("%+v\n", dp)
	return max
}

func findMax(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))           // 4
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}
