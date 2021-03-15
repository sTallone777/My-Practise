package main

import (
	"fmt"
	"strconv"
)

type MyQueue struct {
	queue       []int
	front, rear int
}

func Constructor() MyQueue {
	return MyQueue{
		queue: []int{},
		front: -1,
		rear:  -1,
	}
}

func (this *MyQueue) Push(x int) {
	this.queue = append(this.queue, x)
	if this.front < 0 {
		this.front++
	}
	this.rear++
}

func (this *MyQueue) Pop() int {
	ret := this.queue[this.front]
	this.front++
	return ret
}

func (this *MyQueue) Peek() int {
	return this.queue[this.front]
}

func (this *MyQueue) Empty() bool {
	if this.rear == -1 || this.rear < this.front {
		return true
	}
	return false
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Printf("Peek: %d\r\n", obj.Peek())
	fmt.Printf("Pop: %d\r\n", obj.Pop())
	fmt.Printf("Empty: %s\r\n", strconv.FormatBool(obj.Empty()))
}
