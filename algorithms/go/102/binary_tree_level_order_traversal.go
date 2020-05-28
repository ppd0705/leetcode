package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		level := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[i]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, level)
		queue = queue[l:]
	}

	return ans
}

func dfs(node *TreeNode, level int, ans *[][]int) {
	if node == nil {
		return
	}
	if len(*ans) == level {
		*ans = append(*ans, []int{})
	}
	(*ans)[level] = append((*ans)[level], node.Val)

	dfs(node.Left, level+1, ans)
	dfs(node.Right, level+1, ans)
}
func levelOrder2(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	dfs(root, 0, &ans)
	return ans
}
