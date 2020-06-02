package main

func lemonadeChange(bills []int) bool {
	bank := make(map[int]int, 2)
	for _, b := range bills {
		if b == 5 {
			bank[5]++
		} else if b == 10 {
			if bank[5] == 0 {
				return false
			} else {
				bank[5]--
				bank[10]++
			}
		} else {
			if bank[10] > 0 && bank[5] > 0 {
				bank[10]--
				bank[5]--
			} else if bank[5] >= 3 {
				bank[5] = bank[5] - 3
			} else {
				return false
			}
		}
	}
	return true
}
