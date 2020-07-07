package main

func isValidSudoku0(board [][]byte) bool {
	numberCount := make(map[byte]int, 9)

	// row
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				numberCount[board[i][j]]++
			}
		}
		for k, v := range numberCount {
			if v > 1 {
				return false
			}
			numberCount[k] = 0
		}
	}

	// column
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				numberCount[board[j][i]]++
			}
		}
		for k, v := range numberCount {
			if v > 1 {
				return false
			}
			numberCount[k] = 0
		}
	}

	// column
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				numberCount[board[j][i]]++
			}
		}
		for k, v := range numberCount {
			if v > 1 {
				return false
			}
			numberCount[k] = 0
		}
	}

	// box

	ranges := [4]int{0, 3, 6, 9}
	for m := 0; m < 3; m++ {
		for l := 0; l < 3; l++ {
			for i := ranges[m]; i < ranges[m+1]; i++ {
				for j := ranges[l]; j < ranges[l+1]; j++ {
					if board[i][j] != '.' {
						numberCount[board[i][j]]++
					}
				}
			}
			for k, v := range numberCount {
				if v > 1 {
					return false
				}
				numberCount[k] = 0
			}
		}
	}
	return true
}

func isValidSudoku1(board [][]byte) bool {
	rowCount := [9][9]bool{}
	columnCount := [9][9]bool{}
	boxCount := [9][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b := board[i][j]
			if b != '.' {
				n := b - 49
				if rowCount[i][n] {
					return false
				} else {
					rowCount[i][n] = true
				}
				if columnCount[j][n] {
					return false
				} else {
					columnCount[j][n] = true
				}

				k := (i/3)*3 + j/3
				if boxCount[k][n] {
					return false
				} else {
					boxCount[k][n] = true
				}

			}
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	var row, column, box [9]uint16
	var cur uint16
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b := board[i][j]
			if b == '.' {
				continue
			}
			cur = 1 << (b - '1')

			k := i/3*3 + j/3
			if (cur&row[i])|(cur&column[j])|(box[k]&cur) != 0 {
				return false
			}

			row[i] |= cur
			column[j] |= cur
			box[k] |= cur
		}
	}
	return true
}
