package main

import "fmt"

type Trie struct {
	childrens map[rune]*Trie
	word      bool
}

func Constructor() Trie {
	return Trie{
		childrens: make(map[rune]*Trie),
		word:      false,
	}
}

func (this *Trie) Insert(word string) {
	current := this
	for _, v := range word {
		if child, ok := current.childrens[v]; ok {
			current = child
		} else {
			child = &Trie{
				childrens: make(map[rune]*Trie),
			}
			current.childrens[v] = child
			current = child
		}
	}
	current.word = true
}

func (this *Trie) Search(word string) bool {
	current := this
	for _, v := range word {
		child, ok := current.childrens[v]
		if !ok {
			return false
		}
		current = child
	}
	return current.word
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this
	for _, v := range prefix {
		child, ok := current.childrens[v]
		if !ok {
			return false
		}
		current = child
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	t := Constructor()
	t.Insert("apple")
	fmt.Println(t)
	fmt.Println(t.Search("apple"))
}
