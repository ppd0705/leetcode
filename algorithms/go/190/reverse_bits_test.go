package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		num    uint32
		target uint32
	}{
		{43261596, 964176192},
		{4294967293, 3221225471},
	}

	for _, tt := range tests {
		if ans := reverseBits(tt.num); ans != tt.target {
			t.Fatalf("expect %d, got %d\n", tt.target, ans)
		}
	}
}
