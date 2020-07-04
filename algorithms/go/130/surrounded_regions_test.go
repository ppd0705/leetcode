package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		board  [][]byte
		expect [][]byte
	}{
		{[][]byte{
			{'x', 'x'},
		}, [][]byte{
			{'x', 'x'},
		},
		},
		{[][]byte{
			{'x', 'x', 'x', 'x'},
			{'x', 'o', 'o', 'x'},
			{'x', 'x', 'o', 'x'},
			{'x', 'o', 'x', 'x'},
		}, [][]byte{
			{'x', 'x', 'x', 'x'},
			{'x', 'x', 'x', 'x'},
			{'x', 'x', 'x', 'x'},
			{'x', 'o', 'x', 'x'},
		},
		},
		{[][]byte{
			{'x', 'x', 'x', 'x'},
			{'x', 'o', 'o', 'x'},
			{'x', 'x', 'o', 'x'},
			{'x', 'o', 'o', 'x'},
		}, [][]byte{
			{'x', 'x', 'x', 'x'},
			{'x', 'o', 'o', 'x'},
			{'x', 'x', 'o', 'x'},
			{'x', 'o', 'o', 'x'},
		},
		},
	}

	for _, tt := range tests {
		if solve2(tt.board); !reflect.DeepEqual(tt.board, tt.expect) {
			t.Fatalf("expect %s\ngot %s\n", tt.expect, tt.board)
		}
	}
}
