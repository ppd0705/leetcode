package main

import (
	"reflect"
	"testing"
)

func TestEmptyTreeNode(t *testing.T) {
	var except []int

	ret = preorderTraversal(nil)

	if !reflect.DeepEqual(ret, except) {
		t.Fatalf("failed, ret: %v, except: %v", ret, except)
	}
}

func TestTreeNode1(t *testing.T) {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node4 := TreeNode{Val: 4}
	node1.Left = &node2
	node1.Right = &node3
	node3.Left = &node4

	except := []int{1, 2, 3, 4}

	ret = preorderTraversal(&node1)

	if !reflect.DeepEqual(ret, except) {
		t.Fatalf("failed, ret: %v", ret)
	}
}

func TestTreeNode2(t *testing.T) {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node4 := TreeNode{Val: 4}
	node1.Left = &node2
	node2.Left = &node3
	node3.Left = &node4

	except := []int{1, 2, 3, 4}

	ret = preorderTraversal(&node1)

	if !reflect.DeepEqual(ret, except) {
		t.Fatalf("failed, ret: %v", ret)
	}
}

func TestTreeNode3(t *testing.T) {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node4 := TreeNode{Val: 4}
	node1.Right = &node2
	node2.Right = &node3
	node3.Right = &node4

	except := []int{1, 2, 3, 4}

	ret = preorderTraversal(&node1)

	if !reflect.DeepEqual(ret, except) {
		t.Fatalf("failed, ret: %v", ret)
	}
}

