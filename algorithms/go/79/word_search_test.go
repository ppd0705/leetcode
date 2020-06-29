package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		board   [][]byte
		words   []string
		expects []bool
	}{
		{[][]byte{
			{'a','a','a'},
			{'a','b','b'},
			{'a','b','b'},
			{'b','b','b'},
			{'b','b','b'},
			{'a','a','a'},
			{'b','b','b'},
			{'a','b','b'},
			{'a','a','b'},
			{'a','b','a'},
		}, []string{"aabaaaabbb"}, []bool{false}},
		{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		},
		[]string{"AAA", "ASS", "ABCCED", "SEE", "ASFS"},
		[]bool{false, false, true, true, false}},
	}

	for _, tt := range tests {
		for i := 0; i < len(tt.words); i++ {
			if ans := exist(tt.board, tt.words[i]); ans != tt.expects[i] {
				t.Fatalf("word %s search failed, got %v\n", tt.words[i], ans)
			}
		}

	}
}
