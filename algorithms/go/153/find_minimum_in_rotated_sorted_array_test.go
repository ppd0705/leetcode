package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{[]int{1, 2}, 1},
		{[]int{2, 1}, 1},
		{[]int{4, 1, 2, 3}, 1},
		{[]int{4, 5, 1, 2, 3}, 1},
		{[]int{4, 5, 6, 1, 2, 3}, 1},
		{[]int{4, 5, 6, 7, 1, 2, 3}, 1},
		{[]int{4, 5, 6, 7, 1}, 1},
	}

	for _, tt := range tests {
		if ans := findMin(tt.nums); ans != tt.expect {
			t.Fatalf("failed, nums %v, got %d\n", tt.nums, ans)
		}
	}
}
