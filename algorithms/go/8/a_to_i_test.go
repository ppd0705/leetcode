package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		s string
		i int
	}{
		{" -42", -42},
		{"a -42", 0},
		{" -4223w", -4223},
		{"+1", 1},
		{"-2147483649", -2147483648},
		{"2147483649", 2147483647},
	}

	for _, tt := range tests {
		if ans := myAtoi(tt.s); ans != tt.i {
			t.Fatalf("target %d, got %d", tt.i, ans)
		}
	}
}
