package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func divides(s, t string) bool {
	if len(t)%len(s) != 0 {
		return false
	}

	repeat := int(len(t) / len(s))
	checkString := ""
	for i := 0; i < repeat; i++ {
		checkString = checkString + s
	}

	return checkString == t
}

func gcdOfStrings(str1 string, str2 string) string {
	minVal := min(len(str1), len(str2))
	result := ""
	for i := 1; i < minVal+1; i++ {
		canditate := str1[:i]
		if divides(canditate, str1) && divides(canditate, str2) {
			result = canditate
		}
	}

	return result
}
