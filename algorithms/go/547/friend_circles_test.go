package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		M      [][]int
		expect int
	}{
		{[][]int{{1, 1}, {1, 1}}, 1},
	}

	for _, tt := range tests {
		if ans := findCircleNum(tt.M); ans != tt.expect {
			t.Fatalf("expect %d got %d\n", tt.expect, ans)
		}
	}
}
