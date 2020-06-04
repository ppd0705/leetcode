package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		expect bool
	}{
		{[]int{}, true},
		{[]int{2, 0, 0}, true},
		{[]int{1, 2, 3}, true},
		{[]int{3, 2, 1, 0, 4}, false},
	}
	for _, tt := range tests {
		if ans := canJump(tt.nums); ans != tt.expect {
			t.Fatalf("expect %t, got %t\n", tt.expect, ans)
		}
	}
}
