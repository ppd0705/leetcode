package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect int
	}{
		{[]int{4, 5, 6, 7, 8, 1, 2, 3}, 8, 4},
		{[]int{5, 1, 3, 4}, 5, 0},
		{[]int{5, 1, 3, 4}, 1, 1},
		{[]int{5, 1, 3, 4}, 2, -1},
		{[]int{5, 1, 3, 4}, 3, 2},
		{[]int{5, 1, 3, 4}, 4, 3},
		{[]int{5, 1, 3, 4}, 6, -1},
		{[]int{5, 1, 3, 4}, 0, -1},
		{[]int{1, 3, 4}, 0, -1},
		{[]int{1, 3, 4}, 1, 0},
		{[]int{1, 3, 4}, 3, 1},
		{[]int{1, 3, 4}, 4, 2},
		{[]int{1, 3, 4}, 5, -1},
		{[]int{1, 3, 4}, 2, -1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
	}

	for _, tt := range tests {
		if ans := search(tt.nums, tt.target); ans != tt.expect {
			t.Fatalf("target %d, expect %d, ans %d\n", tt.target, tt.expect, ans)
		}
	}
}
