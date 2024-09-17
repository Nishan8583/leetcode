/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 Find kth element
 So we need to traverse correctly and append lowest first
 First traverse left, if both left and right is empty, append current, means we are at leeaf node
 and since we are traversing left first, this has to be leftmost node
Then we append current node, since rigt will be bigger
THen we go right
*/

func kthSmallest(root *TreeNode, k int) int {
	s := []int{}

	var dfsLeft func(node *TreeNode)
	dfsLeft = func(node *TreeNode) {

		// Since we go left first, we wll reach left most leaf fisrt
		// this if both left and right is empty we are at left most leaf
		// so append it
		if node.Left == nil && node.Right == nil {
			s = append(s, node.Val)
			return
		}

		// if left node is not nil
		// continue down left first
		if node.Left != nil {
			dfsLeft(node.Left)
		}

		// after appending left, append current, because right will always be bigger
		s = append(s, node.Val)

		// go to right node at the last
		if node.Right != nil {
			dfsLeft(node.Right)
		}

	}

	dfsLeft(root)

	// k is going to be one greater since k is counted by 1
	k = k - 1

	if k >= len(s) {
		return s[0]
	}

	return s[k]
}