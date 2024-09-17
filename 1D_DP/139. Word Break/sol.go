/*
139. Word Break
Solved
Medium
Topics
Companies

Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.



Example 1:

Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:

Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.

Example 3:

Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false



Constraints:

    1 <= s.length <= 300
    1 <= wordDict.length <= 1000
    1 <= wordDict[i].length <= 20
    s and wordDict[i] consist of only lowercase English letters.
    All the strings of wordDict are unique.



*/

/*
Solution:
 1. We do this by decision tree
 2. Initial  thought process:
    Loop through i := 0
    check if any word in wordDict matches
    if it does, update i to legnth of word
    start for next, and check the whole word again
 3. Caching decides to go from end to first,true
*/
package main

/*
This code implements a dynamic programming solution for the word break problem.
It uses a boolean array dp to track whether substrings of the input string s can be segmented into words from the wordDict.
The algorithm iterates through s from the end and checks if any substring can be segmented into words. If a segmentation is possible,
it updates the dp array accordingly. Finally, it returns whether the entire s can be segmented into words from wordDict by checking dp[0].
*/

func wordBreak(s string, wordDict []string) bool {

	// dp will be cahce of wethere this slice was able to break into words or not
	dp := make([]bool, len(s)+1)
	dp[len(s)] = true // if we have reached end of decision treee, we return true, so for last it will be true

	// we start  from the end
	for i := len(s) - 1; i >= 0; i-- {
		for _, w := range wordDict {

			// (i + len(w)) <= len(s) checks if current word is less than length of the given string
			if (i+len(w)) <= len(s) && s[i:i+len(w)] == w {
				dp[i] = dp[i+len(w)]
			}

			// if one of them matches, just break
			if dp[i] {
				break
			}
		}

	}

	// dp[0] will be set if all of them have been set to true
	return dp[0]
}
