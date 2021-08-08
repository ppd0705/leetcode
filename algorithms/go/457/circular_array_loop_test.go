package main

import "testing"

func TestCircularArray(t *testing.T) {
	samples := []struct {
		arr    []int
		target bool
	}{
		{[]int{-2,-3, -9}, false},
		{[]int{1,-1}, false},
		{[]int{1}, false},
		{[]int{1,1}, true},
		{[]int{1,2}, false},
		{[]int{1,2, 1, 2}, true},
	}

	for _, s := range samples {
		nums := append([]int{}, s.arr...)
		if res := circularArrayLoop(nums); res != s.target {
			t.Fatalf("arr: %v test failed!\n", s.arr)
		}
	}
}
