package main

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		board  [][]int
		expect int
	}{
		{[][]int{{1, 2, 3}, {4, 0, 5}}, 1},
	}

	for _, tt := range tests {
		if ans := slidingPuzzle(tt.board); ans != tt.expect {
			t.Fatalf("expect %d, got %d\n", tt.expect, ans)
		}
	}

}
