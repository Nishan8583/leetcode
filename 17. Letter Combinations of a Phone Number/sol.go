/*
https://leetcode.com/problems/letter-combinations-of-a-phone-number/
https://neetcode.io/courses/advanced-algorithms/12
17. Letter Combinations of a Phone Number
Medium
17.2K
901
Companies

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
Example 1:

Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Example 2:

Input: digits = ""
Output: []

Example 3:

Input: digits = "2"
Output: ["a","b","c"]

# ITs backtracking

A map of number to its respective ascii

	mappings := map[byte][]byte{
		'1': []byte{},
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}

	Go through the given digits
	ex: 23
	so we go to 2 first
	[a,b,c]
	for a we go through each of values for 3, ad, ae, af
	for b we do same
	for c also
*/
package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	mappings := map[byte][]byte{
		'1': []byte{},
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}

	final := []string{}
	current := ""
	backtrack(0, current, digits, &final, mappings)
	return final
}

func backtrack(i int, current string, digits string, res *[]string, mappings map[byte][]byte) {
	fmt.Println(i, current, digits)
	// initially the i is 0

	if len(current) == len(digits) {
		*res = append(*res, current)
		return
	}

	for _, v := range mappings[digits[i]] { // go through each digit, when"23", first is gonna be "2", then we pass i+1 which is 3
		backtrack(i+1, fmt.Sprintf("%s%s", current, string(v)), digits, res, mappings)
	}
}

func main() {
	fmt.Println(letterCombinations("35"))
}
