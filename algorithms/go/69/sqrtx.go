package main

func mySqrt(x int) int {
	i, j := 0, x
	for i < j {
		mid := (i + j) / 2
		square := mid * mid
		if square == x {
			return mid
		} else if square > x {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	if i*i > x {
		i--
	}
	return i
}
