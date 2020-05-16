package main

var cache = map[int]int{}

func climbStairs2(n int) int {
	if res, ok := cache[n]; ok {
		return res
	}

	var res int
	if n < 3 {
		res =  n
	} else {
		res = climbStairs2(n-1) + climbStairs2(n-2)
	}
	cache[n] = res
	return res
}
