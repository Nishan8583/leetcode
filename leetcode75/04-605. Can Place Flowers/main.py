"""
1. Loop, if current is 0, check if next is 0, and next after that is also 0 if so decrease n
else continue
2. If current is 1,next mut be 0 and third must be 0
"""


class Solution:
    def canPlaceFlowers(self, flowerbed: list[int], n: int) -> bool:
        count = 0
        for i in range(len(flowerbed)):
            if flowerbed[i] == 1:
                continue

            empty_left_plot = (i == 0) or (flowerbed[i - 1] == 0)
            empty_right_plot = (i == len(flowerbed) - 1) or (flowerbed[i + 1] == 0)

            if empty_left_plot and empty_right_plot:
                flowerbed[i] = 1
                count += 1

        return count >= n
