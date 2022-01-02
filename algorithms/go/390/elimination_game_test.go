package main

import "testing"

func TestLastRemaining(t *testing.T) {
	args := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 2},
		{5, 2},
		{6, 4},
		{7, 4},
		{8, 6},
		{9, 6},
		{100000, 6},
	}

	for _, arg := range args {
		if got := lastRemaining(arg.n); got != arg.want {
			t.Fatalf("n: %d, want: %d, got: %d\n", arg.n, arg.want, got)
		}
	}
}
