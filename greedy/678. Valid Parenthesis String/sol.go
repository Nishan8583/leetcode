/*
678. Valid Parenthesis String
Solved
Medium
Topics
Companies
Hint

Given a string s containing only three types of characters: '(', ')' and '*', return true if s is valid.

The following rules define a valid string:

    Any left parenthesis '(' must have a corresponding right parenthesis ')'.
    Any right parenthesis ')' must have a corresponding left parenthesis '('.
    Left parenthesis '(' must go before the corresponding right parenthesis ')'.
    '*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string "".



Example 1:

Input: s = "()"
Output: true

Example 2:

Input: s = "(*)"
Output: true

Example 3:

Input: s = "(*))"
Output: true


Solution:
we count
() leftMin  = 1, leftMax = 1, when ), both will be set to 0

(*) first, both set to 1, second, leftMin is substracted to 0, leftMax is 2 (assuming * is ( ), on last both substracted, if leftMinx is below 0 update to 0

if leftMax is every below 0, it means even if we assume * to be ( , it is not enough, i.e. number of left paranthesis will be lower than right paranthesis in order required by the question

*/

package main

func checkValidString(s string) bool {
	// case when we consider * to not be (
	leftMin := 0

	// leftMax when we consider * to be (
	leftMax := 0

	for _, v := range s {
		switch v {
		case '(':
			leftMax++
			leftMin++
		case ')':
			leftMax--
			leftMin--
		default:
			leftMin--
			leftMax++
		}

		if leftMin < 0 {
			leftMin = 0
		}

		if leftMax < 0 {
			return false
		}
	}

	return leftMin == 0
}
