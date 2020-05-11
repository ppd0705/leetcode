package main

func inorderTraversal(root *TreeNode) []int {
	var ret2 []int
	var pre *TreeNode
	curr := root
	for curr != nil {
		if curr.Left == nil {
			ret2 = append(ret2, curr.Val)
			curr = curr.Right
		} else {
			pre = curr.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = curr
			curr = curr.Left
			pre.Right.Left = nil
		}
	}
	return ret2
}
