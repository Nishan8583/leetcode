package main

import (
	"fmt"
	"time"
)

/*func wordBreak(s string, wordDict []string) bool {

	adj_list := map[string][]string{}
	for i, b := range s {
		if i == len(s)-1 {
			break
		}
		if v, ok := adj_list[string(b)]; ok {

			v = append(v, string(s[i+1]))
			adj_list[string(b)] = v
		} else {

			adj_list[string(b)] = []string{string(s[i+1])}
		}
	}
	fmt.Println(adj_list)

	for _, word := range wordDict {
		fmt.Println("Visiting", word)
		for i, letter := range word {
			if i == len(word)-1 {
				break
			}
			if n, ok := adj_list[string(letter)]; ok {
				found := false
				for _, nn := range n {
					if nn == string(word[i+1]) {
						found = true

						break
					}
				}
				if !found {
					return false
				}

			} else {
				return false
			}
		}
		fmt.Println("Found", word)
	}
	return true
}*/

func wordBreak(s string, wordDict []string) bool {

	// holds index of source string thats been already visited
	//visited := map[int]int{}

	doesStringExist := func(source string, word string, index int) (int, bool) {
		c := 0 // index for the word

		for index < len(word) {

			// check if each consequetive characters in both strings are not equal
			if word[c] != source[index] {
				//fmt.Printf("Words dont match? %s %s \n", string(word[c]), string(source[i]))
				return index, false
			}
			index++
			c++

		}
		fmt.Println("Match complete", index)
		return index, true
	}

	index_to_search := 0
	for {
		// we have found everything and now we returning
		if index_to_search >= len(s) {
			return true
		}

		found := false
		for _, word := range wordDict {
			if s[index_to_search] == word[0] {
				fmt.Println("sending index", word, index_to_search)
				i, ok := doesStringExist(s, word, index_to_search)
				if ok {
					index_to_search = i
					fmt.Println("found word", word)
					fmt.Println("new index", index_to_search)
					found = true
					time.Sleep(3 * time.Second)
					break
				}
			}
		}
		if !found {
			break
		}
	}

	return false
	/*for _, word := range wordDict {
		//fmt.Println("checking ", word)
		found := false
		for i, v := range s[index_to_search:] {
			if _, ok := visited[i]; ok {
				continue
			}
			if v == rune(word[0]) {
				//fmt.Println("Found first matching", string(v), i)
				if checkVisited(s, word, i)
			}
		}

	}*/

}

func main() {
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("bb", []string{"a", "b", "bbb", "bbbb"}))
}
