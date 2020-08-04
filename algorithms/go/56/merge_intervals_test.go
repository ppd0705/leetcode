package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		intervals [][]int
		target    [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 3}, {1, 4}, {7, 10}, {6, 20}}, [][]int{{1, 4}, {6, 20}}},
	}

	for _, tt := range tests {
		if ans := merge(tt.intervals); !reflect.DeepEqual(ans, tt.target) {
			t.Fatalf("got %v, expect %v\n", ans, tt.target)
		}
	}
}
