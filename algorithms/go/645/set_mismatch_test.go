package main

import (
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	samples := []struct {
		nums   []int
		target []int
	}{
		{[]int{3,2,3,4,6,5}, []int{3, 1}},
		{[]int{1, 1}, []int{1, 2}},
		{[]int{1, 2, 3, 2}, []int{2, 4}},
	}

	for _, s := range samples {
		ans := findErrorNums(s.nums)
		if !reflect.DeepEqual(ans, s.target) {
			t.Fatalf("ans: %v, target: %v\n", ans, s.target)
		}
	}
}

