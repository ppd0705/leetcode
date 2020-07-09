package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		grid   [][]int
		expect int
	}{
		{[][]int{{0, 1}, {1, 0}}, 2},
		{[][]int{
			{0, 0, 0}, {1, 1, 0}, {1, 1, 0},
		}, 4},
	}

	for _, tt := range tests {
		if ans := shortestPathBinaryMatrix(tt.grid); ans != tt.expect {
			t.Fatalf("expect %d, got %d\n", tt.expect, ans)
		}
	}

}
