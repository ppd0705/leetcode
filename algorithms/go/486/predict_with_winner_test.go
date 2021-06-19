package main

import (
	"testing"
)

func Test(t *testing.T) {
	samples := []struct{
		nums []int
		target bool
	}{
		{[]int{1}, true},
		{[]int{1, 1}, true},
		{[]int{1, 1, 1}, true},
		{[]int{1, 3, 1}, false},
	}

	for _, sample := range samples {
		res := PredictTheWinner2(sample.nums)
		if res != sample.target {
			t.Fatalf("failed for %v\n", sample.nums)
		}
	}
}