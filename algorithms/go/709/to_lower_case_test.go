package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		src string
		dst string
	}{
		{"q123", "q123"},
		{"", ""},
		{"12", "12"},
		{"hOlll", "holll"},
	}

	for _, tt := range tests {
		if ans := toLowerCase(tt.src); ans != tt.dst {
			t.Fatalf("expect %s, got %s\n", tt.dst, ans)
		}
	}
}
