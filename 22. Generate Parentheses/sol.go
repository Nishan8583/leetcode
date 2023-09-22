/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.



Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:

Input: n = 1
Output: ["()"]

*/

import "strings"

func generateParenthesis(n int) []string {
	var stack []string
	var res []string

	var backtrack func(int, int)
	backtrack = func(openN int, closedN int) {
		if openN == n && closedN == n && openN == closedN {
			res = append(res, strings.Join(stack, ""))
			return
		}
		if openN < n {
			stack = append(stack, "(")
			backtrack(openN+1, closedN)
			pop(&stack)
		}
		if closedN < openN {
			stack = append(stack, ")")
			backtrack(openN, closedN+1)
			pop(&stack)
		}
	}
	backtrack(0, 0)
	return res

}

func pop(list *[]string) {
	length := len(*list)
	*list = (*list)[:length-1]
}