/*
two pointers, one holds index in string(i), another holds final value in array(fr)
first convert to lower case, NOw after that we are ignoring all space and special characters, and increase i and decrease fr until we get lowercase or numbers
we are doing this to avoid another loop to remove such characters, since ignoring them will have same affect.
keep checking out of bound of fr < i

	first loop {
		check if ascii, and not space and numbers is ok
		if match
		another loop {
			until (check if ascii, and not space and numbers is ok)  is matched for fr index (i.e. final value)
			we decrease fr index value
			(basically we are avoiding space and special characters here as well)
		}
		check if fr inbound
		check if fr and i values are equal, if not return true

}
*/
package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	fr := len(s) - 1
	for i := 0; i < len(s); i++ {
		v := s[i]
		if fr < 0 || fr < i {
			break
		}
		g := s[fr]
		//fmt.Printf("Comparing i=%d and fr=%d\n", i, fr)
		if (v <= 122 && v >= 97 && (v != ' ')) || (v >= 48 && v <= 57) {
			for {
				if fr < 0 {
					break
				}
				if (g <= 122 && g >= 97 && (g != ' ')) || (g >= 48 && g <= 57) {
					break
				}
				fr = fr - 1
				g = s[fr]
			}
			//fmt.Printf("Comparing v=%s and g=%s\n", string(v), string(g))
			if v != g {
				return false
			}
			fr = fr - 1
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"), " Expected True")

	fmt.Println(isPalindrome("race a car"), " Expected False")

	fmt.Println(isPalindrome(" "), " Expected True")

	fmt.Println(isPalindrome("0P"), " Expected False")
}
