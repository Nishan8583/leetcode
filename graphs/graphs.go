package main

import (
	"fmt"
)

type stack[T any] struct {
	values []T
}

func (s *stack[T]) push(value T) {
	s.values = append(s.values, value)
}

func (s *stack[T]) pop() T {
	poped := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return poped
}

// DSF use stack
func dfs_traversal(graph map[string][]string, firstElem string) []string {

	result := []string{}
	s := stack[string]{
		values: []string{firstElem},
	}
	for {
		if len(s.values) == 0 {
			break
		}
		value := s.pop()
		result = append(result, value)
		if len(graph[value]) == 0 {
			continue
		}
		for _, v := range graph[value] {
			s.push(v)
		}

	}
	return result
}

type queue[T any] struct {
	values []T
}

func (q *queue[T]) enqueue(value T) {
	q.values = append(q.values, value)
}

func (q *queue[T]) dequeue() T {
	var result T
	if len(q.values) == 0 {
		return result
	} else if len(q.values) == 1 {
		val := q.values[0]
		q.values = []T{}
		return val
	}
	que := q.values[0]
	q.values = q.values[1:]

	return que
}

// BFS = queue
func bfs_traversal(graph map[string][]string, initial string) []string {
	result := []string{}
	q := queue[string]{
		values: []string{initial},
	}

	for {
		if len(q.values) == 0 {
			break
		}
		val := q.dequeue()
		result = append(result, val)
		values := graph[val]
		for _, v := range values {
			q.enqueue(v)
		}
	}

	return result
}

// any traversal, check if we reach dst
func has_path(graph map[string][]string, src, dest string) bool {
	s := stack[string]{
		values: []string{src},
	}
	for {
		if len(s.values) == 0 {
			break
		}
		value := s.pop()
		if value == dest {
			return true
		}
		if len(graph[value]) == 0 {
			continue
		}
		for _, v := range graph[value] {
			s.push(v)
		}

	}
	return false
}

// Traverse, use DFS, mark as visited, check if visiited, out of loop increase connected node couont
func get_component_count(graph map[string][]string) uint {
	type empty struct{}
	var count uint = 0
	visited := map[string]empty{}

	// loop for each nodes in graph
	for node, connections := range graph {

		// if the node has already been visited, then it was already part of a component count, we do not want to count it again.
		if _, ok := visited[node]; ok {
			continue
		}

		count += 1
		visited[node] = empty{}    // adding it to visited node,we do not want to visit it again
		if len(connections) == 0 { // if there is no connections its alone, not connected to any node, just like me lmao !!!
			continue
		}

		// we are using a stack, which means LIFO, i.e. DFS
		s := stack[string]{
			values: []string{connections[0]},
		}

		for {
			if len(s.values) == 0 {
				fmt.Println("breaking out")
				break
			}

			value := s.pop()
			fmt.Println(s.values)
			fmt.Println("popping ", value)
			fmt.Println(s.values)
			if _, ok := visited[value]; ok {
				continue
			}
			visited[value] = empty{}
			if len(graph[value]) == 0 {
				continue
			}
			for _, v := range graph[value] {
				s.push(v)
			}
		}

	}

	return count
}

// get te connected nodes, which has maximum number of nodes
// DO a dfs, check if node already visited, if yes no need to go further
// mark node as visited, increase count, if higher than maximum increase the count
func get_component_max_count(graph map[string][]string) uint {
	type empty struct{}
	var count uint = 0
	visited := map[string]empty{}
	var max_count uint = 0
	// loop for each nodes in graph
	for node, connections := range graph {

		// if the node has already been visited, then it was already part of a component count, we do not want to count it again.
		if _, ok := visited[node]; ok {
			continue
		}

		count += 1
		visited[node] = empty{}    // adding it to visited node,we do not want to visit it again
		if len(connections) == 0 { // if there is no connections its alone, not connected to any node, just like me lmao !!!
			continue
		}

		// we are using a stack, which means LIFO, i.e. DFS
		s := stack[string]{
			values: []string{connections[0]},
		}

		var local_count uint = 1
		for {
			if len(s.values) == 0 {
				fmt.Println("breaking out")
				break
			}

			value := s.pop()
			if _, ok := visited[value]; ok {
				continue
			}
			visited[value] = empty{}
			local_count += 1
			if len(graph[value]) == 0 {
				continue
			}
			for _, v := range graph[value] {
				s.push(v)
			}
		}
		fmt.Println(local_count, max_count)
		if local_count > max_count {
			fmt.Println("Appending for node ", node, local_count)
			max_count = local_count
		}

	}

	return uint(max_count)
}

// Problem is get the smallest path length
// Convert edge list to adgency list
// Set a counter, increase on every traversal
// DO a BFS, when we encounter the destination, return cost
func shortest_path(graph map[string][]string, initial string, dst string) int {
	type set struct {
		value string
		cost  int
	}

	path := []string{}
	q := queue[set]{
		values: []set{
			{value: initial, cost: 0},
		},
	}
	for {
		if len(q.values) == 0 {
			break
		}
		val := q.dequeue()
		path = append(path, val.value)
		if dst == val.value {
			fmt.Println("Shortest path becomes", path)
			return val.cost
		}

		values := graph[val.value]
		for _, v := range values {
			q.enqueue(set{value: v, cost: val.cost + 1})
		}
	}

	return -1
}

type adjecancy_list map[string][]string

func convert_edge_to_list(value [][]string) adjecancy_list {
	v := adjecancy_list{}
	for _, edges := range value {
		if l, ok := v[edges[0]]; ok {
			for _, e := range edges {
				if edges[0] != e {
					v[edges[0]] = append(l, e)
				}
				if _, ok := v[e]; !ok {
					v[e] = []string{edges[0]}
				}
			}
		} else {
			v[edges[0]] = append(v[edges[0]], edges[1])
			v[edges[1]] = append(v[edges[1]], edges[0])
		}

	}
	return v
}

/*
Problem is to basically count the number of connected L, which is island

	{"W", "L", "W", "W", "W"},
	{"W", "L", "W", "L", "W"},
	{"W", "W", "W", "L", "W"},
	{"L", "W", "L", "L", "W"},
	{"L", "L", "W", "W", "W"},

Go through each row, and then column
check if row and column is out of bound
mark as visited
if W no point in going further
else explore row-1, row +1, col -1 and col +1 which is basically moving up, down, left and righ
on each visit we mark these new as visited, and if we encounter visited we dont go futher
amazingly simple explaination by https://youtu.be/tWVWeAqZ0WU?si=rkqQbFMJJWdnNCxO
*/
func numIslands(grid [][]string) int {
	count := 0
	visited := map[string]struct{}{}

	// loop through each row
	for row := 0; row < len(grid); row++ {
		fmt.Println(len(grid[0]))
		// loop through each columns
		for col := 0; col < len(grid[0]); col++ {
			if explore_island(grid, row, col, visited) {
				count += 1
			}
		}
	}

	return count
}

func explore_island(grid [][]string, row int, col int, visited map[string]struct{}) bool {

	// check if out of bound row
	if row < 0 || row >= len(grid) {
		return false
	}

	if col < 0 || col > len(grid[0]) {
		return false
	}

	// its water
	if grid[row][col] == "0" {
		return false
	}

	pos := fmt.Sprintf("%d,%d", row, col)
	if _, ok := visited[pos]; ok {
		return false
	}

	visited[pos] = struct{}{}
	explore_island(grid, row-1, col, visited) // explore up
	explore_island(grid, row+1, col, visited) // explore down
	explore_island(grid, row, col+1, visited) // explore right
	explore_island(grid, row, col-1, visited) // explore left

	return true
}

func minimum_island(grid [][]string) int {
	min_count := 10000
	visited := map[string]struct{}{}

	// loop through each row
	for row := 0; row < len(grid); row++ {

		// loop through each columns
		for col := 0; col < len(grid[0]); col++ {

			temp := min_island_explore(grid, row, col, visited)
			if temp != 0 && temp < min_count {
				min_count = temp
			}
		}
	}

	return min_count
}

func min_island_explore(grid [][]string, row int, col int, visited map[string]struct{}) int {

	// check if out of bound row
	if row < 0 || row >= len(grid) {
		return 0
	}

	if col < 0 || col > len(grid[0]) {
		return 0
	}

	pos := fmt.Sprintf("%d,%d", row, col)
	if _, ok := visited[pos]; ok {
		return 0
	}

	// its water
	if grid[row][col] == "W" {
		return 0
	}

	visited[pos] = struct{}{}
	// if we are here, it means, it was a land.
	v1 := min_island_explore(grid, row-1, col, visited) // explore up
	v2 := min_island_explore(grid, row+1, col, visited) // explore down
	v3 := min_island_explore(grid, row, col+1, visited) // explore right
	v4 := min_island_explore(grid, row, col-1, visited) // explore left

	v := 1 + v1 + v2 + v3 + v4
	return v
}
