package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) <= 1 {
		return 0
	}

	if 2*k >= len(prices) {
		return maxProfitWithoutLimit(prices)
	}
	dp := make([][][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][2]int, k+1)
	}

	for i, p := range prices {
		for j := 1; j <= k; j++ {
			if i == 0 {
				dp[i][j][1] = -p
			} else {
				dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+p)
				dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-p)
			}
		}
	}
	return dp[len(prices)-1][k][0]
}

func maxProfitWithoutLimit(prices []int) int {
	hold, notHold := -prices[0], 0
	for _, p := range prices[1:] {
		hold, notHold = max(hold, notHold-p), max(notHold, hold+p)
	}
	return notHold
}
