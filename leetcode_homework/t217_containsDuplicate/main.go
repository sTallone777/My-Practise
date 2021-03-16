package main

import (
	"fmt"
	"strconv"
)

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for k, v := range nums {
		if _, ok := m[v]; ok {
			return true
		} else {
			m[v] = k
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Printf("contains duplicate: %s\r\n", strconv.FormatBool(containsDuplicate(nums)))
}
