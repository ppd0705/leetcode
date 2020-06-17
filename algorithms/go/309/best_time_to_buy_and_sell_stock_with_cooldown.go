package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = [][]int{{0, 0}, {0, 0}} // 1th: if today sel or not; 2th: if hold stock
	}

	for i, p := range prices {
		if i == 0 {
			dp[i][0][1] = -p
		} else {
			dp[i][0][1] = max(dp[i-1][0][1], dp[i-1][0][0]-p)
			dp[i][0][0] = max(dp[i-1][0][0], dp[i-1][1][0])
			dp[i][1][0] = dp[i-1][0][1] + p
		}
	}
	return max(dp[len(prices)-1][1][0], dp[len(prices)-1][0][0])
}

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	if len(prices) == 2 {
		return max(0, prices[1]-prices[0])
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = []int{0, 0} // if not hold stock or hold
	}

	dp[0][1] = -prices[0]
	dp[1][0] = max(dp[0][0], dp[0][1]+prices[1])
	dp[1][1] = max(-prices[1], dp[0][1])

	for i := 2; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
		fmt.Printf("%d %d %d %d\n", i, prices[i], dp[i][0], dp[i][1])
	}
	return dp[len(prices)-1][0]
}

func maxProfit3(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp10 := 0
	dp11 := -prices[0]
	dp20 := max(0, dp11+prices[1])
	dp21 := max(dp11, dp10-prices[1])
	for i := 2; i < len(prices); i++ {
		dp11, dp21 = dp21, max(dp21, dp10-prices[i])
		dp10, dp20 = dp20, max(dp20, dp11+prices[i])
	}
	return dp20
}

