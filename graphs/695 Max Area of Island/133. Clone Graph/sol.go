package main

/*

Given a reference of a node in a connected undirected graph.

Return a deep copy (clone) of the graph.

Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.

class Node {
    public int val;
    public List<Node> neighbors;
}



Test case format:

For simplicity, each node's value is the same as the node's index (1-indexed). For example, the first node with val == 1, the second node with val == 2, and so on. The graph is represented in the test case using an adjacency list.

An adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.

The given node will always be the first node with val = 1. You must return the copy of the given node as a reference to the cloned graph.


solution:

Use hashmap
map[Old Nodes][]

create new nodes, if new nodes are already present, return its
If clone is not created already, create new, and do recursive function call for the neighbors
*/

type Node struct {
	Val       int
	Neighbors []*Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	oldToNew := map[*Node]*Node{}

	var dfsClone func(*Node) *Node

	dfsClone = func(node *Node) *Node {
		// if node was already cloned
		if clonedNode, ok := oldToNew[node]; ok {
			return clonedNode
		}

		copy := &Node{
			Val: node.Val,
		}
		oldToNew[node] = copy

		for _, n := range node.Neighbors {
			// clone each Neighbours
			copy.Neighbors = append(copy.Neighbors, dfsClone(n))
		}
		return copy
	}

	return dfsClone(node)
}
