package main

func dfs(node *Node, level int, ret *[][]int) {
	if len(*ret) == level {
		*ret = append(*ret, make([]int, 0))
	}
	(*ret)[level] = append((*ret)[level], node.Val)
	for _, c := range node.Children {
		dfs(c, level+1, ret)
	}
}

func levelOrder3(root *Node) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	dfs(root, 0, &ret)
	return ret
}
