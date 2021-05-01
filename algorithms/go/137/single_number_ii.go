package main

func singleNumber(nums []int) int {
	bits := make([]int, 32)
	for _, num := range nums {
		for i:= 0; i< 32;i++ {
			bits[i] += 1 & num
			num >>= 1
		}
	}
	res := 0
	for i := 0; i < 31; i++ {
		if bits[i] % 3 != 0 {
			res +=  1 << i
		}
	}
	if bits[31] % 3 == 1 {
		res -=  1 << 31
	}
	return res
}
