package main

import "fmt"

func missingNumber(nums []int) int {
	nm := make(map[int]int)
	for i, v := range nums {
		nm[v] = i
	}

	for i := 0; i <= len(nums); i++ {
		if _, ok := nm[i]; ok == false {
			return i
		}
	}
	return 0
}

func main() {
	arr1 := []int{0, 1}
	fmt.Printf("output: %d\r\n", missingNumber(arr1))
}
