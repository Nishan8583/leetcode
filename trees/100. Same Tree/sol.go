/*
https://leetcode.com/problems/same-tree/description/
100. Same Tree
Solved
Easy
Topics
Companies

Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

Use queue, 2 queues for 2 different trees, push both left and right values in the queue respecitvely in same order
check if values matches in those queue
*/
package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	values []*TreeNode
}

func (q *Queue) push(n *TreeNode) {
	q.values = append(q.values, n)
}

func (q *Queue) pop() *TreeNode {
	if len(q.values) == 0 {
		return nil
	}

	v := q.values[0]
	q.values = q.values[1:]
	return v
}
func isSameTree(p *TreeNode, q *TreeNode) bool {
	pQueue := Queue{}
	qQueue := Queue{}

	pQueue.push(p)
	qQueue.push(q)

	for {

		if len(pQueue.values) == 0 || len(qQueue.values) == 0 {
			break
		}

		pValue := pQueue.pop()
		qValue := qQueue.pop()
		if pValue == nil && qValue == nil {
			continue
		} else if pValue == nil && qValue != nil {
			return false
		} else if pValue != nil && qValue == nil {
			return false
		}

		if pValue.Val != qValue.Val {
			return false
		}

		pQueue.push(pValue.Left)
		pQueue.push(pValue.Right)

		qQueue.push(qValue.Left)
		qQueue.push(qValue.Right)
	}

	if len(pQueue.values) != len(qQueue.values) {
		return false
	}

	return true
}

func main() {
	f1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}

	f2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}

	f3 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	f4 := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}

	f5 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 1},
	}

	f6 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}
	_ = f1
	_ = f2
	_ = f5
	_ = f6
	fmt.Println("f1 and f2", isSameTree(f1, f2))
	fmt.Println("f3 and f4", isSameTree(f3, f4))
	fmt.Println("f5 and f6", isSameTree(f5, f6))
}
