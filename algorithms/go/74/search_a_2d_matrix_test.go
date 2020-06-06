package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		expect bool
	}{
		{[][]int{{1, 3}}, 3, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 3, true},
	}

	for _, tt := range tests {
		if ans := searchMatrix(tt.matrix, tt.target); ans != tt.expect {
			t.Fatalf("tartget %d, expect %v\n", tt.target, tt.expect)
		}
	}
}
