package main

import "fmt"

func numIslands(grid [][]byte) int {
	row := len(grid)
	count := 0
	if row == 0 {
		return count
	}
	column := len(grid[0])

	var dfs func(x, y int)
	dfs = func(x, y int) {
		grid[x][y] = '0'
		if x > 0 && grid[x-1][y] == '1' {
			dfs(x-1, y)
		}
		if x+1 < row && grid[x+1][y] == '1' {
			dfs(x+1, y)
		}

		if y > 0 && grid[x][y-1] == '1' {
			dfs(x, y-1)
		}
		if y+1 < column && grid[x][y+1] == '1' {
			dfs(x, y+1)
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == '1' {
				fmt.Println(i, j)
				count++
				dfs(i, j)
			}
		}
	}
	return count
}
