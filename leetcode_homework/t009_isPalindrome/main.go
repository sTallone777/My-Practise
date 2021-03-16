package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	i, j := 0, len(s)-1
	for i != j && i < j  {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	x := -121
	fmt.Printf("is palindrome: %s\r\n", strconv.FormatBool(isPalindrome(x)))
}
