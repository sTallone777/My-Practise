package main

import "fmt"

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	retStr := ""

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] && len(s[i:j+1]) > len(retStr) {
				tmpStr := s[i : j+1]
				isOk := true
				for k := 0; k < len(tmpStr)>>1; k++ {
					if tmpStr[k] != tmpStr[len(tmpStr)-1-k] {
						isOk = false
						break
					}
				}
				if isOk {
					retStr = tmpStr
				}
			}
		}
	}

	if retStr == "" {
		return string(s[0])
	}
	return retStr
}

func main() {
	// fmt.Println(longestPalindrome("aacabdkacaa"))
	// fmt.Println(longestPalindrome("ccc"))
	// fmt.Println(longestPalindrome("babac"))
	fmt.Println(longestPalindrome("c"))
	// fmt.Println(longestPalindrome("abbcccbbbcaaccbababcbcabca"))
}
