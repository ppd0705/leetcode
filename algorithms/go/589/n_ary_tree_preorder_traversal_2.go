package main

func preOrder(node *Node, ret *[]int) {
	if node != nil {
		*ret = append(*ret, node.Val)
		for _, child := range node.Children {
			preOrder(child, ret)
		}
	}

}

func preorder2(root *Node) []int {
	var ret []int
	preOrder(root, &ret)
	return ret
}
