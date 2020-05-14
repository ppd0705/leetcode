package main

func preorderTraversal3(root *TreeNode) []int {
	var ret2 []int
	stack := make([]*TreeNode, 0)

	curr := root
	for curr != nil || len(stack) > 0 {
		if curr == nil {
			idx := len(stack) - 1
			curr = stack[idx]
			stack = stack[:idx]
		} else {
			ret2 = append(ret2, curr.Val)
			stack = append(stack, curr.Right)
			curr = curr.Left
		}

	}
	return ret2
}
