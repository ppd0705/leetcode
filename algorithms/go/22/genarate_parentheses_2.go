package main

type Node struct {
	Val   string
	Left  int
	Right int
}

func generateParenthesis2(n int) []string {
	res := make([]string, 0)
	if n <= 0 {
		return res
	}
	queue := []*Node{{Val: "", Left: n, Right: n}}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Right == 0 {
			res = append(res, node.Val)
		} else {
			if node.Left > 0 {
				queue = append(queue, &Node{Val: node.Val + "(", Left: node.Left - 1, Right: node.Right})
			}
			if node.Right > node.Left {
				queue = append(queue, &Node{Val: node.Val + ")", Left: node.Left, Right: node.Right - 1})
			}
		}
	}
	return res
}
