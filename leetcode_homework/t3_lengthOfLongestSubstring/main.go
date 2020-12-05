package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	max := 0
	ss := 0
	k := 0
	if s != "" || len(s) > 0 {
		for i := 0; i <= len(s); i++ {
			str := s[ss:i]

			if i == len(s) {
				k = i - 1
			} else {
				k = i
			}

			p := strings.LastIndexByte(str, s[k])
			if p > -1 {
				ss += p + 1
				if len(str) > max {
					max = len(str)
				}
			}
		}
	}
	return max
}

func main() {
	data, _ := ioutil.ReadFile("C:/Users/nttcom/Downloads/test1.java")
	num := lengthOfLongestSubstring(string(data))

	// str := "aa"
	// num := lengthOfLongestSubstring(str)
	fmt.Println("The length is :", num)
	// str := "asdf/x."
	// fmt.Println(str[0:0])

}
