package main

import "math"

func myAtoi(str string) int {
	i := 0
	sign := 1
	idx := 0
	for idx < len(str) && str[idx] == ' ' {
		idx++
	}
	if idx == len(str) {
		return 0
	}
	if str[idx] == '-' {
		sign = -1
		idx++
	} else if str[idx] == '+' {
		idx++
	}

	for ;idx < len(str) && str[idx] >= '0' && str[idx] <= '9';idx++ {
		if sign == 1 {
			if i > math.MaxInt32/10 || i == math.MaxInt32/10 && str[idx] >= '7' {
				return math.MaxInt32
			}
			i = i*10 + int(str[idx]) - 48
		} else {
			if i < math.MinInt32/10 || i == math.MinInt32/10 && str[idx] >= '8' {
				return math.MinInt32
			}
			i = i*10 - int(str[idx]) + 48
		}
	}
	return i
}
