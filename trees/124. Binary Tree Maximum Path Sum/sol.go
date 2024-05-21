package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxSum int

func maxPathSum(root *TreeNode) int {
	maxSum = math.MinInt
	dfsSum(root)
	return maxSum
}

func dfsSum(node *TreeNode) int {
	if node == nil {
		return 0
	}

	// add values from both left and right
	leftSum := dfsSum(node.Left)
	rightSum := dfsSum(node.Right)

	if leftSum < 0 {
		leftSum = 0
	}
	if rightSum < 0 {
		rightSum = 0
	}
	// sum when both values from left and right are considered
	splitSum := node.Val + leftSum + rightSum
	if splitSum > maxSum {
		maxSum = splitSum
	}

	// return value should be the maximum value we can generate here
	// because when we think about a path, we can not go both left and right
	max := 0
	if leftSum >= rightSum {
		max = leftSum
	} else {
		max = rightSum
	}

	return node.Val + max
}
