package main

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int,  m+1)
	for i := 0; i<= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}

	for i := 1; i <= m; i++ {
		for j :=1; j <= n; j++ {
			if s[i-1] != t[j-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				// t[j-1]和 s[i-1]匹配 +  t[j-1] 和s[:i-1]某个位置匹配
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}