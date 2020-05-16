package main

import "testing"

func TestClimbingStairs(t *testing.T) {
	params := map[int]int{1: 1, 2: 2, 3: 3}

	for step, res := range params {
		if n := climbStairs2(step); n != res {
			t.Fatalf("step %d expect: %d", n, res)
		}
	}

}
