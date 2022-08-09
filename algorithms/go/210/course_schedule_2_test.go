package main

import (
	"log"
	"reflect"
	"testing"
)

func TestFinOrder(t *testing.T) {
	args := []struct {
		num    int
		pres   [][]int
		target []int
	}{
		{1, [][]int{}, []int{0}},
		{2, [][]int{}, []int{0, 1}},
		{2, [][]int{{0, 1}}, []int{1, 0}},
		{2, [][]int{{0, 1}, {1, 0}}, []int{}},
		{3, [][]int{{0, 1}, {1, 0}}, []int{}},
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, []int{0, 1, 2, 3}},
	}

	for _, arg := range args {
		got := findOrder(arg.num, arg.pres)
		if !reflect.DeepEqual(got, arg.target) {
			log.Fatalf("target %v,  got %v\n", arg.target, got)
		}
	}

}
