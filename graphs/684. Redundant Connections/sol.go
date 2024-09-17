/*
In this problem, a tree is an undirected graph that is connected and has no cycles.

You are given a graph that started as a tree with n nodes labeled from 1 to n, with one additional edge added. The added edge has two different vertices chosen from 1 to n, and was not an edge that already existed. The graph is represented as an array edges of length n where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the graph.

Return an edge that can be removed so that the resulting graph is a tree of n nodes. If there are multiple answers, return the answer that occurs last in the input.



Example 1:

Input: edges = [[1,2],[1,3],[2,3]]
Output: [2,3]

Example 2:

Input: edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]
Output: [1,4]

Explaination:
The question is telling us to find the first edge that starts a cycle

Solution:
Using Union find
for  Input: edges = [[1,2],[1,3],[2,3]]
parents = [0,1,2,3]  // each position holds the parents for themselves, 0 is not used, just for simplicity, first nodes are parents to themselves
rank=[1,1,1,1] // rank holds the number of children they have

for 1 and 2, see of they have the same root parent, if they do, return edge, else update parents and edge
the one with the higher rank will be the parent
parents = [0,1,1,3]  // now 2 has parent 1
rank = [1,2,1,1]  // 1 has 2 children now

for 1,3 same
parents = [0,1,1,1] now 3 parent is also 1
rank = [1,3,1,1]

for 2,3 same, we find that both have parent 1. so return edge


*/

package main

func findRedundantConnection(edges [][]int) []int {
	// each position in parents holds parent of that node value
	// 0 is just to make it simple, initially, each node is parent to itself
	parents := []int{}
	for i := 0; i <= len(edges)+1; i++ {
		parents = append(parents, i)
	}

	// rank, holds the number of children the node has, in the beginning each will have 1
	rank := []int{}
	for i := 0; i <= len(edges)+1; i++ {
		rank = append(rank, 1)
	}

	// check if both node has same parent, if they do, return false
	// else the node with higher rank is parent of the other node
	// i.e. increase the rank of that node with rank of the other, and update parent array to indicate parent child relationship
	var union func(int, int) bool
	union = func(node1, node2 int) bool {
		p1, p2 := find_root_parent(node1, parents), find_root_parent(node2, parents)
		if p1 == p2 {
			return false
		}

		if rank[p1] > rank[p2] {
			parents[p2] = p1
			rank[p1] += rank[p2]
		} else {
			parents[p1] = p2
			rank[p2] += rank[p1]
		}
		return true
	}

	for _, nodes := range edges {
		if !(union(nodes[0], nodes[1])) {
			return nodes
		}
	}

	return []int{}
}

func find_root_parent(node int, parents []int) int {
	p := parents[node]

	for p != parents[p] {
		parents[p] = parents[parents[p]]
		p = parents[p]
	}

	return p
}
