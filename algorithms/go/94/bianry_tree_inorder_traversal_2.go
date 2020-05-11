package main

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret []int

func inorder(node *TreeNode) {
	if node != nil {
		inorder(node.Left)
		ret = append(ret, node.Val)
		inorder(node.Right)
	}
}

func inorderTraversal2(root *TreeNode) []int {
	ret = make([]int, 0)
	inorder(root)
	return ret
}
