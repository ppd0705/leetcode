package main

func numEnclaves(grid [][]int) int {
	m := len(grid)
	if m == 1 {
		return 0
	}
	n := len(grid[0])
	if n == 1 {
		return 0
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i == m || j < 0 || j == n {
			return
		}
		if grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	for y := 0; y < n; y++ {
		dfs(0, y)
		dfs(m-1, y)
	}
	for x := 1; x < m-1; x++ {
		dfs(x, 0)
		dfs(x, n-1)
	}
	ret := 0
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			ret += grid[x][y]
		}
	}
	return ret
}
