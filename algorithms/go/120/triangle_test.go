package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		triangle [][]int
		expect   int
	}{
		{[][]int{
			{2},
			{3, 4},
			{6, 5, 7},
			{4, 1, 8, 3},
		},
			11,
		},
	}
	for _, tt := range tests {
		if ans := minimumTotal(tt.triangle); ans != tt.expect {
			t.Fatalf("expect %d, got %d\n", tt.expect, ans)
		}
	}

}
