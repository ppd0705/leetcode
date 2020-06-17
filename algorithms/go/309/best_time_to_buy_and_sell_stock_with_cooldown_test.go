package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		prices []int
		expect int
	}{
		{[]int{1,2,3,0,2}, 3},
	}

	for _, tt := range tests {
		if ans := maxProfit3(tt.prices); ans != tt.expect {
			t.Fatalf("expect %d, got %d\n", tt.expect, ans)
		}
	}
}
