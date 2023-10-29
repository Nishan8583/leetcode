package main

import "fmt"

type NumMatrix struct {
	sumMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	// create a sum matrix, that has extra padding to account for the out of bound issue
	rows := len(matrix) + 1
	cols := len(matrix[0]) + 1
	sumMatrx := make([][]int, rows)
	for i := range sumMatrx {
		sumMatrx[i] = make([]int, cols, cols)
	}
	// add 0 in the first row and col

	for row_index, rows := range matrix {
		row_sum := 0

		for col_index, col_val := range rows {

			row_sum = row_sum + col_val
			fmt.Printf("pos=%d,%d matrix value=%d pos=%d,%d matrix_value=%d\n", row_index, col_index, matrix[row_index][col_index], row_index, col_index, sumMatrx[row_index][col_index])
			sum := row_sum + sumMatrx[row_index][col_index+1] // col +1 gives one more value to the right, because first values are always zero
			sumMatrx[row_index+1][col_index+1] = sum
		}
	}

	fmt.Println("[")
	for _, row := range sumMatrx {
		fmt.Printf("	%v\n", row)
	}
	fmt.Println("[")

	return NumMatrix{
		sumMatrix: sumMatrx,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	fmt.Println(row1, row2, col1, col2)

	// increase row and column sicne we are gonna get from sum matrix, corresponding value will be in one higher position
	row1, row2, col1, col2 = row1+1, row2+1, col1+1, col2+1

	fmt.Println(row1, row2, col1, col2)
	bottom := this.sumMatrix[row2][col2]
	above := this.sumMatrix[row1-1][col2]
	left := this.sumMatrix[row2][col1-1]
	topLeft := this.sumMatrix[row1-1][col1-1] // get the whole box sum

	return bottom - above - left + topLeft
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
