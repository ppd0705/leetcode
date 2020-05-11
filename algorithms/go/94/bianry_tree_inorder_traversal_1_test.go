package main

import (
	"reflect"
	"testing"
)

func TestEmptyTreeNode(t *testing.T) {
	except := make([]int, 0)

	ret = inorderTraversal(nil)

	if reflect.DeepEqual(ret, except) {
		t.Fatalf("failed, ret: %v", ret)
	}
}

func TestTreeNode(t *testing.T) {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node4 := TreeNode{Val: 4}
	node1.Left = &node2
	node1.Right = &node3
	node3.Left = &node4

	except := [4]int{2, 1, 4, 3}

	ret = inorderTraversal(&node1)

	if reflect.DeepEqual(ret, except) {
		t.Fatalf("failed, ret: %v", ret)
	}
}
