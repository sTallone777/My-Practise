package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) add(e int) {
	var newNode = new(ListNode)
	newNode.Val = e
	newNode.Next = l.Next
	l.Next = newNode
}

func addTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	var arr1 []int
	var arr2 []int

	for l1 != nil {
		arr1 = append(arr1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		arr2 = append(arr2, l2.Val)
		l2 = l2.Next
	}

	i := 0
	rem := 0
	var nodeArr []ListNode
	for i < len(arr1) && i < len(arr2) {
		sum := arr1[i] + arr2[i] + rem
		node := ListNode{
			Val: sum % 10,
		}
		nodeArr = append(nodeArr, node)
		rem = sum / 10
		i++
	}

	for i < len(arr1) {
		sum := arr1[i] + rem
		node := ListNode{
			Val: sum % 10,
		}
		nodeArr = append(nodeArr, node)
		rem = sum / 10
		i++
	}

	for i < len(arr2) {
		sum := arr2[i] + rem
		node := ListNode{
			Val: arr2[i] + rem,
		}
		nodeArr = append(nodeArr, node)
		rem = sum / 10
		i++
	}

	if rem > 0 {
		node := ListNode{
			Val: rem,
		}
		nodeArr = append(nodeArr, node)
	}

	if nodeArr != nil {
		for i := 0; i < len(nodeArr); i++ {
			if i < len(nodeArr)-1 {
				nodeArr[i].Next = &nodeArr[i+1]
			}
		}
		return &nodeArr[0]
	}

	return &ListNode{
		Val:  0,
		Next: nil,
	}
}

//3 2 5
//1 7 8
// 523 + 871 = 1394
// 4 9 3 1
func main() {
	var l1 = new(ListNode)
	l1.Val = 9
	l1.add(9)
	l1.Next.add(9)
	l1.Next.Next.add(9)
	var l2 = new(ListNode)
	l2.Val = 9
	l2.add(9)
	l2.Next.add(9)
	node := addTwoNumber(l1, l2)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
