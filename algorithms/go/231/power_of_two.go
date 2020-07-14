package main

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func isPowerOfTwo1(n int) bool {
	return n > 0 && n&-n == n
}
