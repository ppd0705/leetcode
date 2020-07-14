package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		n      int
		target bool
	}{
		{1, true},
		{0, false},
		{3, false},
		{16, true},
	}
	for _, tt := range tests {
		if ans := isPowerOfTwo(tt.n); ans != tt.target {
			t.Fatalf("failed to calculate %d\n", tt.n)
		}
	}
}
