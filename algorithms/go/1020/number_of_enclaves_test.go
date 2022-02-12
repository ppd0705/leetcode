package main

import "testing"

func Test(t *testing.T) {
	args := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{1, 1, 1}}, 0},
		{[][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}, 3},
		{[][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}, 0},
	}

	for i, arg := range args {
		if got := numEnclaves(arg.grid); got != arg.want {
			t.Fatalf("%d got %d, want %d\n", i, got, arg.want)
		}
	}
}
