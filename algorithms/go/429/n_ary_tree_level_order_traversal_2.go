package main

import "container/list"

func levelOrder2(root *Node) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		length := queue.Len()
		val := make([]int, 0)
		for i := 0; i < length; i++ {
			ele := queue.Front()
			queue.Remove(ele)
			node := ele.Value.(*Node)
			val = append(val, node.Val)
			for _, c :=range node.Children {
				queue.PushBack(c)
			}
		}
		ret = append(ret, val)
	}

	return ret
}
