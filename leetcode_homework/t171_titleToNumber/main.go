package main

import "fmt"

func titleToNumber(columnTitle string) int {
	baseNum := 64
	col := 0
	l := len(columnTitle)
	for i := 0; i < l; i++ {
		s1 := 1
		for j := 0; j < l-1-i; j++ {
			s1 *= 26
		}
		col += (int(rune(columnTitle[i])) - baseNum) * s1
	}
	return col
}

func main() {
	s := "ZYP"
	fmt.Printf("title to number: %d\r\n", titleToNumber(s))
}
