package main

func lastRemaining(n int) int {
	a1 := 1
	step := 1
	remain := n
	isLeft := true
	for remain > 1 {
		if isLeft || remain%2 == 1 {
			a1 += step
		}
		remain >>= 1
		step <<= 1
		isLeft = !isLeft
	}
	return a1
}
