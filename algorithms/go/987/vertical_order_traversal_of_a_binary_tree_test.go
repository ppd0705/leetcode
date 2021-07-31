package main

import (
	"reflect"
	"testing"
)

func TestVerticalOrderTraversal(t *testing.T) {
	var res [][]int
	var want [][]int
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 7}

	res = verticalTraversal(node1)
	want = [][]int{{1}}
	if !reflect.DeepEqual(res, want) {
		t.Fatalf("want: %v, got: %v\n", want, res)
	}

	node1.Left = node2
	res = verticalTraversal(node1)
	want = [][]int{{2}, {1}}
	if !reflect.DeepEqual(res, want) {
		t.Fatalf("want: %v, got: %v\n", want, res)
	}

	node1.Right = node3
	node2.Left = node5
	node2.Right = node7
	node3.Right = node6
	node3.Left = node4
	res = verticalTraversal(node1)
	want = [][]int{{5}, {2}, {1, 4, 7}, {3}, {6}}
	if !reflect.DeepEqual(res, want) {
		t.Fatalf("want: %v, got: %v\n", want, res)
	}
}
