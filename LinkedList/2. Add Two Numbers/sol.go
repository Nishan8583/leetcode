/*

https://leetcode.com/problems/add-two-numbers/description/
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.



Example 1:


Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]


Solution:
Simple, just go next in linked list add the numbers and carryober.
divide by 10 to get the carry over, ex: 12/10=1 which is gonna next carry over
sum is gonna be 12%10=2, because we need to get the last digit
Loop until both list is empty and carry is 0
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	curL1 := l1
	curL2 := l2
	head := &ListNode{} // points to the new head linked list
	currHead := head    // points to the current new linked list
	carry := 0

	for curL1 != nil || curL2 != nil || carry != 0 {

		v1 := 0
		v2 := 0

		if curL1 != nil {
			v1 = curL1.Val
		}
		if curL2 != nil {
			v2 = curL2.Val
		}

		// get the sum also add the previous carry over
		sum := v1 + v2 + carry
		carry = sum / 10
		sum = sum % 10
		currHead.Next = &ListNode{Val: sum, Next: nil}

		if curL1 != nil {
			curL1 = curL1.Next
		}
		if curL2 != nil {
			curL2 = curL2.Next
		}

		currHead = currHead.Next

	}
	return head.Next
}

func generateListNode(arr []int) *ListNode {
	var head *ListNode
	var current *ListNode

	for i := 0; i < len(arr); i++ {
		val := arr[i]
		node := &ListNode{Val: val, Next: nil}

		if head == nil {
			head = node
			current = head
		} else {
			current.Next = node
			current = current.Next
		}
	}

	return head
}

func main() {
	l1 := generateListNode([]int{2, 4, 3})
	l2 := generateListNode([]int{5, 6, 4})
	l3 := addTwoNumbers(l1, l2)
	curr := l3

	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}
