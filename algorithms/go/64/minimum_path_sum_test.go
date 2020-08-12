package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		grid   [][]int
		target int
	}{
		{
			[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7,
		},
	}
	for _, tt := range tests {
		if ans := minPathSum(tt.grid); ans != tt.target {
			t.Fatalf("expect %d, got %d\n", tt.target, ans)
		}
	}
}
