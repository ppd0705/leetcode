package main

func maxProfit(prices []int) int {
	amount := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			amount = amount + prices[i] - prices[i-1]
		}
	}
	return amount
}
