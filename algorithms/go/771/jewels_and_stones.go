package main

func numJewelsInStones(J string, S string) int {
	jSet := make(map[rune]bool, 0)

	for _, r := range J {
		jSet[r] = true
	}
	cnt := 0

	for _, r := range S {
		if jSet[r] {
			cnt++
		}
	}
	return cnt
}
