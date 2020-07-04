package main

var dummy int

func solve2(board [][]byte) {
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

	p := make([]int, m*n+1)
	dummy = len(p) - 1
	for i := 0; i < len(p); i++ {
		p[i] = i
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || i == m-1 || j == 0 || j == n-1) && board[i][j] == o {
				p[i*n+j] = dummy
			}
		}
	}
	dp := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == o {
				if i == 0 || i == m-1 || j == 0 || j == n-1 {
					continue
				} else {
					for k := range dp {
						if board[i+dp[k][0]][j+dp[k][1]] == o {
							union(p, i*n+j, (i+dp[k][0])*n+j+dp[k][1])
						}
					}
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == o && find(p, p[i*n+j]) != dummy {
				board[i][j] = x
			}
		}
	}
}

func find(p []int, x int) int {
	for p[x] != x && x != dummy {
		p[x] = p[p[x]]
		x = p[x]
	}
	return x
}

func union(p []int, x, y int) {
	px := find(p, x)
	py := find(p, y)
	if px == dummy {
		p[py] = px
	} else {
		p[px] = py
	}
}
