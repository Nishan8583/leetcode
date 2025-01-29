class Solution:
    def largestAltitude(self, gain: list[int]) -> int:
        max_gain = 0

        prefix_sum = 0

        for i in gain:
            prefix_sum = prefix_sum + i
            max_gain = max(max_gain, prefix_sum)

        return max_gain


s = Solution()
print(s.largestAltitude([-5, 1, 5, 0, -7]))
print(s.largestAltitude([-4, -3, -2, -1, 4, 3, 2]))
