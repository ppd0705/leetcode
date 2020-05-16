package main

import (
	"reflect"
	"testing"
)

func TestPreOrder(t *testing.T) {
	n1 := Node{Val: 1}
	n2 := Node{Val: 2}
	n3 := Node{Val: 3}
	n4 := Node{Val: 4}
	n5 := Node{Val: 5}
	n1.Children = []*Node{&n2, &n3, &n4}
	n3.Children = []*Node{&n5}

	expect := []int{1, 2, 3, 5, 4}

	if ret := preorder3(&n1); !reflect.DeepEqual(ret, expect) {
		t.Fatalf("expect %v, ret %v", expect, ret)
	}
}
