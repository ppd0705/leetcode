package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		t      int
		target bool
	}{
		{[]int{1,5,9,1,5,9}, 2, 3, false},
		{[]int{-2, -1, 0, 1, 2, 3, 4, 5, 6}, 2, 2, true},
	}

	for _, tt := range tests {
		if ans := containsNearbyAlmostDuplicate(tt.nums, tt.k, tt.t); ans != tt.target {
			t.Fatalf("expect %t, but got %t", tt.target, ans)
		}
	}
}
