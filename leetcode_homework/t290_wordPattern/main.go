package main

import (
	"fmt"
	"strconv"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	if len(pattern) == 0 || len(s) == 0 {
		return false
	}

	sArr := strings.Split(s, " ")
	if len(sArr) != len(pattern) {
		return false
	}

	judgeMap := make(map[rune]string)

	for k, v := range pattern {
		//when string doesn't match pattern
		if val, ok := judgeMap[v]; ok != false {
			if val != sArr[k] {
				return false
			}
			continue
		}
		//when pattern doesn't match string
		for _, mv := range judgeMap {
			if sArr[k] == mv {
				return false
			}
		}

		judgeMap[v] = sArr[k]
	}

	return true
}

func main() {
	pattern := "aaad"
	s := "cat cat cat cat"
	fmt.Printf("result: %s\r\n", strconv.FormatBool(wordPattern(pattern, s)))
}
