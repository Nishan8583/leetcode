package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	count := 0
	visited := map[string]struct{}{}

	maxArea := 0

	// loop through each row
	for row := 0; row < len(grid); row++ {
		// loop through each columns
		for col := 0; col < len(grid[0]); col++ {
			if ok, area := explore_island(grid, row, col, visited, 0); ok {
				count += 1
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

func explore_island(
	grid [][]int,
	row int,
	col int,
	visited map[string]struct{},
	area int,
) (bool, int) {
	// check if out of bound row
	if row < 0 || row >= len(grid) {
		return false, area
	}

	if col < 0 || col >= len(grid[0]) {
		return false, area
	}

	// its water
	if grid[row][col] == 0 {
		return false, area
	}

	pos := fmt.Sprintf("%d,%d", row, col)
	if _, ok := visited[pos]; ok {
		return false, area
	}

	// since its 1 and since its not visited, increase the area
	area += 1

	visited[pos] = struct{}{}
	_, area = explore_island(grid, row-1, col, visited, area) // explore up
	_, area = explore_island(grid, row+1, col, visited, area) // explore down
	_, area = explore_island(grid, row, col+1, visited, area) // explore right
	_, area = explore_island(grid, row, col-1, visited, area) // explore left

	return true, area
}
