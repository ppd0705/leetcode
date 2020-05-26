package main

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		input  string
		expect []string
	}{
		{"2", []string{"a", "b", "c"}},
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}

	for _, tt := range tests {
		if ans := letterCombinations(tt.input); !reflect.DeepEqual(tt.expect, ans) {
			t.Fatalf("expect: %v, got %v\n", tt.expect, ans)
		}
	}
}
