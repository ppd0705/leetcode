package main

func dfs(node *Node, ret *[]int) {
	if node != nil {
		for _, c := range node.Children {
			dfs(c, ret)
		}
		*ret = append(*ret, node.Val)
	}
}

func postorder2(root *Node) []int {
	ret := make([]int, 0)
	dfs(root, &ret)
	return ret
}
