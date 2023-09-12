/**

https://leetcode.com/problems/invert-binary-tree/description/
Given the root of a binary tree, invert the tree, and return its root.
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

 DO a BFS, for me used queue, and then swithch left with right
*/
type queue struct {
	vals []*TreeNode
}

func (q *queue) enque(val *TreeNode) {
	q.vals = append(q.vals, val)
}
func (q *queue) deque() *TreeNode {
	if len(q.vals) == 0 {
		return nil
	}
	f := q.vals[0]

	if len(q.vals) > 1 {
		q.vals = q.vals[1:]
	} else {
		q.vals = []*TreeNode{}
	}

	return f

}
func invertTree(root *TreeNode) *TreeNode {
	q := queue{}
	if root == nil {
		return nil
	}

	l := root.Left
	r := root.Right

	// swapping left and right
	root.Left = r
	root.Right = l

	q.enque(l)
	q.enque(r)

	for {
		if len(q.vals) == 0 {
			break
		}
		left := q.deque()
		right := q.deque()
		if left == nil && right == nil {
			continue
		}

		if left != nil {
			ll := left.Left
			lr := left.Right
			left.Right = ll
			left.Left = lr
			q.enque(ll)
			q.enque(lr)
		}

		if right != nil {
			rl := right.Left
			rr := right.Right
			right.Right = rl
			right.Left = rr
			q.enque(rl)
			q.enque(rr)
		}

	}
	return root
}



