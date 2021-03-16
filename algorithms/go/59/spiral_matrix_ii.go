package main

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i:= 0; i< n; i++ {
		matrix[i] = make([]int, n)
	}
	matrix[0][0] = 1
	dp := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var curX, curY, nextX, nextY, dpIdx int
	for i := 2; i <= n*n; i++ {
		nextX, nextY = curX + dp[dpIdx][0], curY + dp[dpIdx][1]
		if nextX == -1 || nextY == -1 ||nextX == n || nextY == n || matrix[nextX][nextY] != 0 {
			dpIdx = (dpIdx+1) % 4
		}
		curX, curY = curX + dp[dpIdx][0], curY + dp[dpIdx][1]
		matrix[curX][curY] = i
	}
	return matrix
}