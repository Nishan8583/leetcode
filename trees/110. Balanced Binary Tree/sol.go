/*
https://leetcode.com/problems/balanced-binary-tree/description/
Given a binary tree, determine if it is
height-balanced: A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.
.

Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: true

Example 2:

Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
*/
package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, b := dfs(root)
	return b
}

func dfs(node *TreeNode) (int, bool) {

	if node == nil {
		return 0, true
	}

	leftHeight, lisBalanced := dfs(node.Left)
	rightHeight, risBalanced := dfs(node.Right)
	curHeight := 0
	diff := 0
	if leftHeight == rightHeight {
		curHeight = leftHeight + 1
	} else if leftHeight > rightHeight {
		curHeight = leftHeight + 1
		diff = leftHeight - rightHeight
	} else {
		curHeight = rightHeight + 1
		diff = rightHeight - leftHeight
	}

	if !lisBalanced || !risBalanced {
		return curHeight, false
	}

	if diff <= 1 {
		return curHeight, true
	}

	return curHeight, false

}

func main() {
	t1 := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	t2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	fmt.Println(isBalanced(t1))
	fmt.Println(isBalanced(t2))
}
