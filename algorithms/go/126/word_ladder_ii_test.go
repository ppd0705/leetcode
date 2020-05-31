package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		beginWord string
		endWord   string
		wordList  []string
		expect    [][]string
	}{
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"},
			[][]string{
				{"hit", "hot", "dot", "dog", "cog"}, {"hit", "hot", "lot", "log", "cog"},
			}},
	}

	for _, tt := range tests {
		if ans := findLadders4(tt.beginWord, tt.endWord, tt.wordList); !reflect.DeepEqual(ans, tt.expect) {
			t.Fatalf("failed. expect %v, got %v\n", tt.expect, ans)
		}
	}

}
