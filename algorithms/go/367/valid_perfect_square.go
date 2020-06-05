package main

func isPerfectSquare(num int) bool {
	i, j := 0, num
	for i <= j {
		mid := (i + j) / 2
		square := mid * mid
		if square == num {
			return true
		} else if square > num {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return false
}
