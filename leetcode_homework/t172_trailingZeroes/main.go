package main

import (
	"fmt"
	"math/big"
)

func trailingZeroes(n int) int {
	s := big.NewInt(1)
	for i := 1; i <= n; i++ {
		s.Mul(s, big.NewInt(int64(i)))
	}
	str := s.String()
	c := 0
	for i := len(str) - 1; i >= 0; i-- {
		if rune(str[i]) != '0' {
			break
		}
		c++
	}
	return c
}

func main() {
	n := 5
	fmt.Printf("trailing zeroes: %d\r\n", trailingZeroes(n))
}
