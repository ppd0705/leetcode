package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		strs   []string
		target string
	}{
		{[]string{"abc", "a"}, "a"},
		{[]string{}, ""},
		{[]string{"ac", "cc"}, ""},
		{[]string{"ac", "ace"}, "ac"},
		{[]string{"aaa", ""}, ""},
		{[]string{"aer", "a", "aerrrrrr"}, "a"},
	}

	for _, tt := range tests {
		if ans := longestCommonPrefix(tt.strs); ans != tt.target {
			t.Fatalf("target: %s, got %s\n", tt.target, ans)
		}
	}
}
