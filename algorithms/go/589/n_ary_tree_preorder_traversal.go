package main

type Node struct {
	Val int
	Children []*Node
}

func preorder(root *Node) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}

	stack := make([]*Node, 0)
	stack = append(stack, root)

	var node *Node
	var idx int
	for len(stack) > 0 {
		idx = len(stack) - 1
		node = stack[idx]
		stack = stack[:idx]
		ret = append(ret, node.Val)

		for i := len(node.Children)-1; i >=0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return ret
}
