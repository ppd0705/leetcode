package main

func inorderTraversal3(root *TreeNode) []int {
	var ret2 []int
	stack := make([]*TreeNode, 0)

	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		ret2 = append(ret2, curr.Val)
		curr = curr.Right
		stack = stack[:len(stack)-1]
	}
	return ret2
}
