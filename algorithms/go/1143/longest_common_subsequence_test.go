package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		text1  string
		text2  string
		expect int
	}{
		{"abc", "def", 0},
		{"adb", "ab", 2},
		{"", "ab", 0},
		{"", "", 0},
		{"abalpxe", "ape", 3},
	}

	for _, tt := range tests {
		if ans := longestCommonSubsequence3(tt.text1, tt.text2); ans != tt.expect {
			t.Fatalf("text1: %s text2 %s expect %d, got %d\n", tt.text1, tt.text2, tt.expect, ans)
		}
	}
}
