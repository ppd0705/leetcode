package main

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		coins  []int
		amount int
		expect int
	}{
		{[]int{2}, 1, -1},
		{[]int{2}, 2, 1},
		{[]int{1, 9, 10}, 18, 2},
	}

	for _, tt := range tests {
		if ans := coinChange(tt.coins, tt.amount); ans != tt.expect {
			t.Fatalf("expect %d, got %d\n", tt.expect, ans)
		}
	}
}
