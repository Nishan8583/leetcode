"""
Sliding window
[1,1,1,0,0,0,1,1,1,1,0], k = 2
[1,1,1,0], since last one is 0, flip it, [1,1,1,0], zerocount is 1
[1,1,1,1,0] flip it [1,1,1,0,0] now zerocount is 2
check right [1,1,1,0,0,0], flip it, now zero count is 3,
ok so zerocount is greater than k
, now we decrease zerocount only if we find where left is 0, i.e. the first 1 was flipped
0, no, 1, no, 2, no, num[3] == 0, ok, so we now decrease zerocount, no zerocount < k
get count = r - l +1, check if greater than max,
now our new window will be from num[3], the place where we first set to 1,

"""


class Solution:
    def longestOnes(self, nums: list[int], k: int) -> int:
        left, right, zerocount = 0, 0, 0
        max_count = 0
        for right in range(len(nums)):
            if nums[right] == 0:
                zerocount += 1
            if zerocount > k:
                while zerocount > k:
                    if nums[left] == 0:  # if this number was flipped
                        zerocount -= 1
                    left += 1
            max_count = max(max_count, right - left + 1)

        return max_count


s = Solution()
print(s.longestOnes([1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0], 2))
print(s.longestOnes([0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1], 3))
""" 
[0,0,1,1,0,0] r = 5 ,zerocount = 4 l = 0
frst nums[left] == 0, so zerocount - 1
left += 1
count = r-l, 5-1+1 = 5 

again flip, 4= 6, zerocount = 4
nums[1] = 0, so l++,zerocount -1 = 3 
r-l+1 = 6-1+1 = 6

"""
