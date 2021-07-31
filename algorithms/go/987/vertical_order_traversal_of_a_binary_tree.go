package main

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Point struct {
	row int
	col int
}

var pointValMap map[Point][]int

func dfs(node *TreeNode, row, col int) {
	if node == nil {
		return
	}
	p := Point{row, col}
	pointValMap[p] = append(pointValMap[p], node.Val)
	dfs(node.Left, row+1, col-1)
	dfs(node.Right, row+1, col+1)
}
func verticalTraversal(root *TreeNode) [][]int {
	var res [][]int
	var points []Point
	pointValMap = map[Point][]int{}

	dfs(root, 0, 0)

	for point, vals := range pointValMap {
		points = append(points, point)
		sort.Ints(vals)
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i].col < points[j].col || (points[i].col == points[j].col && points[i].row < points[j].row) {
			return true
		}
		return false
	})

	prevCol := -10002

	for _, point := range points {
		if res == nil || point.col != prevCol {
			res = append(res, pointValMap[point])
			prevCol = point.col
		} else {
			res[len(res)-1] = append(res[len(res)-1], pointValMap[point]...)
		}
	}
	return res
}
