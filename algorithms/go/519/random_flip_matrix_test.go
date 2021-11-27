package main

import (
	"log"
	"reflect"
	"sort"
	"testing"
)

func TestSolution(t *testing.T) {
	s := Constructor(3, 1)
	s.Reset()
	var points [][]int
	points = append(points, s.Flip())
	points = append(points, s.Flip())
	points = append(points, s.Flip())
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] < points[j][0] {
			return true
		}
		return false
	})
	want := [][]int{
		[]int{0, 0},
		[]int{1, 0},
		[]int{2, 0},
	}
	if !reflect.DeepEqual(points, want) {
		log.Fatalf("got %v, want: %v\n", points, want)
	}
}
