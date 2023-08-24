package main

import "fmt"

func main() {
	graph := map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}
	fmt.Println(bfs_traversal(graph, "a"))
	fmt.Println(dfs_traversal(graph, "a"))
}
