package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	node1 := &TreeNode{Val:1}
	node2 := &TreeNode{Val:2}
	node3 := &TreeNode{Val:3}
	node4 := &TreeNode{Val:4}
	node5 := &TreeNode{Val:5}
	node6 := &TreeNode{Val:6}

	node1.Left = node2
	node1.Right = node3
	node2.Right = node4
	node3.Left = node5
	node5.Right = node6

	cases := []*TreeNode{node1, node2, node3, node3, node4, node5, node6}

	for _, c := range cases {
		obj := Constructor2()
		data := obj.serialize(c)
		ans := obj.deserialize(data)
		if !reflect.DeepEqual(c, ans) {
			t.Fatalf("failed: expect : %+v, but got: %+v\n", c, ans)
		}
	}

}