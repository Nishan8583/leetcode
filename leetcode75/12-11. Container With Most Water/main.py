class Solution:
    def maxArea(self, height: list[int]) -> int:
        left = 0
        right = len(height) - 1
        max_area = 0

        while left <= right:
            min_height = min(height[left], height[right])
            width = right - left
            area = width * min_height
            max_area = max(area, max_area)

            if height[left] < height[right]:
                left += 1
            else:
                right -= 1
        return max_area


s = Solution()
print(s.maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]))
print(s.maxArea([1, 1]))
