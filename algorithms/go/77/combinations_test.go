package main

import (
	"reflect"
	"testing"
)

func TestCombination(t *testing.T) {
	ans := combine(5, 4)
	expect := [][]int{
		{1, 2, 3, 4},
		{1, 2, 3, 5},
		{1, 2, 4, 5},
		{1, 3, 4, 5},
		{2, 3, 4, 5},
	}
	if !reflect.DeepEqual(ans, expect) {
		t.Fatalf("failed: expect: %v but gotans: %v\n", expect, ans)
	}
}
