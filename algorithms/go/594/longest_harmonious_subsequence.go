package main

func findLHS(nums []int) int {
	counter := make(map[int]int)
	res := 0
	for _, num := range nums {
		counter[num]++
	}
	for num, count := range counter {
		if nextCount, ok := counter[num+1]; ok {
			res = max(res, count+nextCount)
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
