package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	buf.WriteByte('{')
	fmt.Printf("%d \n", buf.Len())

}
