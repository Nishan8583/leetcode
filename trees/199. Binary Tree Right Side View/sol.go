/*

Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.



Example 1:

Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]

Example 2:

Input: root = [1,null,3]
Output: [1,3]

Example 3:

Input: root = []
Output: []


Solution: Turns out we do BFS
					1
				   /  \
				2      3
				\		\
			      5      4
				    \
				    7
We use queue, and loop fr each level
First qeuue will have values for level 0, [1]
Loop through all the values, keep updating right most node, in the end append left and then right. Here right most value will be 1
Second level queue [2,3], Loop till len(queue) i.e. 2, keep updating right most node, in the end append left and then right. Here right most value will be 3
Third LEvel QUEUE [5,4], Loop till len(queue) i.e. 2, keep updating right most node, in the end append left and then right. Here right most value will be 4
Fourth Level queue [7], Loop till len(queue) i.e. 1, keep updating right most node, in the end append left and then right. Here right most value will be 7
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 { // Keep looping till queue is empty

		var rightMostNode *TreeNode // will hold rightmost value at each level
		levelLen := len(queue)      // will hold number of nodes at each level, Helps us prevent loopoing till len(queue) == 0

		for i := 0; i < levelLen; i++ { // Loop for all values in this current level
			node := queue[0]  // get first valu
			queue = queue[1:] // remove first value

			if node == nil { // if node is nil continue
				continue
			}

			rightMostNode = node // update it to current node, since we only loop till last value of current level, last
			// value will be right most node, since we append right most node at last
			// if right most node is nil, the right most will be its left counterpart

			// We are appending the right most node to the end
			// so in the end the node will hold the right most node for the level
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}

		if rightMostNode != nil {
			result = append(result, rightMostNode.Val)
		}

	}

	return result
}
