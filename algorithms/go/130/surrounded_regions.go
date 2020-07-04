package main

func solve(board [][]byte) {
	m := len(board)
	if m < 3 {
		return
	}
	n := len(board[0])
	if n < 3 {
		return
	}
	o := byte('o')
	x := byte('x')

	notSurrounded := make([][]bool, len(board))

	for i := 0; i < m; i++ {
		notSurrounded[i] = make([]bool, n)
	}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x > 0 && board[x-1][y] == o && !notSurrounded[x-1][y] {
			notSurrounded[x-1][y] = true
			dfs(x-1, y)
		}
		if x < m-1 && board[x+1][y] == o && !notSurrounded[x+1][y] {
			notSurrounded[x+1][y] = true
			dfs(x+1, y)
		}

		if y > 0 && board[x][y-1] == o && !notSurrounded[x][y-1] {
			notSurrounded[x][y-1] = true
			dfs(x, y-1)
		}

		if y < n-1 && board[x][y+1] == o && !notSurrounded[x][y+1] {
			notSurrounded[x][y+1] = true
			dfs(x, y+1)
		}
	}

	for i := 0; i < m; i++ {
		if board[i][0] == o && !notSurrounded[i][0] {
			notSurrounded[i][0] = true
			dfs(i, 0)
		}
		if board[i][n-1] == o && !notSurrounded[i][n-1] {
			notSurrounded[i][n-1] = true
			dfs(i, n-1)
		}
	}

	for j := 1; j < n-1; j++ {
		if board[0][j] == o && !notSurrounded[0][j] {
			notSurrounded[0][j] = true
			dfs(0, j)
		}
		if board[m-1][j] == o && !notSurrounded[m-1][j] {
			notSurrounded[m-1][j] = true
			dfs(m-1, j)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == o && !notSurrounded[i][j] {
				board[i][j] = x
			}
		}
	}
}
