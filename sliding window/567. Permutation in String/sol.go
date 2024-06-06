package main

/*
Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.

Solution, we do a sliding window here, and maintain a count of number of chaarcters
Input: s1 = "ab", s2 = "eidbaooo"
for this
for s1: a and b, first count of a is 1 and b is 1, rest is 0
for s2: e is 1 and i is 1

then move on for s2, next windows is id
here r points to d and l points to 0
here, decrease e since we moved the window, check if value matches with count for s1, if so increase matches
also for e decrease to check if previous value did match, if so decrease matches

here
*/
func checkInclusion(s1 string, s2 string) bool {

	// if s1 > s2, s2 can not be a permutation of s1
	if len(s1) > len(s2) {
		return false
	}

	// slice, each position represents ascii character like 0 means a, 1 means b and so on
	// and the value means the count/number of occurences of that character
	s1Count := make([]int, 26)
	s2Count := make([]int, 26)

	matches := 0 // number of position values matched in both s1CharMap and s2CharMap

	// first increase the count of length of s1
	// Ex: Input: s1 = "ab", s2 = "eidbaooo"
	// increase a to 1 and b to 1 for s1
	// increase e to 1 and i to 1 for s2
	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

	// first check if values in char maps are same
	for i := 0; i < 26; i++ {
		if s1Count[i] == s2Count[i] {
			matches++
		}
	}
	l := 0
	// start for length of s1, because the first 3 we have already done above
	for r := len(s1); r < len(s2); r++ {

		// if the chracter count representation mathces, this means there was an instance where count of characters
		// were same, i.e. the string s1 is present in s2
		if matches == 26 {
			return true
		}

		index := s2[r] - 'a'
		s2Count[index]++

		// if the value count matches, increase matches
		if s1Count[index] == s2Count[index] {
			matches++
		} else if s1Count[index]+1 == s2Count[index] {
			// this case implies the previous value in current index matched (because we only increase by 1, -1 means previous values)
			// decrease matches
			matches--
		}

		// Now do the similar for the l (lelft pointer), but we will decrease because we will be increase
		// our sliding window by 1, and left will get updated
		index = s2[l] - 'a'
		s2Count[index]--

		// if the value count matches, increase matches
		if s1Count[index] == s2Count[index] {
			matches++
		} else if s1Count[index]-1 == s2Count[index] {
			// this case implies the previous value in current index matched (because we only decrease by 1, +1 means previous values)
			// decrease matches
			matches--
		}

		l++ // increase left pointer
	}

	return matches == 26
}
