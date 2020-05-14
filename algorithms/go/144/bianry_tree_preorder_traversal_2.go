package main

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret0 []int

func preorder(node *TreeNode) {
	if node != nil {
		ret0 = append(ret0, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
}

func preorderTraversal2(root *TreeNode) []int {
	ret0 = make([]int, 0)
	preorder(root)
	return ret0
}
