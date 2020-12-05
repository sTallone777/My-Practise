package main

import "fmt"

func main() {
	var b int = 10
	//p1 := &b
	fmt.Printf("变量的地址: %x\n", &b)
	fmt.Printf("变量的地址: %x\n", b)
}
