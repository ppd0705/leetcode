package main

import (
	"testing"
)

func TestGetCoordinate(t *testing.T) {
	samples := []struct {
		n  int
		x  int
		dx int
		dy int
	}{
		{6, 13, 3, 0},
		{5, 13, 2, 2},
	}
	for _, s := range samples {
		dx, dy := getCoordinate(s.x, s.n)
		if dx != s.dx || dy != s.dy {
			t.Fatalf("x: %d, target: [%d, %d], get: [%d, %d]\n", s.x, s.dx, s.dy, dx, dy)
		}
	}
}
