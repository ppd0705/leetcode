package main

func PredictTheWinner2(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	for i := n - 2; i >= 0; i-- {
		for j := i+1; j < n; j++ {
			dp[i][j] = max2(nums[i] - dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] >= 0
}

func max2(x, y int) int {
	if x > y {
		return x
	}
	return y
}
