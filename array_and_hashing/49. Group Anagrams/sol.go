/*
solve using a map of array of count of each letter in string (i.e. [0]*26, first represents count of a and so on) to array of anagrams
u can directly append at the required position by substracting current rune with value 'a'
then just append to same position
*/
package main

import "fmt"

type bit_array [26]int

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{
			{strs[0]},
		}
	}
	count := make(map[[26]rune][]string)
	for _, s := range strs {

		temp := [26]rune{}
		for _, v := range s {
			temp[v-'a'] += 1
		}
		count[temp] = append(count[temp], s)
	}

	values := [][]string{}
	for _, v := range count {
		values = append(values, v)
	}

	return values

}

func main() {
	input := []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}
	fmt.Println(groupAnagrams(input))
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
