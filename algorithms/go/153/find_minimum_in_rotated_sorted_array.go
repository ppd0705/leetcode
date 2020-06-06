package main

func findMin(nums []int) int {
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}
	l, r := 0, len(nums)-1

	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		if nums[mid] > nums[0] {
			l = mid + 1
		} else {
			r = mid -1
		}
	}
	return -1
}

func findMin2(nums []int) int {
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}
	l, r := 0, len(nums)-1

	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		if nums[mid] > nums[0] {
			l = mid + 1
		} else {
			r = mid -1
		}
	}
	return -1
}