package main

import "sort"

func findErrorNums2(nums []int) []int {
	res := []int{0, 0}
	sort.Ints(nums)
	pre := 0
	for _, num := range nums {
		if num == pre {
			res[0] = num
		} else if num - pre > 1 {
			res[1] = pre + 1
		}
		pre = num
	}
	if nums[len(nums)-1] != len(nums) {
		res[1] = len(nums)
	}
	return res
}