package main

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}
		if nums[l] == nums[r] {
			l++
			continue
		}
		if nums[mid] >= nums[l] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[len(nums)-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}
