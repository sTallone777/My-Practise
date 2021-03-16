package main

import (
	"fmt"
	"strconv"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[v]; ok {
			if i-m[v] <= k {
				return true
			}
			m[v] = i
		} else {
			m[v] = i
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	fmt.Printf("contains nearby duplicate: %s\r\n", strconv.FormatBool(containsNearbyDuplicate(nums, k)))
}
