package main

func maxDepth2(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	if root == nil {
		return 0
	}
	stack = append(stack, root)
	res := 0
	for len(stack) > 0 {
		l := len(stack)
		for i := 0; i < l; i++ {
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		res++
		stack = stack[l:]
	}
	return res
}
