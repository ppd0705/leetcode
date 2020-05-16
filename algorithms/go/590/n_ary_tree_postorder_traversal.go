package main

type Node struct {
	Val      int
	Children []*Node
}

func reverse(ints []int) {
	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1{
		ints[i], ints[j] = ints[j], ints[i]
	}
}
func postorder(root *Node) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	stack := make([]*Node, 0)
	stack = append(stack, root)
	idx := 0
	for len(stack) > 0 {
		idx = len(stack) - 1
		root = stack[idx]
		stack = stack[:idx]
		ret = append(ret, root.Val)
		for idx = 0; idx < len(root.Children); idx++ {
			stack = append(stack, root.Children[idx])
		}
	}
	reverse(ret)
	return ret
}
