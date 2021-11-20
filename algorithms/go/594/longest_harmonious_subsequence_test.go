package main

import "testing"

func Test_findLHS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{[]int{1, 1}}, 0},
		{"test1", args{[]int{1, 1, 2}}, 3},
		{"test2", args{[]int{1, 1, 3}}, 0},
		{"test3", args{[]int{1, 1, 2, 3, 3, 3}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLHS(tt.args.nums); got != tt.want {
				t.Errorf("findLHS() = %v, want %v", got, tt.want)
			}
		})
	}
}
