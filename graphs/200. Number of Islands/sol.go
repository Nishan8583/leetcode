package main

import "fmt"

func numIslands(grid [][]byte) int {
	count := 0
	visited := map[string]struct{}{}

	// loop through each row
	for row := 0; row < len(grid); row++ {
		fmt.Println(len(grid[0]))
		// loop through each columns
		for col := 0; col < len(grid[0]); col++ {
			if explore_island(grid, row, col, visited) {
				count += 1
			}
		}
	}

	return count
}

func explore_island(grid [][]byte, row int, col int, visited map[string]struct{}) bool {

	// check if out of bound row
	if row < 0 || row >= len(grid) {
		return false
	}

	if col < 0 || col > len(grid[0]) {
		return false
	}

	// its water
	if grid[row][col] == '0' {
		return false
	}

	pos := fmt.Sprintf("%d,%d", row, col)
	if _, ok := visited[pos]; ok {
		return false
	}

	visited[pos] = struct{}{}
	explore_island(grid, row-1, col, visited) // explore up
	explore_island(grid, row+1, col, visited) // explore down
	explore_island(grid, row, col+1, visited) // explore right
	explore_island(grid, row, col-1, visited) // explore left

	return true
}

func main() {
	grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	fmt.Println(numIslands(grid))
}

