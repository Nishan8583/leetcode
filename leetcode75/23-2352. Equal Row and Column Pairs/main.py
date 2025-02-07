from collections import Counter


class Solution:
    def equalPairs(self, grid: list[list[int]]) -> int:
        row_map = Counter()

        # Put the entire row in a hashmap set so that we can look it up again
        # Counter will hold a key and its Counter
        # Key is going to be the row in tuple form
        for i in grid:
            row_map[tuple(i)] += 1

        n = len(grid)
        total = 0

        # Loop through each row, and get the column for the entire column
        for i in range(0, n):
            curr = []
            for j in range(
                0, n
            ):  # we keep on increasing the column number, and keep the row number same, we get the entire column
                curr.append(grid[j][i])
            # if we have seen the tuple before, it means we can encountered it
            # so it can be formed, increase it
            # We use this because, if the column can be matched with one row, then it can be matched with all the rows that are the same, so we want all of the count of all the rows that are the same

            total += row_map[tuple(curr)]
            return total


s = Solution()
# print(s.equalPairs([[3, 2, 1], [1, 7, 6], [2, 7, 7]]))
print(
    s.equalPairs(
        [
            [3, 3, 3, 6, 18, 3, 3, 3, 3, 3],
            [3, 3, 3, 3, 1, 3, 3, 3, 3, 3],
            [3, 3, 3, 3, 1, 3, 3, 3, 3, 3],
            [3, 3, 3, 3, 1, 3, 3, 3, 3, 3],
            [1, 1, 1, 11, 19, 1, 1, 1, 1, 1],
            [3, 3, 3, 18, 19, 3, 3, 3, 3, 3],
            [3, 3, 3, 3, 1, 3, 3, 3, 3, 3],
            [3, 3, 3, 3, 1, 3, 3, 3, 3, 3],
            [3, 3, 3, 1, 6, 3, 3, 3, 3, 3],
            [3, 3, 3, 3, 1, 3, 3, 3, 3, 3],
        ]
    )
)
# print(s.equalPairs([[3, 1, 2, 2], [1, 4, 4, 5], [2, 4, 2, 2], [2, 4, 2, 2]]))
