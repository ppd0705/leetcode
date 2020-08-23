package main

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	pre, cur := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				cur, pre = pre, cur
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
			cur, pre = cur+pre, cur
		} else {
			pre = cur
		}
	}
	return cur
}
