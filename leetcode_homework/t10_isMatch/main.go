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
	si := 0

	for ps < len(p) {
		if (ps+1 < len(p)) && (p[ps+1] == '*') {
			si = ss
			for ss < len(s) && compare(s, p, ss, ps) {
				ss++
			}
			ps += 2
			if ss == len(s) && ps != len(p) && ss != si {
				if strings.ContainsRune(p[ps:], '*') {
					ss = ss - (strings.IndexRune(p[ps:], '*') - 1)
				} else {
					ss = ss - len(p[ps:])
				}
			}
		} else {
			ret = ret && compare(s, p, ss, ps)
			ss++
			ps++
		}

	}

	if ss != len(s) || ps != len(p) {
		return false
	}

	return ret
}

func compare(s string, p string, ss int, ps int) bool {
	if ss >= len(s) {
		if (ps+1 < len(p)) && (p[ps+1] == '*') {
			return true
		}
		return false
	}

	if p[ps] != '.' {
		return p[ps] == s[ss]
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
