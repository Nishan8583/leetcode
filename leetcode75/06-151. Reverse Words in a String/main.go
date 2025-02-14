package main

import (
	"slices"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	l := strings.Split(s, " ")
	slices.Reverse(l)
	result := ""

	for _, v := range l {
		value := string(v)
		if value == "" {
			continue
		}
		result = result + value + " "
	}

	return strings.TrimSpace(result)
}
