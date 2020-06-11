package main

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := nums[0]
	curr := nums[0]
	for i := 1; i < len(nums); i++ {
		if curr > 0 {
			curr += nums[i]
		} else {
			curr = nums[i]
		}
		ans = max(ans, curr)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray2(nums []int) int {
	return maxSubArray2Helper(nums, 0, len(nums)-1)
}

func maxSubArray2Helper(nums []int, l, r int) int {
	if l == r {
		return nums[l]
	}
	m := (l + r) / 2
	lSum := maxSubArray2Helper(nums, l, m)
	rSum := maxSubArray2Helper(nums, m+1, r)
	mSUm := findMaxCrossingArray(nums, l, m, r)
	return max(max(lSum, rSum), mSUm)
}

func findMaxCrossingArray(nums []int, l, m, r int) int {
	lMaxSum := math.MinInt64
	rMaxSum := math.MinInt64
	sum := 0
	for i := m; i >= l; i-- {
		sum += nums[i]
		lMaxSum = max(lMaxSum, sum)
	}
	sum = 0
	for j := m + 1; j <= r; j++ {
		sum += nums[j]
		rMaxSum = max(rMaxSum, sum)
	}
	return lMaxSum + rMaxSum
}
