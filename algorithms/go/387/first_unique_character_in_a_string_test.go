package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		s      string
		target int
	}{
		{"leetcode", 0},
		{"leetlover", 3},
	}

	for _, tt := range tests {
		if ans := firstUniqChar(tt.s); ans != tt.target {
			t.Fatalf("expect %d, got %d\n", tt.target, ans)
		}
	}
}
