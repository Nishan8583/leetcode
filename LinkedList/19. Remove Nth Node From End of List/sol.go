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

func (q *queue) dequeNthElement(n int) {
	count := 1
	for i := len(q.vals) - 1; i >= 0; i-- {
		if count == n {

			firstHalf := q.vals[:i]

			s := q.vals[i+1:]
			for _, v := range s {
				fmt.Println("Second Half", v)
				firstHalf = append(firstHalf, v)
			}
			//final := append(firstHalf, seconbdHalf...)
			q.vals = firstHalf

			return
		}
		count++
	}
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	curr := head
	q := queue{}
	for curr != nil {
		q.enque(curr)
		fmt.Println("value", curr)
		curr = curr.Next
	}

	q.dequeNthElement(n)
	//fmt.Printf("%+v\n", q.vals)
	for i := 0; i < len(q.vals); i++ {
		fmt.Println("Current value", q.vals[i])

	}

	head = nil
	if len(q.vals) == 0 {
		fmt.Println("Empty values, returning nil")
		return nil
	}
	head = q.vals[0]
	curr = head
	//fmt.Println("First value", q.vals[0])
	for i := 1; i < len(q.vals); i++ {
		//fmt.Println("Current value", q.vals[i])
		curr.Next = q.vals[i]
		curr = curr.Next
	}
	curr.Next = nil
	fmt.Println("Final Message")
	curr = head
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}

	return nil
}

func prepareNodeForTest(values []int) *ListNode {

	head := &ListNode{Val: values[0]}

	current := head

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}

	current = head
	fmt.Println("----------------<Prepared values>-------------------")
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
	fmt.Println("----------------</Prepared values>-------------------")
	return head
}

func main() {

	head := prepareNodeForTest([]int{1, 2, 3, 4, 5})

	removeNthFromEnd(head, 2)

	current := head

	fmt.Println("----------------<Result>-------------------")
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}

	fmt.Println("----------------</Result>-------------------")

	head = prepareNodeForTest([]int{1, 2})

	removeNthFromEnd(head, 1)

	current = head

	fmt.Println("----------------<Result>-------------------")
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}

	fmt.Println("----------------</Result>-------------------")

}
