/*
https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/submissions/1256262178/
Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”



Example 1:

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.

Example 2:

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.

Example 3:

Input: root = [2,1], p = 2, q = 1
Output: 2


My solution:
We do dfs for bot the values so lets say we are searching for the 2,4 as shwon in eexamplee 2
We do dfs for 2, pass in the ancestors including itself in the stack, do again for 4
Then we check in stack, if ancestors match, since top of stack will have itself, current,
we are going from top to bottom, our slices will look like this
// 2,6
// 4,2,6
When we pop, we are getting the last value first,
first 6 is same, so least common ancestor is 6,
then 2
and its returned
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	values []*TreeNode
}

func (s *Stack) push(t *TreeNode) {
	s.values = append(s.values, t)
}

func (s *Stack) pop() *TreeNode {
	if len(s.values) == 0 {
		return nil
	}

	v := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return v
}

// 2,6
// 4,2,6
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	pStack := Stack{}
	qStack := Stack{}
	dfsSearch(root, p, &pStack)
	dfsSearch(root, q, &qStack)

	var lca *TreeNode

	for {
		if len(pStack.values) == 0 || len(qStack.values) == 0 {
			break
		}

		pVal := pStack.pop()
		qVal := qStack.pop()

		if pVal == qVal {
			lca = pVal
		}
	}
	return lca
}

func dfsSearch(node, v *TreeNode, stack *Stack) bool {

	if node == nil {
		return false
	}

	if node.Val == v.Val {
		stack.push(node)
		return true
	}
	ok := false
	if v.Val > node.Val {
		ok = dfsSearch(node.Right, v, stack)
	} else {
		ok = dfsSearch(node.Left, v, stack)
	}

	if ok {
		stack.push(node)
		return true
	}
	return false
}
