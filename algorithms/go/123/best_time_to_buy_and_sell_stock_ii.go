package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp10 := 0
	dp20 := 0
	dp11 := -prices[0]
	dp21 := -prices[0]

	for _, p := range prices[1:] {
		dp20 = max(dp20, dp21+p)
		dp21 = max(dp21, dp10-p)
		dp10 = max(dp10, dp11+p)
		dp11 = max(dp11, dp10-p)
	}
	return dp20
}

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	n := 2
	dp := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = []int{0, 0}
		}
	}
	for i := 0; i < len(prices); i++ {
		for j := 1; j <= n; j++ {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
			} else {
				dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
				dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
			}
		}
	}
	return dp[len(prices)-1][n][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
