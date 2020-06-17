package main

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		prices []int
		fee    int
		expect int
	}{
		{[]int{1, 3, 2, 8, 4, 9}, 2, 8},
	}

	for _, tt := range tests {
		if ans := maxProfit(tt.prices, tt.fee); ans != tt.expect {
			t.Fatalf("expect %d, got %d \n", tt.expect, ans)
		}
	}
}
