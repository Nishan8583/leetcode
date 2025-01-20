"""
Question:
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.



Example 1:

Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

Example 2:

Input: nums = [5], k = 1
Output: 5.00000


Solution:
Intially use 2 pointers of length k, just get the sum of that subarray, and update max sum,
increase left and right, keep getting new sum, and keep updating the max sum

My optimized solution is ,get the first subarray sum, then increase left, right can be same, becasue we are not looping now
i.e. (range (0,right) missed the actual right value because it loops till right - 1,
now decrease previous left vlaue and add the new right value to get the new subarray sum, update max_sum
"""


class Solution:
    def findMaxAverage(self, nums: list[int], k: int) -> float:
        left = 0
        right = k
        if len(nums) < 2:
            return nums[left] / k
        max_sum = 0  # holds the max sum we have encountered
        prev_sum = 0  # holds the sum of previous subarray, we use this so that we do not havae to calculate sum of each subarray,
        # and can instead just get the sum by substracting prevous left value and adding new right value
        #
        # Get the first subarray sum
        for i in range(left, right):
            # ls = nums[i]
            max_sum = max_sum + nums[i]
            # print(
            #    f"First Sum with index {i} and value {nums[i]} is {ls} total sum is {max_sum}"
            # )

        prev_sum = max_sum  # update previous subarray sum
        max_sum = max_sum / k  # update max sum
        left += 1  # update left pointer, no need to update right, because range only loops till right-1

        # print(f"final sum is {sum}")
        while right < len(nums):
            # print("inside while", left, right)
            sum = (
                prev_sum - nums[left - 1]
            )  # delete previous left value, because it no longer is within the subarray
            # print(
            #    f"previous sum was {prev_sum}, now we are removing {nums[left-1]}, so our new sum becomses {sum}"
            # )
            # print(f"SUM====={sum}")
            # sum = prev_sum + nums[left]
            sum = sum + nums[right]  # add the new right value
            # print(f"adding {nums[right]} new sum becomses {sum} ")
            prev_sum = sum  # update previous subarray sum
            sum = sum / k
            # sum = 0
            # print(f"final sum is {sum}")
            if sum > max_sum:
                # print(f"curent sum {sum} max_sum {max_sum}")
                max_sum = sum
            left += 1
            right += 1

        return max_sum


s = Solution()
print(s.findMaxAverage([1, 12, -5, -6, 50, 3], 4))
print(s.findMaxAverage([5], 1))
print(s.findMaxAverage([3, 3, 4, 3, 0], 3))
print(s.findMaxAverage([0, 1, 1, 3, 3], 4))
