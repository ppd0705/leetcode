package three_sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var ret [][]int
	numsLen := len(nums)
	sort.Ints(nums)
	for i := 0; i < numsLen-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, numsLen-1
		for j < k {
			s := nums[i] + nums[j] + nums[k]
			if s > 0 {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if s < 0 {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}

				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}

		}

	}

	return ret
}
