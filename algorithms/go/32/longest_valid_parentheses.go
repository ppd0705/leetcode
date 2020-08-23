package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func longestValidParentheses(s string) int {
	left, right, maxLen := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		}else{
			right++
		}

		if left == right {
			maxLen = max(maxLen, 2 * left)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s)-1; i >=0; i-- {
		if s[i] == '(' {
			left++
		}else{
			right++
		}

		if left == right {
			maxLen = max(maxLen, 2 * left)
		} else if right < left {
			left, right = 0, 0
		}
	}
	return maxLen
}
