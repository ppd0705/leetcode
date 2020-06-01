package main

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt64
		for _, c := range coins {
			if c <= i {
				if dp[i-c] < dp[i] {
					dp[i] = dp[i-c] + 1
				}
			}
		}
	}

	if dp[amount] != math.MaxInt64 {
		return dp[amount]
	} else {
		return -1
	}
}

func coinChange2(coins []int, amount int) int {
	memo := make(map[int]int, 0)
	memo[0] = 0
	var dp func(n int) int
	dp = func(n int) int {
		if n < 0 {
			return -1
		}
		if v, ok := memo[n]; ok {
			return v
		}
		res := math.MaxInt64
		for _, c := range coins {
			sub := dp(n - c)
			if sub != -1 && sub < res {
				res = sub + 1
			}
		}
		if res == math.MaxInt64 {
			memo[n] =  -1
		} else {
			memo[n] = res
		}
		return memo[n]
	}
	return dp(amount)
}
