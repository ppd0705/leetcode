package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
	}{
		{[]int{-2,-2,1,1,4,1,4,4,-4,-2}, -4},
		{[]int{2, 2, 3, 2}, 3},
		{[]int{9}, 9},
		{[]int{9, 1, 4, 1, 4, 1, 4}, 9},
	}

	for _, tt := range tests {
		if ans := singleNumber2(tt.nums); ans != tt.target {
			t.Fatalf("expect %d, got %d\n", tt.target, ans)
		}
	}
}
