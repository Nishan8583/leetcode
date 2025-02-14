package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	ws1 := make([]string, 0, len(word1))
	ws2 := make([]string, 0, len(word2))
	for _, v := range word1 {
		ws1 = append(ws1, string(v))
	}

	for _, v := range word2 {
		ws2 = append(ws2, string(v))
	}

	final := ""
	for len(ws1) > 0 || len(ws2) > 0 {
		if len(ws1) > 0 {
			val := ws1[0]
			ws1 = ws1[1:]
			final = final + val
		}

		if len(ws2) > 0 {
			val := ws2[0]
			ws2 = ws2[1:]
			final = final + val
		}
	}

	return final
}

func main() {
	fmt.Println(mergeAlternately("abc", "pqr"))
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcd", "pq"))
}
