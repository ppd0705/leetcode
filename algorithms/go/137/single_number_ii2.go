package main

func singleNumber2(nums []int) int {
	res := 0
	for i  := 0; i < 32; i++ {
		total := 0
		for _, num := range nums {
			total += num >> i & 1
		}
		if total % 3 == 1 {
			if i == 31 {
				res -= 1 << i
			} else {
				res |= 1 << i
			}
		}
	}
	return res
}
