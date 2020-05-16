package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	var node *Node
	var val []int
	level := []*Node{root}
	for len(level) > 0 {
		val = make([]int, 0)
		length := len(level)
		for _, node = range level {
			val = append(val, node.Val)
			for i:= 0; i < len(node.Children); i++ {
				level = append(level, node.Children[i])
			}
		}
		ret = append(ret, val)
		level = level[length:]
	}
	return ret
}
