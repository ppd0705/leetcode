package main

import (
"reflect"
"testing"
)

func TestLevelOrder(t *testing.T) {
	n1 := Node{Val: 1}
	n2 := Node{Val: 2}
	n3 := Node{Val: 3}
	n4 := Node{Val: 4}
	n5 := Node{Val: 5}
	n6 := Node{Val: 6}
	n7 := Node{Val: 7}
	n1.Children = []*Node{&n2, &n3, &n4}
	n3.Children = []*Node{&n5}
	n5.Children = []*Node{&n6}
	n4.Children = []*Node{&n7}
	expect := [][]int{{1}, {2, 3, 4}, {5, 7}, {6}}

	if ret := levelOrder3(&n1); !reflect.DeepEqual(ret, expect) {
		t.Fatalf("expect %v, ret %v", expect, ret)
	}
}
