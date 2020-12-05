package main

import "fmt"

func isMatch(s string, p string) bool {
	if len(s) == 0 || len(p) == 0 {
		return false
	}

	if p[:1] == "*" {
		return false
	}

	ret := true
	ss := 0
	ps := 0

	for ps < len(p) {
		if (ps+1 < len(p)) && (p[ps+1] == '*') {
			for ss < len(s) && compare(rune(s[ss]), rune(p[ps])) {
				ss++
			}
			ps += 2
		} else {
			ret = ret && compare(rune(s[ss]), rune(p[ps]))
			ss++
			ps++
		}
	}

	if ss != len(s) {
		return false
	}

	return ret
}

func compare(s rune, p rune) bool {
	if p != '.' {
		return p == s
	}
	return true
}

func main() {
	// str1 := "aa"
	// str2 := "a"
	// str1 := "ab"
	// str2 := ".*"
	// str1 := "aa"
	// str2 := "a*"
	// str1 := "aab"
	// str2 := "c*a*b"
	str1 := "mississippi"
	str2 := "mis*is*p*."
	fmt.Println(isMatch(str1, str2))
}
