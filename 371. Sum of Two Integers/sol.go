/*
Execlusive OR does our addition, but it does not handle the carryover, so we need to add and bit shift left to handle  it
https://leetcode.com/problems/sum-of-two-integers/description/
https://www.youtube.com/watch?v=gVUrDV4tZfY&ab_channel=NeetCode
*/
func getSum(a int, b int) int {
	carry := 0
	for b != 0 {
		carry = (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a
}