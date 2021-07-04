package main

func findErrorNums(nums []int) []int {
	res := []int{0, 0}
	n := len(nums)
	counter := make([]int, n+1)

	total := (n+1) * n / 2

	sum := 0
	set  := 0
	for _, num := range nums {
		sum += num
		if counter[num] == 0 {
			set += num
		}
		counter[num] += 1
	}
	res[0] = sum - set
	res[1] = total - set
	return res
}
