/*
994. Rotting Oranges
Solved
Medium
Topics
Companies

You are given an m x n grid where each cell can have one of three values:

    0 representing an empty cell,
    1 representing a fresh orange, or
    2 representing a rotten orange.

Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.

solution:

First I thought Just do a simple DFS and it would work, but as is with most things in life, I was wrong.
Put the rotten oranges in a queue, so there goes one loop already.

Next loop, pop each oranges one by one, from the queue, and then infect 4 directionally, and then
pop these new in the queue as well, there is a inside loop that does this, and if value is not 1, it was already
infected, so continue
*/

package main

import "fmt"

type queue [][2]int

func (q *queue) push(arr [2]int) {
	*q = append(*q, arr)
}

func (q *queue) pop() [2]int {
	if (len(*q)) == 0 {
		return [2]int{}
	}

	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func orangesRotting(grid [][]int) int {
	q := queue{}

	// time to get all oranges rotten and number of fresh oranges count
	time, fresh := 0, 0

	row, cols := len(grid), len(grid[0])

	// first setup, get the number of fresh oranges and rotten oranges
	for r := 0; r < row; r++ {
		for c := 0; c < cols; c++ {

			// this is a fresh oranges
			if grid[r][c] == 1 {
				fresh++
			}
			if grid[r][c] == 2 {
				q.push([2]int{r, c})
			}
		}
	}

	directions := [][]int{
		{0, 1},  // right
		{0, -1}, // left
		{1, 0},  // down
		{-1, 0}, // up
	}
	fmt.Println(len(q))
	for len(q) != 0 && fresh > 0 {
		length := len(q)

		// Pop rotten oranges one by one and see how it infects others
		for i := 0; i < length; i++ {
			rC := q.pop()

			// infect in all directions diiagonally
			for _, v := range directions {
				dir := [2]int{
					rC[0] + v[0], rC[1] + v[1],
				}

				row := dir[0]
				col := dir[1]

				if (row < 0 || row == len(grid)) ||
					(col < 0 || col == len(grid[0])) {
					continue
				}

				if grid[row][col] != 1 {
					continue
				}

				fmt.Printf("%d %d got infected previous value=%d\n", row, col, grid[row][col])
				grid[row][col] = 2
				q.push([2]int{row, col})
				fresh -= 1
			}
		}
		time += 1
	}
	if fresh == 0 {
		return time
	}
	return -1
}

func main() {
	v := orangesRotting([][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	})
	fmt.Println(v)
}
