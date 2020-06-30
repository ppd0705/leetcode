package main

type Trie struct {
	children map[byte]*Trie
	isEnd    bool
}

func findWords(board [][]byte, words []string) []string {
	var res []string
	mark := byte('#')
	trie := &Trie{children: map[byte]*Trie{}}

	for _, w := range words {
		node := trie
		for _, c := range []byte(w) {
			if _, ok := node.children[c]; !ok {
				node.children[c] = &Trie{children: map[byte]*Trie{}}
			}
			node = node.children[c]
		}
		node.isEnd = true
	}
	offsets := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	row := len(board)
	column := len(board[0])
	var dfs func(i, j int, s string, t *Trie)
	dfs = func(i, j int, s string, t *Trie) {
		v := board[i][j]
		if t, ok := t.children[v]; !ok {
			return
		} else {
			s += string(v)
			if t.isEnd {
				res = append(res, s)
				t.isEnd = false
			}

			board[i][j] = mark
			for k := 0; k < len(offsets); k++ {
				x, y := i+offsets[k][0], j+offsets[k][1]
				if x >= 0 && x < row && y >= 0 && y < column && board[x][y] != mark {
					dfs(x, y, s, t)
				}
			}
			board[i][j] = v
		}

	}

	for r := 0; r < row; r++ {
		for c := 0; c < column; c++ {
			dfs(r, c, "", trie)
		}
	}
	return res
}
