package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit(prices []int, fee int) int {
	hold, notHold := -prices[0]-fee, 0
	for _, p := range prices[1:] {
		hold, notHold = max(hold, notHold-p-fee), max(notHold, hold+p)
	}
	return notHold
}
