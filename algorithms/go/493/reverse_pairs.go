package main

func bisectLeft(arr []int, n int) int {
	l, r := 0, len(arr)
	for l < r {
		m := (l + r) / 2
		if arr[m] < n {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func bisectInsert(arr []int, n int) []int {
	l, r := 0, len(arr)
	for l < r {
		m := (l + r) / 2
		if arr[m] < n {
			l = m + 1
		} else {
			r = m
		}
	}
	arr = append(arr, 0)
	copy(arr[l+1:len(arr)], arr[l:len(arr)-1])
	arr[l] = n
	return arr
}

func reversePairs(nums []int) int {
	sorted := make([]int, 0)
	ans := 0

	for i := len(nums) - 1; i >= 0; i-- {
		ans += bisectLeft(sorted, nums[i])
		sorted = bisectInsert(sorted, nums[i]*2)
	}
	return ans
}
