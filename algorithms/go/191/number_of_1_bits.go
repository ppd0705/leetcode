package main

func hammingWeight(num uint32) int {
	ans := 0
	for num > 0 {
		if num&1 > 0 {
			ans++
		}
		num >>= 1
	}
	return ans
}

func hammingWeight2(num uint32) int {
	ans := 0
	for num > 0 {
		ans++
		num &= num - 1
	}
	return ans
}
