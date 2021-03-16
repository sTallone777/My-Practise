package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for numbers[i]+numbers[j] != target {
		if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}
	return []int{i + 1, j + 1}
}

func main() {
	numbers := []int{2, 3, 4}
	target := 6
	fmt.Printf("two sum: %v\r\n", twoSum(numbers, target))
}
