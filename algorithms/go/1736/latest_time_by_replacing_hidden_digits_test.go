package main

import "testing"

func TestLatestTime(t *testing.T) {
	samples := []struct {
		time   string
		target string
	}{
		{"??:??", "23:59"},
		{"11:22", "11:22"},
		{"1?:3?", "19:39"},
		{"?2:?1", "22:51"},
		{"?5:?5", "15:55"},
		{"0?:0?", "09:09"},
	}

	for _, s := range samples {
		if res := maximumTime(s.time); res != s.target {
			t.Fatalf("target: %s, got: %s\n", s.target, res)
		}
	}
}
