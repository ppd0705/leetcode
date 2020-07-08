package main

func boxIdx(i, j int) int {
	return i/3*3 + j/3
}

func solveSudoku(board [][]byte) {
	row := [9][9]bool{}
	column := [9][9]bool{}
	box := [9][9]bool{}
	toFill := [9][9]bool{}

	missed := 0

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				missed++
				toFill[i][j] = true
			} else {
				v := board[i][j] - 49
				row[i][v] = true
				column[j][v] = true
				box[boxIdx(i, j)][v] = true
			}
		}
	}

	canFill := func(i, j, v int) bool {
		return !row[i][v] && !column[j][v] && !box[boxIdx(i, j)][v]
	}

	getNext := func(i, j int) (int, int) {
		if j == 8 {
			return i + 1, 0
		}
		return i, j + 1
	}
	var dfs func(i, j, cnt int) bool
	dfs = func(i, j, cnt int) bool {
		if cnt == 0 {
			return true
		}
		if toFill[i][j] {
			for k := 0; k < 9; k++ {
				if canFill(i, j, k) {
					row[i][k] = true
					column[j][k] = true
					box[boxIdx(i, j)][k] = true
					board[i][j] = byte('1' + k)

					nxtI, nxtJ := getNext(i, j)

					if dfs(nxtI, nxtJ, cnt-1) {
						return true
					}
					row[i][k] = false
					column[j][k] = false
					box[boxIdx(i, j)][k] = false
					board[i][j] = '.'
				}
			}
		} else {
			nxtI, nxtJ := getNext(i, j)
			return dfs(nxtI, nxtJ, cnt)
		}
		return false
	}
	dfs(0, 0, missed)
}
