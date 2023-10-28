package main

import "fmt"

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{
		matrix: matrix,
	}

}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	total := 0
	fmt.Println("added", total)
	total += helper(this.matrix, row1, col1, row2, col2)
	return total
}

func helper(matrix [][]int, row, column, max_row, max_col int) int {
	sum := 0
	if row > max_row && column > max_col {
		//fmt.Println("end reached at", row, column, max_row, max_col)
		return 0
	}

	if row > max_row {
		sum = helper(matrix, row, column+1, max_row, max_col)
	} else if column > max_col {
		sum = helper(matrix, row+1, column, max_row, max_col)
	} else {
		fmt.Printf("\nAdding row=%d column=%d value=%d \n", row, column, matrix[row][column])
		sum = matrix[row][column] + helper(matrix, row+1, column, max_row, max_col) + helper(matrix, row, column+1, max_row, max_col)
	}
	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func main() {
	t := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	nm := Constructor(t)
	fmt.Println(nm.SumRegion(2, 1, 4, 3))
}
