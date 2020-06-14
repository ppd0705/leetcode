package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	ans := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > minPrice {
			ans = max(ans, prices[i]-minPrice)
		} else {
			minPrice = prices[i]
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
