/*
Given an integer array nums, handle multiple queries of the following type:

Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
Implement the NumArray class:

NumArray(int[] nums) Initializes the object with the integer array nums.
int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).

Example 1:

Input
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
Output
[null, 1, -1, -3]

Explanation
NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3
*/
package main

type NumArray struct {
	prefixSums []int
}

func Constructor(nums []int) NumArray {
	n := NumArray{
		prefixSums: []int{},
	}

	n.prefixSums = append(n.prefixSums, nums[0])
	c := 1
	for _, v := range nums[1:] {
		n.prefixSums = append(n.prefixSums, n.prefixSums[c-1]+v)
		c++
	}

	return n
}

func (this *NumArray) SumRange(left int, right int) int {
	r := this.prefixSums[right]
	l := 0
	if left > 0 {
		l = this.prefixSums[left-1]
	}
	return r - l
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
