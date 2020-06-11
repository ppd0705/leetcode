package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}

	for _, tt := range tests {
		if ans := maxSubArray2(tt.nums); ans != tt.expect {
			t.Fatalf("expect %d, got %d\n", tt.expect, ans)
		}
	}
}
