package main

func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	left, right := 1, n
	for left <= right {
		mid := (left  + right) / 2
		if isBadVersion(mid) {
			if mid == 1 || !isBadVersion(mid-1) {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}