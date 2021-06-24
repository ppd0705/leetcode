package main

import "testing"

func Test(t *testing.T) {
	samples := []struct {
		points [][]int
		target int
	}{
		{[][]int{
			{1, 1},
			{2, 2},
			{4, 2},
		},
			2,
		},
		{[][]int{
			{3, 2},
			{2, 2},
			{4, 2},
		},
			3,
		},
		{[][]int{
			{0, 0},
			{2, 2},
			{4, 4},
		},
			3,
		},
	}

	for _, s := range samples {
		if ans := maxPoints(s.points); ans != s.target {
			t.Fatalf("target: %d, got %d\n", s.target, ans)
		}
	}
}
