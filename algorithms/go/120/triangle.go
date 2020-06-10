package main

func minimumTotal(triangle [][]int) int {
	row := len(triangle)
	if row == 0 {
		return -1
	}
	column := len(triangle[row-1])
	if column == 0 {
		return -1
	}
	dp := make([]int, column)
	copy(dp, triangle[row-1])
	for i := row - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
		}
	}
	return dp[0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
