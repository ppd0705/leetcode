package main

import "testing"

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree2(preorder, inorder)

	if root.Val != 3 || root.Left.Val != 9 || root.Right.Right.Val != 7 {
		t.Fatal("failed to build tree")
	}
}
