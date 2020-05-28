package main

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}

	node1.Left = node2
	node1.Right = node3
	node3.Left = node4
	node4.Right = node5

	expect := [][]int{
		{1},
		{2, 3},
		{4},
		{5},
	}

	if ans := levelOrder2(node1); !reflect.DeepEqual(ans, expect) {
		t.Fatalf("failed. expect: %+v, ans: %+v\n", expect, ans)
	}
}
