package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, lower int, upper int) bool {
	if node == nil {
		return true
	}
	return lower < node.Val && node.Val < upper && dfs(node.Left, lower, node.Val) &&
		dfs(node.Right, node.Val, upper)
}
func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt64, math.MaxInt64)
}
