package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		cost   []int
		expect int
	}{
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
		{[]int{10, 15, 20}, 15},
	}

	for _, tt := range tests {
		if ans := minCostClimbingStairs(tt.cost); ans != tt.expect {
			t.Fatalf("expect %d, got %d\n", tt.expect, tt.cost)
		}
	}
}
