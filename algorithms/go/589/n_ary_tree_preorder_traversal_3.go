package main

type Stack struct {
	Val  *Node
	Next *Stack
}

func preorder3(root *Node) []int {
	var ret []int
	if root == nil {
		return ret
	}

	stack := &Stack{Val: root}

	for s := stack; s != nil; s = s.Next {
		ret = append(ret, s.Val.Val)

		next := s.Next

		ss := s
		for _, c := range s.Val.Children {
			vv := &Stack{Val: c}
			ss.Next = vv
			ss = vv
		}
		ss.Next = next
	}
	return ret
}
