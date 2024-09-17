package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {

	if t == "" {
		return ""
	}

	tCount := map[rune]int{}
	window := map[rune]int{}

	// setting up the map of character to count in tCount
	for _, char := range t {
		if _, ok := tCount[char]; ok {
			tCount[char]++
		} else {
			tCount[char] = 1
		}
	}

	have, need := 0, len(t)
	res := [2]int{-1, -1}
	resLen := math.MaxInt
	l := 0 // left pointer
	for r, char := range s {
		if _, ok := window[char]; ok {
			window[char]++
		} else {
			window[char] = 1
		}

		// if the current character is in t, and the count is exactly the same as in tcount i.e. the condition just fulfulled
		if _, ok := tCount[char]; ok && tCount[char] == window[char] {
			have++

		}

		// when we have the exact count of character counts that we need
		for have == need {

			// if the portion we found is less then previous set portion
			if (r - l + 1) < resLen {
				res = [2]int{l, r}
				resLen = (r - l + 1)
			}

			// pop from the left of the window
			window[rune(s[l])] -= 1
			if _, ok := tCount[char]; ok && tCount[char] > window[char] {
				have--
			}
			l += 1

		}
	}

	if resLen != math.MaxInt {
		return s[res[0] : res[1]+1]
	}
	return ""
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
