// My logic was, both should have same number of characters
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[byte]int, len(s))
	tMap := make(map[byte]int, len(t))

	for i := 0; i < len(s); i++ {
		if v, ok := sMap[s[i]]; ok {
			v += 1
			sMap[s[i]] = v
		} else {
			sMap[s[i]] = 1
		}
		if v, ok := tMap[t[i]]; ok {
			v += 1
			tMap[t[i]] = v
		} else {
			tMap[t[i]] = 1
		}
	}

	for k, v := range sMap {
		count, ok := tMap[k]
		if !ok {
			return false
		}
		if count != v {
			return false
		}
	}
	return true
}
