package main

import "testing"

func Test(t *testing.T) {
	samples := []struct {
		preferences [][]int
		pairs       [][]int
		target      int
	}{
		{
			[][]int{{1, 2, 3}, {3, 2, 0}, {3, 1, 0}, {1, 2, 0}},
			[][]int{{0, 1}, {2, 3}},
			2,
		},
		{
			[][]int{{1}, {0}},
			[][]int{{0, 1}},
			0,
		},
		{
			[][]int{{1, 3, 2}, {2, 3, 0}, {1, 3, 0}, {0, 2, 1}},
			[][]int{{3, 1}, {0, 2}},
			4,
		},
	}

	for _, s := range samples {
		n := len(s.preferences)
		res := unhappyFriends2(n, s.preferences, s.pairs)
		if res != s.target {
			t.Fatalf("preferences: %v, paires: %d, want: %d, got %d",
				s.preferences, s.pairs, s.target, res,
			)
		}
	}
}
