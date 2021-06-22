package main

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch2(turnedOn int) []string {
	ans := make([]string, 0)
	for i := 0; i < 1024; i++ {
		hour := i >> 6
		minute := i & 63
		if hour < 12 && minute < 60 && bits.OnesCount(uint(i)) == turnedOn {
			ans = append(ans, fmt.Sprintf("%d:%02d", hour, minute))
		}
	}
	return ans
}
