/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.


	Example 1:

Input: s = "()"
Output: true

Example 2:

Input: s = "()[]{}"
Output: true

Example 3:

Input: s = "(]"
Output: false

Solution:
We keep on pushing until we find a close tag, if we find a close tag, we check if last pop tag was the corresponding open tag
if it was we pop it, else we keep on pushing until we match the case

Simple create a map that holds close to open i.e.e ) to (

use stack
for ()[](), 1st ( ,now when u push ) check if it was opened before, by checking is stack, if so dont push it and pop stack
do same for [] and so on
([{}])
in this case ([{ and finally when we push } we see that it has opened, dont push and pop,  then we get ] and same thing repeats

*/

func isValid(s string) bool {
	valid_brases := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]byte, 0, len(s))
	stack = append(stack, s[0])

	for i := 1; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
			continue
		}

		// Top of the stack must have a open tag, if the next is the closing tag
		poped_tag := stack[len(stack)-1]

		// check if the next is the closing tag
		if corresponding_close_tag, ok := valid_brases[s[i]]; ok {

			// if it was a closing tag, check if the poped tag is correspoding open tag
			if corresponding_close_tag == poped_tag {
				stack = stack[:len(stack)-1] // removing from stack
			} else {
				stack = append(stack, s[i])
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	// If the condition was matched, then our stack will be empty
	if len(stack) != 0 {
		return false
	}

	return true
}