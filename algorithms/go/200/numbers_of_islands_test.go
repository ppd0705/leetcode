package main

import (
	"testing"
)

func Test(t *testing.T) {
	grid := [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'},
	}
	ans := numIslands(grid)
	if ans != 3 {
		t.Fatalf("failed. got : %d", ans)
	}
}
