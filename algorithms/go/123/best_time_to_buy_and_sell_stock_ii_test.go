package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		prices []int
		expect int
	}{
		{[]int{3, 3, 5, 0, 0, 3, 1, 4}, 6},
	}

	for _, tt := range tests {
		if ans := maxProfit2(tt.prices); ans != tt.expect {
			t.Fatalf("expect %d, got %d\n", tt.expect, ans)
		}
	}
}
