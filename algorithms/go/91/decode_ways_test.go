package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		s      string
		target int
	}{
		{"1", 1},
		{"0", 0},
		{"100", 0},
		{"1010", 1},
		{"1213", 5},
	}

	for _, tt := range tests {
		if ans := numDecodings(tt.s); ans != tt.target {
			t.Fatalf("%s expect %d, got %d\n", tt.s, tt.target, ans)
		}
	}
}
