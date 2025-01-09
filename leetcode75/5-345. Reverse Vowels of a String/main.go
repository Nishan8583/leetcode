package main

import "fmt"

func reverseVowels(s string) string {
	vowel_set := map[rune]struct{}{
		'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}, 'A': {},
		'E': {}, 'I': {}, 'O': {}, 'U': {},
	}
	stack := []rune{}

	for _, c := range s {
		if _, ok := vowel_set[c]; ok {
			stack = append(stack, c)
		}
	}
	// fmt.Println("Stack", stack)
	result := ""
	for _, c := range s {
		if _, ok := vowel_set[c]; ok {
			final_val := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = result + string(final_val)
		} else {
			result = result + string(c)
		}
	}

	return result
}

func main() {
	fmt.Println(reverseVowels("OE"))

	fmt.Println(reverseVowels("IceCreAm"))
}
