package main

func lengthOfLastWord(s string) int {
	cnt := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if cnt == 0 {
				continue
			} else {
				break
			}
		} else {
			cnt++
		}
	}
	return cnt
}
