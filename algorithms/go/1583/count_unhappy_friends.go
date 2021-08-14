package main

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	pairMap := make(map[int]int, n)
	for _, pair := range pairs {
		pairMap[pair[0]] = pair[1]
		pairMap[pair[1]] = pair[0]
	}
	res := 0

	for x := 0; x < n; x++ {
		y := pairMap[x]
	OUTER:
		for _, u := range preferences[x] {
			if u == y {
				break
			}
			v := pairMap[u]
			for _, i := range preferences[u] {
				if i == x {
					res += 1
					break OUTER
				} else if i == v {
					break
				}
			}
		}
	}
	return res
}
