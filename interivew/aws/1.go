package main

import (
	"fmt"
	"sort"
)

func solve(array []int, steps int) bool {
	new_arr := make([]int, 0, len(array))
	for _, v := range array {
		new_arr = append(new_arr, v)
	}

	sort.Ints(new_arr)
}

func main() {
	input := []int{4, 5, 1, 3, 2}
	fmt.Println(solve(input, 4))
}
