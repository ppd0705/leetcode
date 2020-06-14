package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
	}

	for _, tt := range tests {
		if ans := lengthOfLIS2(tt.nums); ans != tt.expect {
			t.Fatalf("expect %d, got %d\n", tt.expect, ans)
		}
	}
}
