package main

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	bucketMap := make(map[int]int)
	for i, n := range nums {
		idx := getBucket(n, t+1)
		if _, ok := bucketMap[idx]; ok {
			return true
		} else {
			if m, ok := bucketMap[idx+1]; ok && abs(n-m) <= t {
				return  true
			}
			if m, ok := bucketMap[idx-1]; ok && abs(n-m) <= t {
				return  true
			}
			bucketMap[idx] = n
		}
		if i >= k {
			delete(bucketMap, getBucket(nums[i-k], t+1))
		}
	}
	return false
}

func getBucket(num, size int) int {
	if num >= 0 {
		return num / size
	}
	return (num+1)/size-1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

