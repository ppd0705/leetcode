package main

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	var dfs func(cur []int, idx int)
	dfs = func(cur []int, idx int) {
		ans = append(ans, append([]int{}, cur...))
		for i := idx; i < len(nums); i++ {
			cur = append(cur, nums[i])
			dfs(cur, i+1)
			cur = cur[:len(cur)-1]
		}
	}
	empty := make([]int, 0)
	dfs(empty, 0)
	return ans
}
