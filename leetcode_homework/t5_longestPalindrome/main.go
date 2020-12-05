package main

import "fmt"

func longestPalindrome(s string) string {
	pm := make(map[rune]int)
	retStr := ""
	for i, v := range s {
		if _, ok := pm[v]; ok {
			if len(s[pm[v]:i+1]) > len(retStr) {
				retStr = s[pm[v] : i+1]
			}
		} else {
			pm[v] = i
		}
	}
	if len(retStr) > 1 {
		return retStr
	}
	return ""
}

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}
