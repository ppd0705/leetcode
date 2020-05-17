package main

type pair struct {
	depth int
	node  *TreeNode
}

func maxDepth3(root *TreeNode) int {
	stack := make([]*pair, 0)
	if root == nil {
		return 0
	}
	stack = append(stack, &pair{depth: 1, node: root})
	res := 0
	for len(stack) > 0 {
		node, depth := stack[len(stack)-1].node, stack[len(stack)-1].depth
		stack = stack[:len(stack)-1]
		if node != nil {
			if depth > res {
				res = depth
			}
			stack = append(stack, &pair{depth: depth + 1, node: node.Left})
			stack = append(stack, &pair{depth: depth + 1, node: node.Right})
		}
	}
	return res
}
