package main

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Right = node6

	tests := []struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		{node1, node2, node3, node1},
		{node1, node2, node6, node1},
		{node1, node2, node4, node2},
		{node1, node3, node6, node3},
		{node1, node1, node5, node1},
	}
	for _, tt := range tests {
		if ans := lowestCommonAncestor2(tt.root, tt.p, tt.q); ans != tt.want {
			t.Fatalf("want: %+v, but got %+v\n", tt.want, ans)
		}
	}
}
