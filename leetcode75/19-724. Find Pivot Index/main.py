class Solution:
    def pivotIndex(self, nums: list[int]) -> int:
        total_sum = 0

        # Get the sum of of all the values
        for i in nums:
            total_sum += i

        prefix_sum = 0
        for i in range(0, len(nums)):
            # postfix sum will be the total sum minus prefix sum, minus the current value
            # Because after deleting those from total, all that will be left is postfix sum
            postfix_sum = total_sum - nums[i] - prefix_sum
            # print(nums, i, prefix_sum, postfix_sum, total_sum)
            if postfix_sum == prefix_sum:
                return i
            prefix_sum += nums[i]
        return -1


s = Solution()
print(s.pivotIndex([1, 7, 3, 6, 5, 6]))
