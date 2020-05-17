package main

func dfs(node *TreeNode, level int) int {
	if node == nil {
		return level
	}
	return maxInt(dfs(node.Left, level+1), dfs(node.Right, level+1))
}
func maxDepth4(root *TreeNode) int {
	return dfs(root, 0)
}
