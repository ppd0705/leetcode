package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect bool
	}{
		{[]int{1, 3, 1, 1, 1}, 3, true},
		{[]int{1, 1, 1, 3, 1}, 3, true},
		{[]int{1, 2, 2, 3, 1}, 3, true},
		{[]int{1, 2, 2, 3, 4}, 3, true},
	}

	for _, tt := range tests {
		if ans := search(tt.nums, tt.target); ans != tt.expect {
			t.Fatalf("nums %v, target %d\n", tt.nums, tt.target)
		}
	}
}
