package main

import (
	"crypto/sha256"
	"fmt"
)

//比较两个sha256生成的哈希码的不同bit数
func main() {
	c1 := sha256.Sum256([]byte("y"))
	c2 := sha256.Sum256([]byte("Y"))

	fmt.Printf("%x\n%x\n%t\n%T\n%T\n", c1, c2, c1 == c2, c1, c2)

	count := 0

	for i := 0; i < len(c1); i++ {

		b := c1[i] ^ c2[i]

		for j := 0; j < 32; j++ {
			if b>>j&1 > 0 {
				count++
			}
		}
	}

	fmt.Printf("%d", count)
}
