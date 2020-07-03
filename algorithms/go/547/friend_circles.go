package main

func findCircleNum(M [][]int) int {
	n := len(M)
	ans := 0
	visted := make([]bool, n)

	var dfs func(int)
	dfs = func(i int) {
		for j := 0; j < n; j++ {
			if !visted[j] && M[i][j] == 1 {
				visted[j] = true
				dfs(j)
			}
		}
	}
	for i := 0; i < n; i++ {
		if !visted[i] {
			ans++
			visted[i] = true
			dfs(i)
		}
	}

	return ans
}
