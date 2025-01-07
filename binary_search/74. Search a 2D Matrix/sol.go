package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	l := [2]int{0, 0}                        // left most value
	r := [2]int{len(matrix[0]), len(matrix)} // right most value

	for {
		if l[0] >= r[0] || // if left exceeds beyond right break
			l[1] >= r[1] {
			break
		}

		// get mid row and mid low
		mid_row := (l[0] + (r[0] - l[0])) / 2
		mid_col := (l[1] + (r[1] - l[1])) / 2
		fmt.Println("searching in index", mid_row, mid_col)
		val := matrix[mid_row][mid_col]
		if val > target {
			r = [2]int{mid_row - 1, mid_col - 1}
			fmt.Println("Value was greater so changing r", r)
		} else if val < target {
			l = [2]int{mid_row + 1, mid_col + 1}
			fmt.Println("value was smaller so changing l", l)
		} else {
			return true
		}
	}
	return false
}
