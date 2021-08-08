package main

func circularArrayLoop(nums []int) bool {
	n := len(nums)
	next := func(cur int) int {
		return (((cur + nums[cur]) % n) + n) % n
	}

	for i, num := range nums {
		if num == 0 {
			continue
		}
		slow := i
		fast := next(i)

		for nums[slow] * nums[fast] > 0 && nums[slow] * nums[next(fast)] > 0 {
			if slow == fast {
				// circle len is 1
				if slow  == next(slow) {
					break
				}
				return true
			}
			slow = next(slow)
			fast = next(next(fast))
		}

		// mark 0 as visited
		for nums[i] * nums[next(i)] > 0 {
			nums[i] = 0
			i = next(i)
		}

	}

	return false
}