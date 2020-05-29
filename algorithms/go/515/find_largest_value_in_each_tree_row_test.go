package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}

	node1.Left = node2
	node1.Right = node3
	node2.Right = node4
	node4.Left = node5
	node4.Right = node6

	expect := []int{1, 3, 4, 6}
	if ans := largestValues2(node1); !reflect.DeepEqual(ans, expect) {
		t.Fatalf("failed. expect %v, got %v\n", expect, ans)
	}
}
