package main

var cache map[[2]int]int

func PredictTheWinner(nums []int) bool {
	cache = map[[2]int]int{}
	return play(nums, 0, len(nums)-1) >= 0
}

func play(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	if val, ok := cache[[2]int{left, right}]; ok {
		return val
	}
	chooseLeft := nums[left] - play(nums, left+1, right)
	chooseRight := nums[right] - play(nums, left, right-1)
	res := max(chooseLeft, chooseRight)
	cache[[2]int{left, right}] = res
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
