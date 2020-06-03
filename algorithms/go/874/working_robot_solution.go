package main

type pointer struct {
	x int
	y int
}

func robotSim(commands []int, obstacles [][]int) int {
	x, y, dx, dy := 0, 0, 0, 1
	obstacleMap := make(map[pointer]bool)
	for _, o := range obstacles {
		obstacleMap[pointer{o[0], o[1]}] = true
	}
	nextX, nextY := 0, 0
	maxDistance := 0
	for _, c := range commands {
		if c == -2 {
			dx, dy = -dy, dx
		} else if c == -1 {
			dx, dy = dy, -dx
		} else {
			for i := 0; i < c; i++ {
				nextX = x + dx
				nextY = y + dy
				if obstacleMap[pointer{nextX, nextY}] {
					break
				}
				x = nextX
				y = nextY
			}

			distance := x*x + y*y
			if distance > maxDistance {
				maxDistance = distance
			}

		}

	}
	return maxDistance
}
