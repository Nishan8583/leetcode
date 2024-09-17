// https://leetcode.com/problems/maximum-subarray/description/
// https://www.youtube.com/watch?v=5WZl3MMT0Eg&ab_channel=NeetCode
// just looop through, firs the sum will be the first element, at any
// time when the su, is negative, make it 0, else keep adding,
// if local_sum is greater than max_sum, then use that new local_sum
func maxSubArray(nums []int) int {
	max_sum := nums[0]
	local_sum := 0

	for _, v := range nums {
		if local_sum < 0 {
			local_sum = 0
		}

		local_sum += v
		if local_sum > max_sum {
			max_sum = local_sum
		}
	}
	return max_sum

}