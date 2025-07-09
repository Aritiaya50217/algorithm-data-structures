package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insert(node *TreeNode, newData int) *TreeNode {
	if node == nil {
		return &TreeNode{
			Val:   newData,
			Left:  nil,
			Right: nil,
		}
	}
	return node
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	left := inorderTraversal(root.Left)
	current := []int{root.Val}
	right := inorderTraversal(root.Right)

	result := append(append(left, current...), right...)
	return result
}

func main() {
	var root *TreeNode
	root = insert(root, 1)
	root.Right = insert(root.Right, 2)
	root.Right.Left = insert(root.Right.Left, 3)

	res := inorderTraversal(root)
	fmt.Println(res)
}
