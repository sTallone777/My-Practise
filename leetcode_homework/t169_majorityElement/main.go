package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Printf("majority elements: %d\r\n", majorityElement(nums))
}
