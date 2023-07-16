package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Logic behind the addition
	// let number be a and b
	// add a+b, if a+b is >= 10, we need to carryover the more than ten value to the next node, i.e. if we
	// get 9, then 1, is more than 9, so 1 must be carried to the next nodes addition
	//  (a+b)/10, get the whole number, multiply the whole number with 10, substrasct this product with the sum, put it in carryover
	// 14 + 7 = `2`1 => 21/10 => 2 => 10*2=20 => 21-20 => 1that weill be carried over to next node
	// 33+25. sum =  58 => 58/10 => q = 5 ,58-(10*5) = 8, Here carryover=q=5, current value = 8
	// In next node, check if carryover has value, if so add it too, set carryover to 0
	// keep on going
	// Create a final node, if 2 input list nodes are empty, return the node
	// else Call AddNode, this is a recursive function
	// if both input node is empty,
	// put the final carryover value as its current value, and return
	// if either one of them are nil, set value to 0 and call recursively
	// if none of them is empty
	// Do the logic of addition as mentioned above, and call recursively if neithfer of them is zero.

	var final *ListNode

	if (l1 != nil) && (l2 != nil) {
		final = &ListNode{}
		AddNode(l1, l2, final, 0)
	}

	return final
}

func AddNode(l1 *ListNode, l2 *ListNode, result *ListNode, carryOver int) *ListNode {

	next := &ListNode{}

	if (l1 == nil) && (l2 == nil) {
		if carryOver == 0 {
			return nil
		}
		result.Val = carryOver
		return result
	} else if l1 == nil {
		val, carryOver := calculation(l2.Val, 0, carryOver)
		result.Val = val
		result.Next = AddNode(nil, l2.Next, next, carryOver)
		return result
	} else if l2 == nil {
		val, carryOver := calculation(l1.Val, 0, carryOver)
		result.Val = val
		result.Next = AddNode(l1.Next, nil, next, carryOver)

		return result
	}

	val, carryOver := calculation(l1.Val, l2.Val, carryOver)
	result.Val = val

	if (l1.Next == nil) && (l2.Next == nil) {
		if carryOver != 0 {
			result.Next = &ListNode{Val: carryOver}
		}
		return result
	}

	result.Next = AddNode(l1.Next, l2.Next, next, carryOver)
	return result
}

func calculation(a, b, carryOver int) (result, carryOverR int) {
	sum := a + b + carryOver
	carryOver = 0

	if sum >= 10 {
		q := int(sum / 10)
		result = sum - (10 * q)
		carryOver = q
	} else {
		result = sum
	}
	return result, carryOver
}
func main() {
	l1 := ListNode{
		9, &ListNode{
			9, &ListNode{
				1, nil,
			},
		},
	}

	l2 := ListNode{
		1, nil,
	}

	fmt.Println(addTwoNumbers(&l1, &l2))
}
