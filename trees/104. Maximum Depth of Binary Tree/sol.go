/*
Question: https://leetcode.com/problems/diameter-of-binary-tree/description/
Neetcode Soln: https://www.youtube.com/watch?v=bkxqA8Rfv04&ab_channel=NeetCode
Given the root of a binary tree, return the length of the diameter of the tree.
The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
The length of a path between two nodes is represented by the number of edges between them.

Example 1:

Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].

Example 2:

Input: root = [1,2]
Output: 1

Constraints:

	The number of nodes in the tree is in the range [1, 104].
	-100 <= Node.val <= 100
*/
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var maxDiameter int
	dfs(root, &maxDiameter)
	return maxDiameter

}

func dfs(node *TreeNode, maxDiameter *int) int {

	// a single node will have a height of 1, so an empty will have height of -1
	if node == nil {
		return 0
	}

	l := dfs(node.Left, maxDiameter)
	r := dfs(node.Right, maxDiameter)
	diameter := l + r // height of left + right subtree, gives the max distance between the edge leaf between them
	if diameter > *maxDiameter {
		*maxDiameter = diameter
	}

	return max(l, r) + 1 // add 1 to calculate for the current node as well

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
