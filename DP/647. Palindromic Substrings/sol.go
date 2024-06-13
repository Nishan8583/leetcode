/*
Given a string s, return the number of palindromic substrings in it.

A string is a palindrome when it reads the same backward as forward.

A substring is a contiguous sequence of characters within the string.



Example 1:

Input: s = "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".

Example 2:

Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".



Constraints:

    1 <= s.length <= 1000
    s consists of lowercase English letters.



Solution:
"aaab"

1. We use two pointers l and r
2. l = 0, r = 0, when both same we see,its a palindrome cause "a" is same as "a"
3. Now for Odd increase l and r, if same we see value at l and r is same, its a palindrome, In second loop l=a at i=1 when l is decreased and r is increased, we will have aaa
4. If out of bound, we do for i=1
5. For even l=0, r=l+1, when l = 0 and r = 1, we will have "aa"
6. Repeat
*/

package main

func countSubstrings(s string) int {
	count := 0

	// imaging string "aaab"
	for i := 0; i < len(s); i++ {
		l := i
		r := i

		// counting for odd, intially it will be "a"
		for l >= 0 && r < len(s) && s[l] == s[r] {
			count++
			l--
			r++
			// when i =1, it will be "aaa" on the second loop
		}

		// for even
		l = i
		r = i + 1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			count++
			l--
			r++
		}
	}

	return count
}
