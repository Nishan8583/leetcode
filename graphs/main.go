package main

import "fmt"

func main() {

	edges := [][]string{
		{"w", "x"},
		{"x", "y"},
		{"z", "y"},
		{"z", "v"},
		{"w", "v"},
	}
	al := convert_edge_to_list(edges)
	fmt.Println(shortest_path(al, "w", "z"))
}
