package main

func canJump(nums []int) bool {
	k := 0
	for i := 0; k < len(nums)-1; i++ {
		if i > k {
			return false
		}
		if i+nums[i] > k {
			k = i + nums[i]
		}
	}
	return true
}
