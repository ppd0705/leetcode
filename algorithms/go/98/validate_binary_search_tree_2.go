package main

import "math"

func isValidBST2(root *TreeNode) bool {
	node := root
	stack := make([]*TreeNode, 0)
	pre := math.MinInt64
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		idx := len(stack) - 1
		node = stack[idx]
		if node.Val <= pre {
			return false
		}
		pre = node.Val
		stack = stack[:idx]
		node = node.Right
	}
	return true
}
