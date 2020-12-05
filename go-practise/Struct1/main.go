package main

import "fmt"

type Rect struct {
	width  int
	height int
}

func (r *Rect) Area() int {
	return r.width * r.height
}

func main() {
	r := Rect{width: 10, height: 5}
	fmt.Println("area: ", r.Area())
}
