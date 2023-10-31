package main

import (
	"container/heap"
	"fmt"
)

type cs [][3]int // first is index of first point, second is index of second point, third is cost of connection between those poitns

func (c cs) Len() int           { return len(c) }
func (c cs) Less(i, j int) bool { return c[i][2] < c[j][2] }
func (c cs) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func (c *cs) Push(x any) {
	*c = append(*c, x.([3]int))
}

func (c *cs) Pop() any {
	old := *c
	new := len(old)
	x := old[new-1]
	*c = old[0 : new-1]

	return x

}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func minCostConnectPoints(points [][]int) int {
	min_heap := &cs{}
	heap.Init(min_heap)

	// populating heap
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			heap.Push(min_heap, [3]int{i, j, abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])})
		}
	}
	visited := map[int]bool{}
	total := 0
	list := [][3]int{}
	for min_heap.Len() > 0 {
		e := heap.Pop(min_heap).([3]int)
		// if both points is visited continue
		if visited[e[0]] && visited[e[1]] {
			continue
		}

		// if total was not zero, and both are not visited, append
		if total != 0 && !visited[e[0]] && !visited[e[1]] {
			fmt.Println("appending", points[e[0]])
			list = append(list, e)
			continue
		}

		fmt.Println("increasing ", points[e[0]], points[e[1]], e[2])
		total += e[2]
		visited[e[0]] = true
		visited[e[1]] = true
		// now append all the connections
		for _, v := range list {
			fmt.Println("increasing ", points[v[0]], points[v[1]], v[2])
			heap.Push(min_heap, v)
		}

		list = [][3]int{}
	}

	return total

}

func main() {

	//value := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	//fmt.Println(minCostConnectPoints(value))
	//fmt.Println(minCostConnectPoints([][]int{{3, 12}, {-2, 5}, {-4, 1}}))
	fmt.Println(minCostConnectPoints([][]int{{0, 0}, {1, 1}, {1, 0}, {-1, 1}}))
	fmt.Println(minCostConnectPoints([][]int{{2, -3}, {-17, -8}, {13, 8}, {-17, -15}}))
}
