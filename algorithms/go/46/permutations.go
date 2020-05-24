package main

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	option := make([]int, len(nums))
	var dfs func(int)
	dfs = func(idx int) {
		if idx == len(nums) {
			ans = append(ans, append([]int{}, option...))
			return
		}
		for i := idx; i < len(nums); i++ {
			option[idx] = nums[i]
			nums[i], nums[idx] = nums[idx], nums[i]
			dfs(idx + 1)
			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}
	dfs(0)
	return ans
}

func permute0(nums []int) [][]int {
	ans := make([][]int, 0)
	option := make([]int, len(nums))
	visted := make([]int, len(nums))
	var dfs func(int)
	dfs = func(idx int) {
		if idx == len(nums) {
			ans = append(ans, append([]int{}, option...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if visted[i] == 0 {
				option[idx] = nums[i]
				visted[i] = 1
				dfs(idx + 1)
				visted[i] = 0
			}
		}
	}
	dfs(0)
	return ans
}
