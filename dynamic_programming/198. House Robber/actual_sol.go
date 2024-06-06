package main

/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.



Example 1:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.

Example 2:

Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.


*/
//[1,2,3,1]
/*
on each position i,
we keep on calculating the max u can rob
by checking the max values of wether including i-2 or not
rob1 will hold max until i-2
rob2 will hold max until i-1

[1,2,3,1]
for this,
rob1=0
rob2=0
in position 0, max(rob1+v,rob2) = (0+1,0) = 1, rob1=rob2=0, rob2=1
in position 1, max(rob1+v,rob2) = (0+2,1) = 2, rob1=0, rob2=1, rob2=2,rob1 is still 0 here because no values before 1, 2 is greater so use 2
In position 2, max(rob1+v,rob2) = (1+3,2) = 3, rob1=1, rob2=2, rob2=1, Here since 3+1 is greater, sklipping 2, update, rob2 = 3, rob1 is 2,
In position 3, max(rob1+v,rob2) = (2+1,3) = 1, rob1=2 rob2=3 , 2+1 is three, and 3 is 3

In the end rob2 will have the max value so return that
*/
func rob(nums []int) int {
	rob1, rob2 := 0, 0

	for _, v := range nums {
		temp := max(rob1+v, rob2)
		rob1 = rob2
		rob2 = temp
	}
	return rob2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
