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

func removeElements(head *ListNode, val int) *ListNode {
	// copy the head
	s := head
	// pre node
	pre := s
	// current node index of ListNode
	c := 0
	for s != nil {
		if s.Val == val {
			// when the delete target node is head
			if c == 0 {
				head = head.Next
				s = head
				pre = s
				c--
			} else {
			// when the delete target node is not head
				pre.Next = s.Next    // delete the target node
				s = s.Next    // move to next node
			}
		} else {
			// move the pointer to next node
			pre = s
			s = s.Next
		}
		c++
	}
	return head
}

func main() {
	s := create(7).add(7).add(7).add(7).add(7)
	// s := create(1)
	s = removeElements(s, 7)
	s.display()
}
