package main

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	parent := make(map[int]*TreeNode)
	visted := make(map[int]bool)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			parent[node.Left.Val] = node
			dfs(node.Left)
		}
		if node.Right != nil {
			parent[node.Right.Val] = node
			dfs(node.Right)
		}
	}
	dfs(root)
	for p != nil {
		visted[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visted[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}
