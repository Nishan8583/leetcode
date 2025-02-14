"""
1. Find the greatest number in the array
2. For  each integer in array, add the extra candies,  check if it is greater than or equals to the greatest.
3. Set values accordingly
"""


class Solution:
    def kidsWithCandies(self, candies: list[int], extraCandies: int) -> list[bool]:
        greatest = max(candies)

        result = [False] * len(candies)

        for index, value in enumerate(candies):
            if value + extraCandies >= greatest:
                result[index] = True
        return result
