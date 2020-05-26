package main

import "testing"

func Test_POW_X_N(t *testing.T) {
	tests := []struct {
		x      float64
		n      int
		expect float64
	}{
		{2, -3, 0.125},
		{2, 3, 8},
		{-2, 1, -2},
		{-2, -2, 0.25},
		{-2, 0, 1},
	}

	for _, tt := range tests {
		if ans := myPow2(tt.x, tt.n); ans != tt.expect {
			t.Fatalf("failed. expect: %+v, got %f\n", tt, ans)
		}
	}
}
