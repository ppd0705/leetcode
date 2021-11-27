package main

import "math/rand"

type point struct {
	x, y int
}

type Solution struct {
	m, n   int
	points map[point]bool
}

func Constructor(m int, n int) Solution {
	return Solution{
		m:      m,
		n:      n,
		points: make(map[point]bool),
	}
}

func (this *Solution) Flip() []int {
	var i, j int
	for {
		i = rand.Intn(this.m)
		j = rand.Intn(this.n)
		if !this.points[point{i, j}] {
			this.points[point{i, j}] = true
			break
		}
	}
	return []int{i, j}
}

func (this *Solution) Reset() {
	this.points = map[point]bool{}
}
