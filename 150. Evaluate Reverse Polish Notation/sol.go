/*
https://leetcode.com/problems/evaluate-reverse-polish-notation/description/

You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:

    The valid operators are '+', '-', '*', and '/'.
    Each operand may be an integer or another expression.
    The division between two integers always truncates toward zero.
    There will not be any division by zero.
    The input represents a valid arithmetic expression in a reverse polish notation.
    The answer and all the intermediate calculations can be represented in a 32-bit integer.

What is reverse polish notation https://en.wikipedia.org/wiki/Reverse_Polish_notation
*/

// Just use stack, push into stack, and on each enounter of operands, do operation and do the opreation and push the result in stack
// and finally pop and return, look at wikepeida article, it helps lmao
import "strconv"

type stack struct {
	values []int
}

func (s *stack) push(v int) {
	s.values = append(s.values, v)
}

func (s *stack) pop() int {
	v := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return v
}
func evalRPN(tokens []string) int {
	s := stack{}
	for _, num := range tokens {
		n, err := strconv.Atoi(num)
		if err == nil {
			s.push(n)
			continue
		}

		v1 := s.pop()
		v2 := s.pop()
		switch num {
		case "+":
			s.push(v2 + v1)
		case "-":
			s.push(v2 - v1)
		case "/":
			s.push(int(v2 / v1))
		case "*":
			s.push(v1 * v2)
		}
	}

	return s.pop()
}