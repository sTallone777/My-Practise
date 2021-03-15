package main

import (
	"fmt"
	"strconv"
)

type MyStack struct {
	queue       []int
	top			int
}

func Constructor() MyStack {
	return MyStack{
		queue: []int{},
		top:   -1,
	}
}

func (this *MyStack) Push(x int) {
	this.top++
	if len(this.queue) == this.top {
		this.queue = append(this.queue, x)
	} else {
		this.queue[this.top] = x
	}
}

func (this *MyStack) Pop() int {
	ret := this.queue[this.top]
	this.top--
	return ret
}

func (this *MyStack) Top() int {
	return this.queue[this.top]
}

func (this *MyStack) Empty() bool {
	if this.top < 0 {
		return true
	}
	return false
}

func main() {
	obj := Constructor()
	obj.Push(1)

	fmt.Printf("Pop: %d\r\n", obj.Pop())
	obj.Push(2)
	fmt.Printf("Top: %d\r\n", obj.Top())
	fmt.Printf("Pop: %d\r\n", obj.Pop())
	fmt.Printf("Empty: %s\r\n", strconv.FormatBool(obj.Empty()))
}
