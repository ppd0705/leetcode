package main

import "math/rand"

type Solution2 struct {
	m, n, total int
	points      map[int]int
}

func Constructor2(m int, n int) Solution2 {
	return Solution2{
		m:      m,
		n:      n,
		total:  m * n,
		points: make(map[int]int),
	}
}

func (this *Solution2) Flip() []int {
	var ans []int
	x := rand.Intn(this.total)
	this.total--
	if y, ok := this.points[x]; ok {
		ans = []int{y / this.n, y % this.n}
	} else {
		ans = []int{x / this.n, x % this.n}
	}
	if y, ok := this.points[this.total]; ok {
		this.points[x] = y
	} else {
		this.points[x] = this.total
	}
	return ans
}

func (this *Solution2) Reset() {
	this.points = map[int]int{}
	this.total = this.m * this.n
}
