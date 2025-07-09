package main

import (
	"fmt"
	"testing"
)

var root *TreeNode

func TestInorderTraversal(t *testing.T) {
	root = insert(root, 1)
	root.Right = insert(root.Right, 2)
	root.Right.Left = insert(root.Right.Left, 3)

	result := inorderTraversal(root)
	fmt.Println(result)
}

func BenchmarkInorderTraversal(b *testing.B) {
	root = insert(root, 1)
	root.Right = insert(root.Right, 2)
	root.Right.Left = insert(root.Right.Left, 3)

	for i := 0; i < b.N; i++ {
		_ = inorderTraversal(root)
	}
}
