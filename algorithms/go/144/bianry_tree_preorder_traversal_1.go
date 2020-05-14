package main

func preorderTraversal(root *TreeNode) []int {
	var ret []int
	var pre *TreeNode
	curr := root
	for curr != nil {
		ret = append(ret, curr.Val)
		if curr.Left == nil {
			curr = curr.Right
		} else {
			pre = curr.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = curr.Right
			curr = curr.Left
		}
	}
	return ret
}
