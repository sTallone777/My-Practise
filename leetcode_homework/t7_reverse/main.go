package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var xAbs int

	isMinu := x < 0

	if isMinu {
		xAbs = -x
	} else {
		xAbs = x
	}

	var nums []int

	for xAbs > 0 {
		nums = append(nums, xAbs%10)
		xAbs = xAbs / 10
	}

	var retNum int = 0

	for i, v := range nums {
		retNum += v * int(math.Pow10(len(nums)-1-i))
	}

	if isMinu {
		if retNum > 0x7fffffff {
			return 0
		}
		return -retNum
	}

	if retNum > 0x7ffffffe {
		return 0
	}

	return retNum
}

func main() {
	// fmt.Printf("result is : %d \n", reverse(-123))
	// fmt.Printf("result is : %d \n", reverse(123))
	// fmt.Printf("result is : %d \n", reverse(0xfffffffff))
	fmt.Printf("result is : %d \n", reverse(1534236469))
}
