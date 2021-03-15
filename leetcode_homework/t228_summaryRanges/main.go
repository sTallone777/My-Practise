package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	ret := []string{}
	s := 0

	for i := s; i < len(nums); i = s {
		j := i
		for ; j < len(nums); j++ {
			if nums[j]-nums[s] != j-i {
				break
			}
		}
		if j-1 == s {
			ret = append(ret, strconv.Itoa(nums[i]))
		} else {
			ret = append(ret, strconv.Itoa(nums[s])+"->"+strconv.Itoa(nums[j-1]))
		}
		s = j
	}
	return ret
}

func main() {
	nums := []int{-3, -1, 0, 1}
	fmt.Printf("Summary Range: %v\r\n", summaryRanges(nums))
}
