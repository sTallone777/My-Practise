package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func create(e int) *ListNode {
	return &ListNode{
		Val:  e,
		Next: nil,
	}
}

func (ln *ListNode) add(e int) *ListNode {
	newNode := ListNode{
		Val:  e,
		Next: nil,
	}

	tmp := ln

	for tmp.Next != nil {
		tmp = tmp.Next
	}

	tmp.Next = &newNode
	return ln
}

func (ln *ListNode) display() {
	tmp := ln
	fmt.Print("ListNodes:")
	for tmp != nil {
		fmt.Printf(" %d", tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	s := create(1).add(2).add(3).add(4).add(5)
	s.display()
	s = reverseList(s)
	s.display()
}
