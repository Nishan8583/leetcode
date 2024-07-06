/*
19. Remove Nth Node From End of List
Solved
Medium
Topics
Companies
Hint

Given the head of a linked list, remove the nth node from the end of the list and return its head.



Example 1:

Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:

Input: head = [1], n = 1
Output: []

Example 3:

Input: head = [1,2], n = 1
Output: [1]

solution:
Use 2 pointers, keep the distance between those two pointers to n,

dummy -> 1 -> 2 -> 3 -> 4 -> 5
Left first pointer points to dummy, we use dummy because at the end we want left to point to one node before the value we want to delete, could we have increased the n to 1 instead and tried that?
Right is gonna point to n nodes after L
For n =2,
so first l = dummy, r = head, we do the first loop to initialize right, and then l=dummy, r = 3rd node
Then  we do our loop till right is not nil
We keep increasing the left and right pointer by next, because we want the distance between them to be constant of n
L = 1, R = 4
L = 2 R = 5
L = 3 R = nil

Now R is nil
Now make 3 point to 5
3.Next = 3.Next.Next

return dummy.Next
*/

package main

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	left := dummy
	right := head

	for n > 0 && right != nil {
		right = right.Next
		n--
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}

	// delete the node
	left.Next = left.Next.Next
	return dummy.Next
}
