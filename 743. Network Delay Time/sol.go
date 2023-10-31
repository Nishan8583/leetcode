package main

/*
https://leetcode.com/problems/network-delay-time/submissions/
We use dikstra to solve this problem
Dikstra is greedy, 	Push the weight and destination in a min heap
 get the latest one and put in the smallest

 there are some minor addition for specific question
 https://neetcode.io/courses/advanced-algorithms/14
*/
import (
	"container/heap"
	_ "container/heap"
	"fmt"
)

type node struct {
	name   int
	weight int
}

type nodeSlice []node

func (n nodeSlice) Len() int           { return len(n) }
func (n nodeSlice) Less(i, j int) bool { return n[i].weight < n[j].weight }
func (n nodeSlice) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

func (n *nodeSlice) Push(x any) {
	*n = append(*n, x.(node))
}

func (n *nodeSlice) Pop() any {
	old := *n
	new := len(old)
	x := old[new-1]
	*n = old[0 : new-1]
	return x
}

func convertToMatric(times [][]int) map[int]map[int]int {
	matrix := make(map[int]map[int]int)
	for _, v := range times {
		if conn, ok := matrix[v[0]]; ok {
			conn[v[1]] = v[2]
			matrix[v[0]] = conn
		} else {
			matrix[v[0]] = map[int]int{
				v[1]: v[2],
			}
		}
	}
	return matrix
}

func networkDelayTime(times [][]int, n int, k int) int {
	minHeap := &nodeSlice{}
	heap.Init(minHeap)

	/*edges := map[int][][]int{}
	for u,v := range times{
		if ws,ok := edges[u];ok {
			ws = append(ws,[]int{v[1],v[0]} )
		}
	}*/
	m := convertToMatric(times)
	smallest := 0
	//fmt.Println(m)
	//fmt.Println("Searching for ", k)
	heap.Push(minHeap, node{k, 0})

	//fmt.Println("minHeap", minHeap)

	routes := map[int]int{}
	for {
		if minHeap.Len() == 0 {
			break
		}
		node_value := heap.Pop(minHeap).(node)
		if _, ok := routes[node_value.name]; ok {
			continue
		}
		routes[node_value.name] = node_value.weight

		/*if node_value.weight > smallest && node_value.weight != 0 {
			smallest = node_value.weight
		}*/
		smallest = node_value.weight

		if v, ok := m[node_value.name]; ok {
			for dest, weight := range v {
				if _, ok2 := routes[dest]; ok2 {
					continue
				}
				heap.Push(minHeap, node{dest, weight + node_value.weight})
			}
		}
	}
	if len(routes) == n {
		return smallest
	} else {
		return -1
	}
}

func main() {
	times := [][]int{
		{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}

	fmt.Println(networkDelayTime(times, 4, 2))
	fmt.Println(networkDelayTime([][]int{{1, 2, 1}}, 2, 2))
	fmt.Println(networkDelayTime([][]int{{1, 2, 1}, {2, 1, 3}}, 2, 2))
	fmt.Println(networkDelayTime([][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 1}}, 3, 2))
}
