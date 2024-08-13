/*
763. Partition Labels
Solved
Medium
Topics
Companies
Hint

You are given a string s. We want to partition the string into as many parts as possible so that each letter appears in at most one part.

Note that the partition is done so that after concatenating all the parts in order, the resultant string should be s.

Return a list of integers representing the size of these parts.

Example 1:

Input: s = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
This is a partition so that each letter appears in at most one part.
A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits s into less parts.

Example 2:

Input: s = "eccbbbbdec"
Output: [10]

Constraints:

	1 <= s.length <= 500
	s consists of lowercase English letters.

Solution:
1. Hashmap containng rune to its last last index.
2. size = current size, end  = max end of character reached
3. if i == end, current chracter was not reached till end
Ex: "ababcbacadefegdehijhklij"
lasTIndex map for a its 8, for b its 5
looping, a end is set to 8, b has last index 5, so end is not updated,
so on, until i is end then update result
*/package main

func partitionLabels(s string) []int {
	lastIndex := map[rune]int{}
	for i, v := range s {
		lastIndex[v] = i
	}

	size := 0
	end := 0
	res := []int{}
	for i, v := range s {
		size += 1
		end = max(lastIndex[v], end)

		if i == end {
			res = append(res, size)
			size = 0
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
