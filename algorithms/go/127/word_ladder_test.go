package main

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		beginWord string
		endWord   string
		wordList  []string
		expect    int
	}{
		{"ymain", "oecij", []string{"ymann","yycrj","oecij","ymcnj","yzcrj","yycij","xecij","yecij","ymanj","yzcnj","ymain"}, 10},
		{"a", "c", []string{"a", "b", "c"}, 2},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, 0},
		{"hit", "cog", []string{"hot", "cog"}, 0},
	}

	for _, tt := range tests {
		if ans := ladderLength5(tt.beginWord, tt.endWord, tt.wordList); ans != tt.expect {
			t.Fatalf("failed. expect %d, got %d\n", tt.expect, ans)
		}
	}
}