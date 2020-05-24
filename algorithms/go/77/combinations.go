package main

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	option := make([]int, k)
	var dfs func(start int, idx int)
	dfs = func(start int, idx int) {
		for i := start; i <= n-(k-1-idx); i++ {
			option[idx] = i
			if idx == k-1 {
				ans = append(ans, append([]int{}, option...))
			} else {
				dfs(i+1, idx+1)
			}
		}
	}
	dfs(1, 0)
	return ans
}

func combine0(n int, k int) [][]int {
	ans := make([][]int, 0)
	var dfs func(start int, option []int)
	dfs = func(start int, option []int) {
		for i := start; i <= n+len(option)+1-k; i++ {
			option := append(option, i)
			if len(option) == k {
				ans = append(ans, append([]int{}, option...))
			} else {
				dfs(i+1, option)
			}
			option = option[:len(option)-1]
		}
	}
	option := make([]int, 0)
	dfs(1, option)
	return ans
}
