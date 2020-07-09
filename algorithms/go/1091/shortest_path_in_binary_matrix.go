package main

type point struct {
	x int
	y int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	row := len(grid)
	column := len(grid[0])

	if grid[0][0] != 0 || grid[row-1][column-1] != 0 {
		return -1
	}

	queue := []*point{&point{0, 0}}
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, row)
	}
	visited[0][0] = true
	ans := 0

	offsets := [8]point{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{-1, 1},
		{-1, -1},
		{1, 1},
		{1, -1},
	}
	for len(queue) > 0 {
		ans++
		l := len(queue)
		for k := 0; k < l; k++ {
			p := queue[k]
			if p.x == row-1 && p.y == column-1 {
				return ans
			}
			for _, o := range offsets {
				x, y := p.x+o.x, p.y+o.y
				if x < 0 || x >= row || y < 0 || y >= column || grid[x][y] == 1 || visited[x][y] {
					continue
				}
				visited[x][y] = true
				queue = append(queue, &point{x, y})
			}
		}
		queue = queue[l:]
	}
	return -1
}
