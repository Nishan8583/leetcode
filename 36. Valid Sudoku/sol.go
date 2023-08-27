package main

import (
	"fmt"
)

/*
DO the grid traversal
each row and column
check if row is valid
check if valid column, map column number to already checked because we might get to same column twice
check if the current box is valid 3x3
map already visited position, so we do not visit that again, array of 2 int
*/
type empty struct{}

func isValidSudoku(board [][]byte) bool {

	col_visited := map[int]empty{}
	visited := make(map[[2]int]empty) // holds 'row,col' thats already visited

	// loop through each row
	for row := 0; row < len(board); row++ {
		if !check_if_valid_sudoku(board[row]) { // check if row is valid

			return false
		}
		for col := 0; col < len(board[0]); col++ {
			//fmt.Println("checking column", row, col)
			if _, ok := col_visited[col]; !ok {
				if !check_if_col_valid(board, row, col) { // if the column is invalid, return false
					//fmt.Println("column invalid", row, col)
					return false
				}
				col_visited[col] = empty{}
			}
			if !check_if_box_is_valid(board, row, col, visited) {
				//fmt.Println("box invalid", row, col)
				return false
			}

		}
	}
	return true
}

func check_if_valid_sudoku(row []byte) bool {
	seen_numbers := map[byte]empty{}
	for _, v := range row {
		if v == '.' {
			continue
		}
		if _, ok := seen_numbers[v]; ok {
			return false
		}
		seen_numbers[v] = struct{}{}
	}
	return true
}

// go from current row to bottom of the row of the same colum
func check_if_col_valid(board [][]byte, row_index int, col_index int) bool {
	number_of_col := len(board[row_index])
	col_values := map[byte]empty{}
	for i := row_index; i < number_of_col; i++ {
		if _, ok := col_values[board[i][col_index]]; ok { // if the value was already present, then we got a duplicate, return false
			return false
		}
		if board[i][col_index] == '.' {
			continue
		}

		col_values[board[i][col_index]] = empty{}
	}
	return true
}

func check_if_box_is_valid(grid [][]byte, row, col int, visited map[[2]int]empty) bool {
	// check if row and column are within the range of the grid
	// if not, these values were already checked in last iteration
	pos := [2]int{row, col}
	if _, ok := visited[pos]; ok {
		return true
	}
	if row+2 >= len(grid) || col+2 >= len(grid[0]) {
		return true
	}
	checked_values := map[byte]empty{}
	for r := row; r <= (row + 2); r++ {
		for c := col; c <= (col + 2); c++ {
			pos := [2]int{r, c}
			if _, ok := visited[pos]; ok {
				continue
			}
			visited[pos] = empty{}
			//fmt.Println(string(grid[r][c]), pos)
			if _, ok := checked_values[grid[r][c]]; ok {
				return false
			}
			if grid[r][c] == '.' {
				continue
			}

			checked_values[grid[r][c]] = empty{}

		}
	}

	return true
}

func main() {
	v := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	v2 := [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Println(isValidSudoku(v))
	fmt.Println(isValidSudoku(v2))
}
