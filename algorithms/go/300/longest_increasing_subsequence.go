package main

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := 1
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		m := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				m = max(m, dp[j])
			}
		}
		dp[i] = m + 1
		ans = max(ans, dp[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLIS2(nums []int) int {
	tails := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if len(tails) == 0 || nums[i] > tails[len(tails)-1] {
			tails = append(tails, nums[i])
		} else {
			l, r := 0, len(tails)

			for l < r {
				m := (l + r) / 2
				if tails[m] >= nums[i] {
					r = m
				} else {
					l = m + 1
				}
			}
			tails[l] = nums[i]
		}
	}
	return len(tails)
}
