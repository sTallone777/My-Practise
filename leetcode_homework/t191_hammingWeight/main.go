package main

import "fmt"

func hammingWeight(num uint32) int {
	ret := 0
	for i := 0; i < 32; i++ {
		if (num>>i)&1 > 0 {
			ret++
		}
	}
	return ret
}

func main() {
	var num uint32 = 13
	fmt.Printf("hammingweight: %d\r\n", hammingWeight(num))
}
