package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}
func main() {
	tree1 := &TreeNode{
		Val:   10,
		Left:  &TreeNode{4, &TreeNode{}, &TreeNode{}},
		Right: &TreeNode{6, &TreeNode{}, &TreeNode{}},
	}

	fmt.Println(checkTree(tree1))
}
