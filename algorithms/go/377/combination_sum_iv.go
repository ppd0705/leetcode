package main

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i:= 1; i <= target; i++ {
		for _, n := range nums {
			if n <= i {
				dp[i] += dp[i-n]
			}
		}
	}
	return  dp[target]
}