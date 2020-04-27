package three_sum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := [][]int{
		{1, 0, 100, 1},
		{0, 0, 0},
		{1, 0, -1, 1},
		{1, 0, -1, 1, -3, -6, 9},
		{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0},
	}

	targets := [][][]int{
		nil,
		{{0, 0, 0}},
		{
			{-1, 0, 1},
		},
		{
			{-6, -3, 9},
			{-1, 0, 1},
		},
		{
			{-5, 1, 4},
			{-4, 0, 4},
			{-4, 1, 3},
			{-2, -2, 4},
			{-2, 1, 1},
			{0, 0, 0},
		},
	}

	for i := 0; i < len(targets); i++ {
		ret := threeSum(tests[i])
		fmt.Printf("nums：%v, target: %v, result: %v\n", tests[i], targets[i], ret)
		if ! reflect.DeepEqual(ret, targets[i]) {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
