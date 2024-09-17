package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	values []*TreeNode
}

func (q *Queue) push(r *TreeNode) {
	q.values = append(q.values, r)
}

func (q *Queue) pop() *TreeNode {
	if len(q.values) == 0 {
		return nil
	}
	t := q.values[0]
	q.values = q.values[1:]

	return t
}
func checkIfSameTree(root, subtree *TreeNode) bool {
	if subtree.Left == nil && subtree.Right == nil {
		return true
	}
	rQueue := Queue{}
	sQueue := Queue{}

	rQueue.push(root)
	sQueue.push(subtree)

	for {

		// if either one of the queue is empty break out of loop
		if len(rQueue.values) == 0 || len(sQueue.values) == 0 {
			break
		}

		// pop the value from quee
		rValue := rQueue.pop()
		sValue := sQueue.pop()

		// if both is nil continue to check the other value
		if rValue == nil && sValue == nil {
			continue
		} else if rValue == nil && sValue != nil { //if one rValue is nil but sValue is not nil, values do not match, thus return false
			return false
		} else if sValue == nil && rValue != nil { // same as above case
			return false
		}

		// if val does not match, return false
		if rValue.Val != sValue.Val {
			return false
		}

		// push left
		rQueue.push(rValue.Left)
		sQueue.push(sValue.Left)

		// push right
		rQueue.push(rValue.Right)
		sQueue.push(sValue.Right)
	}

	// if length does not match, means the number of nodes does not match
	if len(sQueue.values) != len(rQueue.values) {
		return false
	}

	return true
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	mainQueue := Queue{}
	mainQueue.push(root)

	ret := false
	for {
		if len(mainQueue.values) == 0 {
			break
		}

		v := mainQueue.pop()
		if v == nil {
			continue
		}

		if v.Val == subRoot.Val {
			ret = checkIfSameTree(v, subRoot)
			//fmt.Printf("matched but not equals %+v and %+v\n", v, subRoot)
			if ret == true {
				return ret
			}
		}

		mainQueue.push(v.Left)
		mainQueue.push(v.Right)

	}

	return ret
}

func main() {
	r1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,

			Right: &TreeNode{
				Val: 1,

				Right: &TreeNode{
					Val: 1,

					Right: &TreeNode{
						Val: 1,

						Right: &TreeNode{
							Val: 1,

							Right: &TreeNode{
								Val: 1,

								Right: &TreeNode{
									Val: 1,

									Right: &TreeNode{
										Val: 1,

										Right: &TreeNode{
											Val: 1,

											Right: &TreeNode{
												Val:  1,
												Left: &TreeNode{Val: 2},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	s2 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,

			Right: &TreeNode{
				Val: 1,

				Right: &TreeNode{
					Val: 1,

					Right: &TreeNode{
						Val: 1,

						Right: &TreeNode{
							Val: 1,

							Left: &TreeNode{Val: 2},
						},
					},
				},
			},
		},
	}
	fmt.Println(isSubtree(r1, s2))
}
