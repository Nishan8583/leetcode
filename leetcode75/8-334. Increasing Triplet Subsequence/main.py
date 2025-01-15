"""
1. (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]
2. so it does not have to be linear, thanks for rereading the question nishan
3. [1,2,3,4,5]
4. 2 pointers?
5. L = 0, R = 2
6. Check if something within that range?
start = l+1
loop:
get lowest between L,i,and R
lowest :=
for i = start and i < R:
    if i < lowest and
        keep track of lowest
        return True
if not found increase R and try again
7. [2,1,5,0,4,6]
L = 0, R = 2
2,1,5 here is is smaller than 2,so it breaks out, lowest value here is 1
2,1,5,0, it loops till condition here 5 is greater than 1 which is lowest, so it does not work
2,1,5,0,4 no
2,1,5,0,4,6, here 0,4,6 matches condition
but it is still in adjacent
for 4, it needs to check if there is any vlaue less than it on the left, we should keep track of lowest value
1. [1,2,3,4,5]
lowest in i = 0 is 1
in i = 2 is also 1, check if next is greater than current, yes, then return true

5,4,3,2,1
i = 0, lowest 5
i = 1, is it greater than lowest on side? no, update lowest, lowest = 4
i = 3, same, lowest = 3
i = 4, same lowest = 4
i = 5, same lowest = 1, first conditon is never met


[2,1,5,0,4,6]
i = 0, lowest = 2
i = 1, first does not match, lowest = 1
i = 2, fist matches, but i+1 is not greatr than it,no need to update lowest
i = 3, first conditon does not match, update lowest to 0
i = 4, first conditon matches, yes 0 < 4, check next, 4<6 yes, return true



"""


class Solution:
    def increasingTriplet(self, nums: list[int]) -> bool:
        if len(nums) < 3:
            return False

        lowest = [0] * len(nums)
        highest = [0] * len(nums)
        low = nums[0]
        i = 0

        # each index in the lowest array will have the lowest value till that point
        # from left to right
        while i < len(nums):
            v = nums[i]
            if v < low:
                low = v
            lowest[i] = low
            i += 1

        i = len(nums) - 1
        high = nums[len(nums) - 1]
        # Each index in the highest array will have the highest value till that index from right to left
        while i >= 0:
            v = nums[i]
            if v > high:
                high = v
            highest[i] = high
            i -= 1
        # print(f"lowest={lowest},highest={highest}")
        i = 1

        # so for each value in nums, check if corresponing lowest value is lower than that nums[i]
        # and corresponing highest value is greater than nums[i], if so condition is matched
        while i < len(nums):
            v = nums[i]
            if lowest[i] < v and v < highest[i]:
                return True
            i += 1
        return False


s = Solution()
print(s.increasingTriplet([1, 2, 3, 4, 5]))
print(s.increasingTriplet([5, 4, 3, 2, 1]))
print(s.increasingTriplet([2, 1, 5, 0, 4, 6]))
print(s.increasingTriplet([2, 4, -2, -3]))
print(s.increasingTriplet([0, 4, 2, 1, 0, -1, -3]))
print(s.increasingTriplet([20, 100, 10, 12, 5, 13]))
print(s.increasingTriplet([5, 1, 6]))
