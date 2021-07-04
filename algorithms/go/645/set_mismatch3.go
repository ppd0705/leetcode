package main

func findErrorNums3(nums []int) []int {
	res := []int{0, 0}
	n := len(nums)
	counter := make(map[int]int, n+1)

	for _, num := range nums {
		counter[num] += 1
	}
	for i := 1; i <= n; i++ {
		if counter[i] == 0 {
			res[1] = i
		} else if counter[i] == 2 {
			res[0] = i
		}
	}
	return res
}
