class Solution(object):
    def rob(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        dp1=[-1]*len(nums)
        dp2=[-1] * len(nums)
        sum1 = self.get_max_sum(len(nums)-1,nums,dp1)
        sum2 = self.get_max_sum(len(nums)-2,nums,dp2)
        return max(sum1,sum2)
    def get_max_sum(self,index,nums,dp):
        if index < 0:
            return 0
        
        # if the max value was already calculated then just return it
        if dp[index] != -1:
            return dp[index]
        
        # 2 Cases
        # Case 1: When we include the next house
        sum1 = nums[index] + self.get_max_sum(index-2,nums,dp)
        
        # Case 2: We skip the next house
        sum2 = nums[index] + self.get_max_sum(index-3,nums,dp)
        dp[index] = max(sum1,sum2)
        return max(sum1,sum2)

def get_max_sum(index,nums):
    if index < 0:
        return 0
    
    sum = nums[index] + get_max_sum(index-2,nums)
    return sum
