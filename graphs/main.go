package main

import "fmt"

func main() {

	grid := [][]string{
		{"W", "L", "W", "W", "W"},
		{"W", "L", "W", "L", "W"},
		{"W", "W", "W", "L", "W"},
		{"L", "W", "L", "L", "W"},
		{"L", "L", "W", "W", "W"},
	}
	fmt.Println(minimum_island(grid))

}
