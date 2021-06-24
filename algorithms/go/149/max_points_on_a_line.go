package main

func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}
	ans := 0
	for i, p := range points {
		if ans >= n - i {
			break
		}
		cnt := make(map[int]int, 0)
		for j := i+1; j < n; j++ {
			x, y := points[j][0] - p[0], points[j][1] - p[1]
			if x == 0 {
				y = 1
			} else if y == 0 {
				x = 1
			} else {
				// let y always be a positive int
				if y < 0 {
					y = -y
					x = -x
				}
				g := gcd(abs(x), abs(y))
				y /= g
				x /= g
			}
			// max x: 10^4 * 2
			cnt[x*20001+y] += 1
		}
		for _, v := range cnt {
			ans = max(ans, v+1)
		}
	}
	return ans
}


func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// greatest common divisor
func gcd(x, y int) int {
	for x != 0 {
		x , y =  y % x, x
	}
	return y
}