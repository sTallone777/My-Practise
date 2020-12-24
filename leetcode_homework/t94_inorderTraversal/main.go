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

//Get node value
func (t *TreeNode) getValue() int {
	return t.Val
}

//Set node value
func (t *TreeNode) setValue(v int) {
	t.Val = v
}

//Search node by value
func (t *TreeNode) findNode(root *TreeNode, v int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.getValue() == v {
		return root
	}
	node := t.findNode(root.Left, v)
	if node != nil {
		return node
	}
	return t.findNode(root.Right, v)
}

//Calculate tree height
func (t *TreeNode) getTreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := t.getTreeHeight(root.Left)
	rh := t.getTreeHeight(root.Right)
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}

//Print leafnode
func (t *TreeNode) printLeafNode(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		fmt.Printf("%d\r\n", root.Val)
	}
	t.printLeafNode(root.Left)
	t.printLeafNode(root.Right)
}

//Print treenode
func (t *TreeNode) printTreeNode() {
	if t == nil {
		return
	}
	fmt.Printf("Node Value: %d\r\n", t.Val)
	fmt.Printf("Left Node: %p\r\n", t.Left)
	fmt.Printf("Right Node: %p\r\n", t.Right)
}

//pre-order traversal
func preorderTraversal(root *TreeNode) []int {
	retArr := []int{}
	if root == nil {
		return nil
	}
	retArr = append(retArr, root.Val)
	retArr = append(retArr, preorderTraversal(root.Left)...)
	retArr = append(retArr, preorderTraversal(root.Right)...)

	return retArr
}

//post-order tarversal
func postorderTraversal(root *TreeNode) []int {
	retArr := []int{}
	if root == nil {
		return nil
	}

	retArr = append(retArr, postorderTraversal(root.Left)...)
	retArr = append(retArr, postorderTraversal(root.Right)...)
	retArr = append(retArr, root.Val)

	return retArr
}

//in-order traversal
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

	if lr == "root" {
		fmt.Printf("Please input the deep: %d %s node of root:\r\n", d, lr)
	} else {
		fmt.Printf("Please input the deep: %d %s node of node: %d:\r\n", d, lr, parent)
	}

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

	preorder := preorderTraversal(root)
	fmt.Printf("Pre-order traversal: %v\r\n", preorder)

	inorder := inorderTraversal(root)
	fmt.Printf("In-order traversal: %v\r\n", inorder)

	postorder := postorderTraversal(root)
	fmt.Printf("Post-order traversal: %v\r\n", postorder)

	node1 := root.findNode(root, 3)
	node1.printTreeNode()
	fmt.Printf("This tree's height is : %d\r\n", root.getTreeHeight(node1))
	node2 := root.findNode(root, 5)
	node2.printTreeNode()
	fmt.Printf("This tree's height is : %d\r\n", root.getTreeHeight(node2))

	fmt.Println("Print LeafNode: ")
	root.printLeafNode(root)
}
