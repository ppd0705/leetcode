package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		board  [][]byte
		click  []int
		expect [][]byte
	}{
		{[][]byte{
			{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'M', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'}},
			[]int{1, 2},
			[][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'X', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'}},
		},
		{[][]byte{
			{'E', 'E', 'E', 'E', 'E'},
			{'E', 'E', 'M', 'E', 'E'},
			{'E', 'E', 'E', 'E', 'E'},
			{'E', 'E', 'E', 'E', 'E'}},
			[]int{3, 0},
			[][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'}},
		},
	}
	for _, tt := range tests {
		if ans := updateBoard2(tt.board, tt.click); !reflect.DeepEqual(ans, tt.expect) {
			t.Fatalf("expect: %+v \ngot %+v \n", tt.expect, ans)
		}
	}
}
