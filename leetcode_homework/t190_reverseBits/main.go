package main

import "fmt"

func reverseBits(num uint32) uint32 {
    var ret uint32 = 0
    
	for i := 0; i < 32; i++ {
		ret |= ((num >> i) & 0x01 << (31 - i))
	}
	return ret
}

func main() {
	var num uint32 = 4294967293
	fmt.Printf("reverseBits result: %d\r\n", reverseBits(num))
}
