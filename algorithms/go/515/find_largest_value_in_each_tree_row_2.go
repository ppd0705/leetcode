package main

func largestValues2(root *TreeNode) []int {
	ans := make([]int, 0)
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(ans) == level {
			ans = append(ans, node.Val)
		} else if ans[level] < node.Val {
			ans[level] = node.Val
		}

		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)
	return ans
}
