/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

    MinStack() initializes the stack object.
    void push(int val) pushes the element val onto the stack.
    void pop() removes the element on the top of the stack.
    int top() gets the top element of the stack.
    int getMin() retrieves the minimum element in the stack.

You must implement a solution with O(1) time complexity for each function.

basically two stack, one holds the actual values, and the other holds lowest value at that position
EX:
Value	MinStack
-8			-8
5			-2
-2			-2
1			0
0			0
*/

package main

import "math"

type MinStack struct {
	valueStack []int
	minStack   []int
	prevLowest int
}

func Constructor() MinStack {
	return MinStack{
		valueStack: []int{},
		minStack:   []int{},
		prevLowest: math.MaxInt,
	}
}

func (this *MinStack) Push(val int) {
	this.valueStack = append(this.valueStack, val)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		if this.minStack[len(this.minStack)-1] > val {
			this.minStack = append(this.minStack, val)
		} else {
			this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
		}
	}

}

func (this *MinStack) Pop() {
	this.valueStack = this.valueStack[:len(this.valueStack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.valueStack[len(this.valueStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
