/*

Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

    The left subtree of a node contains only nodes with keys less than the node's key.
    The right subtree of a node contains only nodes with keys greater than the node's key.
    Both the left and right subtrees must also be binary search trees.

	Solution: I first trieed a simple DFS, but it turns out we need to check with parent value as well

	We need to pass the leftboundary, i.e.e the current value must be bigger than that
	and pass rght boundary i.e. current value must be smaller than that
	and also do the regular binary search check
	The right boundary is upaded when we do dfs on the left to current value,
	which means next node will have to be less than the updated vlaue we are passing, current nodes parent along with the left child
	The left boundary is updated for right dfs

*/

package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	// during start left boundary is the most minimum value and right is most max value
	return dfs(root, math.MinInt, math.MaxInt)
}

func dfs(node *TreeNode, leftBoundary, rightBoundary int) bool {
	if node == nil {
		return true
	}

	// THe current value must be gte than left and less than right
	if node.Val <= leftBoundary || node.Val >= rightBoundary {
		return false
	}

	// simple BST check
	if node.Left != nil {
		if node.Val <= node.Left.Val {
			return false
		}
	}
	if node.Right != nil {
		if node.Val >= node.Right.Val {
			return false
		}
	}

	// for left node, update the right boundary, i.e. the left child must be less than its grandparent
	// for right node, pass in right boundary, i.e.e the right child must be greater than grandparent
	return dfs(node.Left, leftBoundary, node.Val) && dfs(node.Right, node.Val, rightBoundary)
}
