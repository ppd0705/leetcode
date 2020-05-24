package main

func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	option := make([]int, len(nums))
	elements := make(map[int]int, 0)
	for _, n := range nums {
		elements[n]++
	}

	var dfs func(int)
	dfs = func(idx int) {
		if idx == len(nums) {
			ans = append(ans, append([]int{}, option...))
			return
		}
		for k, v := range elements {
			if v > 0 {
				option[idx] = k
				elements[k]--
				dfs(idx + 1)
				elements[k]++
			}
		}
	}
	dfs(0)
	return ans
}
