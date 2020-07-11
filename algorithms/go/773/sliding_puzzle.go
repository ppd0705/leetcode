package main

func slidingPuzzle(board [][]int) int {
	visited := make(map[string]bool, 0)

	neighbors := map[int][]int{
		0: {1, 3},
		1: {0, 2, 4},
		2: {1, 5},
		3: {0, 4},
		4: {1, 3, 5},
		5: {2, 4},
	}
	m, n := len(board), len(board[0])
	startChars := make([]byte, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			startChars[i*n+j] = byte('0' + board[i][j])
		}
	}
	visited[string(startChars)] = true
	target := "123450"
	queue := []string{string(startChars)}
	ans := -1

	swapStr := func(str string, x, y int) string {
		if x > y {
			x, y = y, x
		}
		return str[:x] + string(str[y]) +  str[x+1:y] + string(str[x]) + str[y+1:]
	}
	for len(queue) > 0 {
		l := len(queue)
		ans++
		for k := 0; k < l; k++ {
			s := queue[k]
			if s == target {
				return ans
			}
			zeroIdx := 0
			for o, c := range []byte(s) {
				if c == '0' {
					zeroIdx = o
					break
				}
			}
			for _, p := range neighbors[zeroIdx] {
				newS := swapStr(s, p, zeroIdx)
				if !visited[newS] {
					visited[newS] = true
					queue = append(queue, newS)
				}
			}
		}
		queue = queue[l:]
	}
	return -1
}
