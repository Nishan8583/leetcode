class Solution:
    def moveZeroes(self, nums: list[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        l = 0
        r = 1
        if len(nums) < 2:
            return
        while l < len(nums) and r < len(nums):
            if nums[l] == 0:
                if nums[r] == 0:
                    # print(f"nums[r] is 0 for r={r}")
                    r += 1
                    continue
                # print(f"Found nums[l]=0 for l={l} and r = {r}")
                temp = nums[l]
                nums[l] = nums[r]
                nums[r] = temp
                l += 1
                # print(f"new l is {l}")
                # print(nums)
            else:
                l += 1
                if r + 1 < len(nums):
                    r += 1
        print(nums)


s = Solution()
print(s.moveZeroes([1, 0, 1]))
# s.moveZeroes([1, 0, 1, 0])
# s.moveZeroes([0, 1, 0, 3, 12])

# s.moveZeroes([1, 0, 3, 12])


# s.moveZeroes([0, 1, 3, 0, 12])
