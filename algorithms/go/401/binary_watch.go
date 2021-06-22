package main

import "fmt"

func readBinaryWatch(turnedOn int) []string {
	hours := map[int][]int{
		0: {0},
	}
	minutes := map[int][]int{
		0: {0},
	}

	nums := []int{1, 2, 4, 8, 16, 32}

	queue := make([]int, 0)
	var calculate func(idx, count, target, sum int)
	calculate = func(idx, count, target, sum int) {
		sum += nums[idx]
		count += 1
		if count == target {
			queue = append(queue, sum)
			return
		}
		if idx == len(nums) {
			return
		}
		for i := idx + 1; i < len(nums); i++ {
			calculate(i, count, target, sum)
		}
	}

	for i := 0; i < 6; i++ {
		for j := 1; j <= turnedOn; j++ {
			calculate(i, 0, j, 0)
			for k := 0; k < len(queue); k++ {
				res := queue[k]
				if res > 0 {
					if res <= 11 {
						if _, ok := hours[j]; !ok {
							hours[j] = []int{}
						}
						hours[j] = append(hours[j], res)
					}
					if res <= 59 {
						if _, ok := minutes[j]; !ok {
							minutes[j] = []int{}
						}
						minutes[j] = append(minutes[j], res)
					}
				}
			}
			queue = queue[:0]
		}
	}

	ans := make([]string, 0)

	for i := 0; i <= turnedOn; i++ {
		j := turnedOn - i
		if hours[i] != nil && minutes[j] != nil {
			for _, hour := range hours[i] {
				for _, minute := range minutes[j] {
					ans = append(ans, fmt.Sprintf("%d:%02d", hour, minute))
				}
			}
		}
	}
	return ans
}
