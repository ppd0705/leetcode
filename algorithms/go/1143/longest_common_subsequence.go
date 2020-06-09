package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func longestCommonSubsequence(text1 string, text2 string) int {
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}

		if text1[i] == text2[j] {
			return dp(i-1, j-1) + 1
		}
		return max(dp(i-1, j), dp(i, j-1))
	}
	return dp(len(text1)-1, len(text2)-1)
}

func longestCommonSubsequence2(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for  i := 0; i <= m;i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func longestCommonSubsequence3(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([]int, n+1)
	// dp[i][j] related with 3 elements: d[i-1][j], d[i][j-1], d[i-1][j-1], first two are d[j-1], d[j] in 1D array,
	// so we stored d[i-1][j-1] as a tmp variable
	for i := 1; i <= m; i++ {
		pre := dp[0]
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if text1[i-1] == text2[j-1] {
				dp[j] = pre + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			pre = tmp
		}
	}
	return dp[n]
}

