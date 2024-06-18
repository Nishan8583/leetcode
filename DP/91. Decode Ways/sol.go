/*
Decode Ways
A string consisting of uppercase english characters can be encoded to a number using the following mapping:

'A' -> "1"
'B' -> "2"
...
'Z' -> "26"
To decode a message, digits must be grouped and then mapped back into letters using the reverse of the mapping above. There may be multiple ways to decode a message. For example, "1012" can be mapped into:

"JAB" with the grouping (10 1 2)
"JL" with the grouping (10 12)
The grouping (1 01 2) is invalid because 01 cannot be mapped into a letter since it contains a leading zero.

Given a string s containing only digits, return the number of ways to decode it. You can assume that the answer fits in a 32-bit integer.

Example 1:

Input: s = "12"

Output: 2

Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).
Example 2:

Input: s = "01"

Output: 0
Explanation: "01" cannot be decoded because "01" cannot be mapped into a letter.

Solution:
Using decision tree, either include next or dont
cache the value we calculated so we do not have to cache again
For 1012

				root

			1				10

		0		10		01		1
	1        1      2

1. If the value is 0, return 0, becasue our numbers start from 1
2. if value is 1, try i+2 recurse
3. if value is 2, check if next value is within range of 0-6, if yes, try i+2 recurse
4. Store the value in a cahche

for 1012
*/
package main

import "strings"

func numDecodings(s string) int {
	dp := make([]int, len(s)+1) // slice, holds cached vakues, each index holds the number of ways to decode till that index
	dp[len(s)] = 1              // If we have reached end of decision treee, we return 1, so for last it will be 1
	return recurse(0, s, &dp)
}

func recurse(i int, s string, dp *[]int) int {
	// check  if value is in cached, if so return the cached value
	if (*dp)[i] != 0 {
		return (*dp)[i]
	} else if s[i] == '0' {
		// if current value is '0' there is no possible mapping for this
		return 0
	}

	// Since we did not find the value in cache, lets branch of to the one where we include the next value
	res := recurse(i+1, s, dp)
	// first condition is if there is second value
	if ((i + 1) < len(s)) && ((string(s[i]) == "1") || (string(s[i]) == "2" && strings.Contains("0123456", string(s[i+1])))) {
		res += recurse(i+2, s, dp)
	}

	(*dp)[i] = res
	return res
}
