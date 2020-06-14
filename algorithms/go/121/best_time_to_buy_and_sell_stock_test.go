package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, tt := range tests {
		if ans := maxProfit(tt.nums); ans != tt.expect {
			t.Fatalf("expect %d, got %d\n", tt.expect, ans)
		}
	}
}
