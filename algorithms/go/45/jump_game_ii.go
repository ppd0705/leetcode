package main

func jump(nums []int) int {
	end, maxPos, step := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(maxPos, i+nums[i])
		if i == end {
			step++
			end = maxPos
		}
	}
	return step
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
