package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
	}{
		{[]int{5, 4, 3, 2, 1}, 4},
		{[]int{1, 3, 2, 3, 1}, 2},
		{[]int{2, 4, 3, 5, 1}, 3},
	}
	for _, tt := range tests {
		if ans := reversePairs(tt.nums); ans != tt.target {
			t.Fatalf("expect %d go %d\n", tt.target, ans)
		}
	}
}
