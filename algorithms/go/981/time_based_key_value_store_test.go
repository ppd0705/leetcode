package main

import "testing"

func TestTimeValue(t *testing.T) {
	tv := timeValueList{
		{1, "a"},
		{3, "b"},
		{5, "c"},
		{200, "d"},
	}

	samples := []struct {
		tm     int
		target string
	}{
		{-100, ""},
		{0, ""},
		{1, "a"},
		{2, "a"},
		{3, "b"},
		{5, "c"},
		{100, "c"},
		{200, "d"},
		{300, "d"},
	}

	for _, s := range samples {
		if ans := tv.search(s.tm); ans != s.target {
			t.Fatalf("tm: %d, target: %s, got: %s\n", s.tm, s.target, ans)
		}
	}

}

func TestTimeMap(t *testing.T) {
	tm := Constructor()
	if ans:= tm.Get("a", 1); ans != "" {
		t.Fatalf("empty map error no-nil value: %s\n", ans)
	}

	tm.Set("a", "1", 1)
	if ans := tm.Get("a", 1); ans != "1" {
		t.Fatalf("expect 1, got %s\n", ans)
	}
	if ans := tm.Get("a", 3); ans != "1" {
		t.Fatalf("expect 1, got %s\n", ans)
	}
	tm.Set("a", "3", 3)
	if ans := tm.Get("a", 3); ans != "3" {
		t.Fatalf("expect 3, got %s\n", ans)
	}

	tm.Set("b", "3b", 3)
	if ans := tm.Get("a", 3); ans != "3" {
		t.Fatalf("expect 3, got %s\n", ans)
	}
	if ans := tm.Get("b", 3); ans != "3b" {
		t.Fatalf("expect 3b, got %s\n", ans)
	}
}
