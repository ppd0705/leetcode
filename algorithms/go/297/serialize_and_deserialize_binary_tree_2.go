package main

import (
	"strconv"
	"strings"
)

type Codec2 struct {
}

func Constructor2() Codec2 {
	return Codec2{}
}

// Serializes a tree to a single string.
func (this *Codec2) serialize(root *TreeNode) string {
	stack := []*TreeNode{root}
	vals := make([]string, 0)
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		if node == nil {
			vals = append(vals, "#")
		} else {
			vals = append(vals, strconv.Itoa(node.Val))
			stack = append(stack, node.Left, node.Right)
		}
	}
	return strings.Join(vals, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec2) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	idx := 0
	if vals[idx] == "#" {
		return nil
	}
	val, _ := strconv.Atoi(vals[idx])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}
	var node, left, right *TreeNode
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		idx++
		if vals[idx] != "#" {
			val, _ = strconv.Atoi(vals[idx])
			left = &TreeNode{Val: val}
			node.Left = left
			queue = append(queue, left)
		}
		idx++
		if vals[idx] != "#" {
			val, _ = strconv.Atoi(vals[idx])
			right = &TreeNode{Val: val}
			node.Right = right
			queue = append(queue, right)
		}
	}
	return root
}
