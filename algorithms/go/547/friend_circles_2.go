package main

var count int

func findCircleNum2(M [][]int) int {
	n := len(M)
	count = n
	p := make([]int, n)
	for i := 1; i < n; i++ {
		p[i] = i
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if M[i][j] == 1 {
				union(p, i, j)
			}
		}
	}
	return count
}

func union(p []int, x, y int) {
	px, py := find(p, x), find(p, y)
	if px != py {
		p[px] = py
		count -= 1
	}
}

func find(p []int, x int) int {
	for p[x] != x {
		p[x] = p[p[x]]
		x = p[x]
	}
	return x
}
