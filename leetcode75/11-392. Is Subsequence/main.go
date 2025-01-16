package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) < len(s) {
		return false
	}
	left := 0
	for _, v := range t {
		if v == rune(s[left]) {
			left += 1
		}
		if left >= len(s) {
			return true
		}
	}

	return left >= len(s)
}
