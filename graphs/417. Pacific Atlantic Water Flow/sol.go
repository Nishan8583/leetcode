/*

There is an m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean. The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges.

The island is partitioned into a grid of square cells. You are given an m x n integer matrix heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).

The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west if the neighboring cell's height is less than or equal to the current cell's height. Water can flow from any cell adjacent to an ocean into the ocean.

Return a 2D list of grid coordinates result where result[i] = [ri, ci] denotes that rain water can flow from cell (ri, ci) to both the Pacific and Atlantic oceans.


solution:

So what the question is saying to us is that, we have a grid, above and left is pacific ocean, right and below is atlantic ocean, so we need to return all position that can reach both atlantic and pacific,
each value in the cell represents the height of that cell, and water can flow from higher to lower.

So we first start at first row and bottom row, because they both can reach pacific and atlantic ocean respectively.
Similarly for left and right colum.
Then we do dfs, since we are going the opposite direction, we need to find value that is higher than previous.
This code is just copy and paste from neetcode, because this solution was a bit too hard for me lmfao

*/

package main

func pacificAtlantic(heights [][]int) [][]int {
	ROWS, COLS := len(heights), len(heights[0])
	pac, atl := make(
		map[int]bool,
	), make(
		map[int]bool,
	) // holds the visited coordinates for pacific and atlantic ocean respectively.

	var dfs func(int, int, map[int]bool, int)
	dfs = func(r, c int, visit map[int]bool, prevHeight int) {
		// if row and columns are out of bounds or current height is less than previous height (because we are doing this in reverse it should be greater), return
		if visit[r*COLS+c] ||
			r < 0 ||
			c < 0 ||
			r == ROWS ||
			c == COLS ||
			heights[r][c] < prevHeight {
			return
		}

		// marked as visited
		visit[r*COLS+c] = true

		// vist east west north south
		dfs(r+1, c, visit, heights[r][c])
		dfs(r-1, c, visit, heights[r][c])
		dfs(r, c+1, visit, heights[r][c])
		dfs(r, c-1, visit, heights[r][c])
	}

	// vist first row and all columns, last row and all columns
	for c := 0; c < COLS; c++ {
		dfs(0, c, pac, heights[0][c])
		dfs(ROWS-1, c, atl, heights[ROWS-1][c])
	}

	// viist last row first column and last row last column
	for r := 0; r < ROWS; r++ {
		dfs(r, 0, pac, heights[r][0])
		dfs(r, COLS-1, atl, heights[r][COLS-1])
	}

	res := make([][]int, 0)

	// if a cordinate is in both visited set for pacific and atlantic, it means it can reach both ocean
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if pac[r*COLS+c] && atl[r*COLS+c] {
				res = append(res, []int{r, c})
			}
		}
	}
	return res
}
