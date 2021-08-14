package main

func unhappyFriends2(n int, preferences [][]int, pairs [][]int) int {
	match := make([]int, n)
	for _, pair := range pairs {
		match[pair[0]] = pair[1]
		match[pair[1]] = pair[0]
	}

	orders := make([][]int, n)
	for i := 0; i < n; i++ {
		orders[i] = make([]int, n)
		for k, j := range preferences[i] {
			orders[i][j] = k
		}
	}
	res := 0
	for x := 0; x < n; x++ {
		y := match[x]
		for _, u := range preferences[x][:orders[x][y]] {
			v := match[u]
			if orders[u][x] < orders[u][v] {
				res += 1
				break
			}
		}

	}
	return res
}
