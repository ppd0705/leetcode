package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{[]int{-2, 3, 5, -4}, 120},
		{[]int{0, 9, 4, -2}, 36},
		{[]int{-2, 9, 4, 0}, 36},
	}

	for _, tt := range tests {
		if ans := maxProduct(tt.nums); ans != tt.expect {
			t.Fatalf("expect %d, got %d\n", tt.expect, ans)
		}
	}
}
