package main

func solveNQueens(n int) [][]string {
	colUsed := make(map[int]bool)
	leftCrossLineUsed := make(map[int]bool)  // row + col
	rightCrossLineUsed := make(map[int]bool) // row - col
	colPos := make([]int, n)

	ans := make([][]string, 0)

	addAns := func() []string {
		queens := make([]string, n)
		template := make([]byte, n)
		for row:= 0;row < n;row++ {
			for col:= 0; col < n; col++{
				if colPos[row] == col {
					template[col] = 'Q'
				} else {
					template[col] = '.'
				}
			}
			queens[row] = string(template)
		}
		return queens
	}

	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			ans = append(ans, addAns())
			return
		}
		for col := 0; col < n; col += 1 {
			l := row + col
			r := row - col
			if colUsed[col] || leftCrossLineUsed[l] || rightCrossLineUsed[r] {
				continue
			}
			colUsed[col] = true
			leftCrossLineUsed[l] = true
			rightCrossLineUsed[r] = true
			colPos[row] = col

			// drill down
			dfs(row + 1)

			// reverse
			colUsed[col] = false
			leftCrossLineUsed[l] = false
			rightCrossLineUsed[r] = false
		}
	}
	dfs(0)
	return ans
}
