package main

func maxProduct(nums []int) int {
	maxF := nums[0]
	minF := nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		m, n := maxF, minF
		maxF = max(n*nums[i], max(nums[i], m*nums[i]))
		minF = min(n*nums[i], min(nums[i], m*nums[i]))
		ans = max(ans, maxF)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
