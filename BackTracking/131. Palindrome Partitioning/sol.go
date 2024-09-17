/*
Given a string s, partition s such that every
substring
of the partition is a
palindrome
. Return all possible palindrome partitioning of s.



Example 1:

Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]

Example 2:

Input: s = "a"
Output: [["a"]]



Constraints:

    1 <= s.length <= 16
    s contains only lowercase English letters.



Solution: Decision tree
*/

package main

func partition(s string) [][]string {
	ans := [][]string{} // holds the final return
	curr := []string{}  // holds the valid palindromes in current loop

	var dfs func(int)

	dfs = func(idx int) {

		// If we reached to this state, means all the substrings turned out to be valid palindromes in this branch
		if idx == len(s) {
			ans = append(ans, append([]string{}, curr...))
		}

		// Now brute force through each of the partitions here to see if they are valid palindromes
		for i := idx; i < len(s); i++ {
			if isValidPalindrome(s[idx : i+1]) {
				// since it was a valid palindrome, append the substring to the part
				curr = append(curr, s[idx:i+1])
				// now check for the next index
				dfs(i + 1)

				// pop off the last value, because for the index we just passed, if it had
				curr = curr[:len(curr)-1]

			}
		}
	}

	dfs(0)
	return ans
}

func isValidPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}
	return true
}
