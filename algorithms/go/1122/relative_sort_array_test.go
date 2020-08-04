package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		arr1   []int
		arr2   []int
		target []int
	}{
		{
			[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			[]int{2, 1, 4, 3, 9, 6},
			[]int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			[]int{2, 3, 1, 3, 2},
			[]int{},
			[]int{1, 2, 2, 3, 3},
		},
	}

	for _, tt := range tests {
		if ans := relativeSortArray2(tt.arr1, tt.arr2); !reflect.DeepEqual(ans, tt.target) {
			t.Fatalf("expcet %v, got %v\n", tt.target, ans)
		}
	}
}
