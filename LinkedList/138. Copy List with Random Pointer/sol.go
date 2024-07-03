/*

https://leetcode.com/problems/copy-list-with-random-pointer/description/
Medium
Topics
Companies
Hint
A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.

Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.

For example, if there are two nodes X and Y in the original list, where X.random --> Y, then for the corresponding two nodes x and y in the copied list, x.random --> y.

Return the head of the copied linked list.

The linked list is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:

val: an integer representing Node.val
random_index: the index of the node (range from 0 to n-1) that the random pointer points to, or null if it does not point to any node.
Your code will only be given the head of the original linked list.

Solution:
Use hashmap, map old to new,
A:a lets say A.Next=B and A.Random=C
B:b
C:c

Now When we do map[A.Next] it will get pointer to B, and when we do A.Random it will get c and WHEN we look that up in hashmap it will give us the new "c"
A.Next=B, and map[B]=b, b is the new node
A.Random=C, and map[C]=c, c here is the new node

*/

package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	// maps to old nodes to the new copy nodes that we will create
	oldToNew := map[*Node]*Node{
		nil: nil,
	}

	// first pass, just create new nodes and map in the hashmap
	cur := head
	for cur != nil {
		oldToNew[cur] = &Node{
			Val: cur.Val,
		}
		cur = cur.Next
	}

	// second pass update the next and random
	cur = head
	for cur != nil {
		copyNode := oldToNew[cur]
		copyNode.Next = oldToNew[cur.Next]
		copyNode.Random = oldToNew[cur.Random]
		cur = cur.Next
	}

	return oldToNew[head]

}
