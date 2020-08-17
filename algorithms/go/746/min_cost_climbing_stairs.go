package main

func minCostClimbingStairs(cost []int) int {
	arr := make([]int, len(cost))
	arr[0], arr[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		arr[i] = min(arr[i-2], arr[i-1]) + cost[i]
	}
	return min(arr[len(cost)-1], arr[len(cost)-2])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
