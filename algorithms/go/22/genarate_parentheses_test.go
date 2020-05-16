package main

import (
	"reflect"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	expect := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}

	if res := generateParenthesis(3); !reflect.DeepEqual(res, expect) {
		t.Fatalf("res: %v, expecet: %v\n", res, expect)
	}
}
