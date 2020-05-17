package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxInt(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + maxInt(maxDepth(root.Left), maxDepth(root.Right))

}
