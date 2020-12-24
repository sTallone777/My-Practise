package main

import (
	"fmt"
	"strconv"
)

//TreeNode :
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	retArr := []int{}
	if root == nil {
		return nil
	}

	retArr = append(retArr, inorderTraversal(root.Left)...)
	retArr = append(retArr, root.Val)
	retArr = append(retArr, inorderTraversal(root.Right)...)

	return retArr
}

//Create treenode by traversal
func createTree(lr string, d int, parent int) *TreeNode {
	var a string

	fmt.Printf("Please input the deep: %d %s node:\r\n", d, lr)
	fmt.Scanln(&a)
	//Stop to creating the branch by typing '#'
	if a == "#" {
		fmt.Printf("Stop traversal on deep: %d\r\n", d)
		return nil
	}
	tree := new(TreeNode)
	tree.Val, _ = strconv.Atoi(a)
	fmt.Printf("Start create deep: %d node: %d.\r\n", d, tree.Val)
	if lr != "root" {
		fmt.Printf("**Current parent node: %d\r\n", parent)
	}
	tree.Left = createTree("left", d+1, tree.Val)
	fmt.Printf("Back to deep: %d node: %d.\r\n", d, tree.Val)
	tree.Right = createTree("right", d+1, tree.Val)
	fmt.Printf("Back to deep: %d node: %d.\r\n", d, tree.Val)

	fmt.Printf("Create deep: %d node: %d successed.\r\n", d, tree.Val)
	return tree
}

func main() {
	//Create root node
	root := createTree("root", 1, -1)

	if root == nil {
		fmt.Println("Create node: root failed.")
	}
	fmt.Println("Create node: root successed.")

	middleTraversal := inorderTraversal(root)
	fmt.Printf("%v", middleTraversal)
}
