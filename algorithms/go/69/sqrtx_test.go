package main

import "testing"

func Test(t *testing.T) {
	tests := []struct{ x, expect int }{
		{2, 1},
		{5, 2},
		{8, 2},
		{10, 3},
		{11, 3},
		{12, 3},
		{15, 3},
	}
	for _, tt := range tests {
		ans := mySqrt(tt.x)
		if ans != tt.expect {
			t.Fatalf("x %d expect %d, got %d\n", tt.x, tt.expect, ans)
		}
	}
}
