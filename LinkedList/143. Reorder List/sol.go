/*
143. Reorder List
Solved
Medium
Topics
Companies

You are given the head of a singly linked-list. The list can be represented as:

# L0 → L1 → … → Ln - 1 → Ln

Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …

You may not modify the values in the list's nodes. Only nodes themselves may be changed.

Solution:

Use queue
On even pop first on odd pop last
*/
package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}
type queue struct {
	vals []*ListNode
}

func (q *queue) enque(val *ListNode) {
	q.vals = append(q.vals, val)
}

func (q *queue) dequeFirst() *ListNode {
	if len(q.vals) != 0 {
		f := q.vals[0]
		if len(q.vals) < 2 {
			q.vals = []*ListNode{}
			return f
		}
		q.vals = q.vals[1:] // removing the first element from list
		return f
	}
	return nil
}

func (q *queue) dequeLast() *ListNode {
	if len(q.vals) != 0 {
		l := q.vals[len(q.vals)-1]
		q.vals = q.vals[:len(q.vals)-1]
		return l
	}
	return nil
}
func reorderList(head *ListNode) {
	temp := queue{}

	curr := head
	for curr != nil {
		temp.enque(curr)
		curr = curr.Next
	}

	head = temp.dequeFirst()
	helper(1, &temp, head)

}

func helper(index int, q *queue, curr *ListNode) {
	if len(q.vals) == 0 {
		curr.Next = nil
		return
	}

	if index%2 == 0 {
		curr.Next = q.dequeFirst()
	} else {
		curr.Next = q.dequeLast()
	}
	cur := curr.Next
	helper(index+1, q, cur)
}

func main() {
	head := &ListNode{Val: 1}
	second := &ListNode{Val: 2}
	third := &ListNode{Val: 3}
	fourth := &ListNode{Val: 4}
	fifth := &ListNode{Val: 5}

	head.Next = second
	second.Next = third
	third.Next = fourth
	fourth.Next = fifth

	current := head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
	reorderList(head)

	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
