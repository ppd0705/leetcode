package main

func getCoordinate(x, n int) (int, int) {
	r, c := (x-1)/n, (x-1)%n
	if r%2 == 1 {
		c = n - 1 - c
	}
	r = n - 1 - r
	return r, c
}

func snakesAndLadders(board [][]int) int {
	step := 0
	n := len(board)
	target := n * n
	start := 1
	queue := []int{start}
	visited := make([]bool, target+1)

	getBoardValue := func(x int) int {
		r, c := (x-1)/n, (x-1)%n
		if r%2 == 1 {
			c = n - 1 - c
		}
		r = n - 1 - r
		return board[r][c]
	}

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			cur := queue[i]
			for j := 1; j <= 6 && cur+j <= target; j++ {
				next := cur + j
				nextValue := getBoardValue(next)
				if nextValue != -1 {
					next = nextValue
				}
				if next == target {
					return step + 1
				}
				if !visited[next] {
					queue = append(queue, next)
					visited[next] = true
				}
			}

		}
		step += 1
		queue = queue[l:]
	}
	return -1
}
