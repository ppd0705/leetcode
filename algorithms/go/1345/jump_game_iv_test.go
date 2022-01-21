package main

import "testing"

func TestJumpGame(t *testing.T) {
	args := []struct {
		arr  []int
		want int
	}{
		{[]int{1}, 0},
		{[]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}, 3},
		{[]int{7, 6, 9, 6, 9, 6, 9, 7}, 1},
	}
	for _, arg := range args {
		if got := minJumps(arg.arr); got != arg.want {
			t.Fatalf("arr: %v, got: %d, want: %d\n", arg.arr, got, arg.want)
		}
	}
}
