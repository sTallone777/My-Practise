package main

import (
	"fmt"
	"strconv"
)

func isHappy(n int) bool {
	m := map[int]bool{}
	sn := n
	for sn != 1 && !m[sn] {
		m[sn] = true
		sn = step(sn)
	}
	return sn == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

func main() {
	num := 12394
	fmt.Printf("is happy: %s\r\n", strconv.FormatBool(isHappy(num)))
}
