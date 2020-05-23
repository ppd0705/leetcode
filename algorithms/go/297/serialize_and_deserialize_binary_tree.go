package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}


func dfs(valsPtr *[]string) *TreeNode {
	val := (*valsPtr)[0]
	*valsPtr = (*valsPtr)[1:]
	if val == "#" {
		return nil
	}
	valInt, _ := strconv.Atoi(val)
	node := &TreeNode{Val: valInt}
	node.Left = dfs(valsPtr)
	node.Right = dfs(valsPtr)
	return node
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	return dfs(&vals)
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */