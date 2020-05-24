package main

func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	idx := 0
	root := &TreeNode{Val: preorder[0]}
	stack := []*TreeNode{root}

	for i := 1; i < len(preorder); i++ {
		node := stack[len(stack)-1]
		preorderVal := preorder[i]
		if node.Val != inorder[idx] {
			node.Left = &TreeNode{Val: preorderVal}
			stack = append(stack, node.Left)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[idx] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				idx++
			}
			node.Right = &TreeNode{Val: preorderVal}
			stack = append(stack, node.Right)
		}
	}
	return root

}
