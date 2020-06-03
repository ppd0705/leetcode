package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		g      []int
		s      []int
		expect int
	}{
		{[]int{1, 2, 3}, []int{2, 3}, 2},
		{[]int{1, 2, 3}, []int{1, 1}, 1},
	}
	for _, tt := range tests {
		if ans := findContentChildren(tt.g, tt.s); ans != tt.expect {
			t.Fatalf("expect %d, got %d\n", tt.expect, ans)
		}
	}
}
