package main

func exist(board [][]byte, word string) bool {
	row := len(board)
	column := len(board[0])
	if len(word) > row*column {
		return false
	}
	visted := make([][]bool, row)
	for i := 0; i < row; i++ {
		visted[i] = make([]bool, column)
	}

	offsets := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	var dfs func(i, j, level int) bool
	dfs = func(i, j, level int) bool {
		if visted[i][j] || board[i][j] != word[level]{
			return false
		}
		if level == len(word)-1 {
			return true
		}
		visted[i][j] = true
		for k := 0; k < len(offsets); k++ {
			x, y := i+offsets[k][0], j+offsets[k][1]
			if x < 0 || x >= row || y < 0 || y >= column {
				continue
			}
			if dfs(x, y, level+1) {
				return true
			}
		}
		visted[i][j] = false
		return false
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
