"""
1. Sliding window.
2. L= 0, R = 1
3. If [0,1,1,1,0,1,1,0,1] n[0] = 0, n[R]=1
4. If n[l] == 0, while n[l] != 0, l++, l = 1
5. keep on increasing till r = 4
6. r = 4, n[r] is 0, then k = 0,
7. increase r, till r = 6,
8, n[r] = 0 when r= 7, can no longer ignore, so
9. Calculate the number of count, if k == 0, then r-l-1 cause we can not count 0 as well
10. while k != 1: l++,
9.
"""


class Solution:
    def longestSubarray(self, nums: list[int]) -> int:
        left = 0
        k = 1
        max_length = 0
        # we have to remove one element anyways
        if len(nums) == 0:
            return 0

        for i in range(0, len(nums)):
            # in this condition we are in a situation where we can not continue, we will need to increase left
            if nums[left] == 0 and nums[i] == 0 and k == 0:
                left += 1
                continue

            # ignore the first 0 we encountered and now set k to 0
            if nums[i] == 0 and k == 1:
                k = 0
                continue

            if nums[i] == 1:
                continue

            # if current element now is 0, but k is not 1, then we have already ignore previous 0, so we can not ignore it anymore
            max_length = max(max_length, i - left - 1)

            enc_first_zero = 0
            while enc_first_zero != 1:
                left += 1
                if (
                    nums[left] == 0
                ):  # we encountered the i where first 0 was ignored, so we will increase by 1 more, and then reach
                    left += 1
                    enc_first_zero = 1
                    break

        if k == 0:
            max_length = max((len(nums) - 1) - left, max_length)
        else:
            max_length = max((len(nums) - 1) - left, max_length)

        if max_length == len(nums):
            return max_length - 1
        else:
            return max_length


s = Solution()
print(s.longestSubarray([1, 1, 0, 1]))
print(s.longestSubarray([0, 1, 1, 1, 0, 1, 1, 0, 1]))
print(s.longestSubarray([1, 1, 1]))
