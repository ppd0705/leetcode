package main

func majorityElement(nums []int) int {
	counter := make(map[int]int, 0)
	for _, n := range nums {
		counter[n]++
		if counter[n]*2 > len(nums) {
			return n
		}
	}
	return -1
}

func majorityElement2(nums []int) int {
	candidate := nums[0]
	count := 0

	for _, n := range nums {
		if count == 0 {
			candidate = n
		}

		if candidate == n {
			count++
		} else {
			count--
		}
	}
	return candidate
}
