/*
130. Surrounded Regions
Solved
Medium
Topics
Companies

You are given an m x n matrix board containing letters 'X' and 'O', capture regions that are surrounded:

    Connect: A cell is connected to adjacent cells horizontally or vertically.
    Region: To form a region connect every 'O' cell.
    Surround: The region is surrounded with 'X' cells if you can connect the region with 'X' cells and none of the region cells are on the edge of the board.

A surrounded region is captured by replacing all 'O's with 'X's in the input matrix board.



Example 1:

Input: board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]

Output: [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]

Solution:

Basically, what the question is saying is that, any cell, connected to the outside (edge row or column) which has 'O'
can not be captured, so in the end it must still be 'O', and any other cell indirectly or directly connected to that cell
will also remain O, any other O will be taken over by O

So first we check if any edge row or column is O. if it is mark it as 'T', chjeck if any adjacent up down right left is also O
and mark them as well.

Now do another loop, convert any O to X

Another loop, conver T to O

*/

package main

import "fmt"

func solve(board [][]byte) {
	rows := len(board)
	cols := len(board[0])

	// holds the coordinates that has temporary values, i.e. '0' will be switchec to 'T' for a while
	tempValues := [][2]int{}

	var capture func(r, c int)

	capture = func(r, c int) {
		// if out of bound or board is not 'O' return
		if (r < 0) || (c < 0) || (r == rows) || (c == cols) || (board[r][c] != 'O') {
			return
		}

		// if within bound and it is 'O', do DFS and try to replace all connected 'O' with 'T'
		board[r][c] = 'T'

		// replace the cordicates in the tempValues
		tempValues = append(tempValues, [2]int{r, c})

		// Now we try to see if any other cell with value 'O' directly connected to this one
		// if it is we want it to be changed to T as well
		capture(r+1, c) // below
		capture(r-1, c) // above
		capture(r, c+1) // right
		capture(r, c-1) // left
	}

	// 1. Mark unsorrounded regions for all edge part
	// Check if the edge row and column as any 'O', then mark it as T
	// and search on any connected cells recursively
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			fmt.Println(r, c, board[r][c], 'O')
			if (board[r][c] == 'O') &&
				((r == 0) || (r == rows-1)) ||
				((c == 0) || (c == cols-1)) {
				capture(r, c)
			}
		}
	}

	// Now since we already converted the onces that are directly and indirectly conected to ocean as T
	// Rest 'O' can be safely converted to X
	// 2. Now capture
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			}
		}
	}

	// 3 Reverse the Temporar values
	for _, coord := range tempValues {
		board[coord[0]][coord[1]] = 'O'
	}
}
