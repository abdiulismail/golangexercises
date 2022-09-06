package main

import "fmt"

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func invertTree(root *TreeNode2) *TreeNode2 {
	if root == nil {
		return nil
	}
	tmp := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(tmp)
	return root
}

func main() {
	tree1 := &TreeNode2{
		Val:   10,
		Left:  &TreeNode2{4, &TreeNode2{}, &TreeNode2{}},
		Right: &TreeNode2{6, &TreeNode2{}, &TreeNode2{}},
	}

	fmt.Println("original tree", tree1.Val, tree1.Left.Val, tree1.Right.Val)

	mytree := invertTree(tree1)
	fmt.Println("inverted tree", mytree.Val, mytree.Left.Val, mytree.Right.Val)
}
