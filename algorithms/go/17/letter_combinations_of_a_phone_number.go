package main

func letterCombinations(digits string) []string {
	numberMap := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	ans := make([]string, 0)
	var dfs func(s string, idx int)
	dfs = func(s string, idx int) {
		if idx == len(digits) {
			ans = append(ans, s)
			return
		}
		for _, x := range numberMap[string(digits[idx])] {
			dfs(s+x, idx+1)
		}
	}
	if len(digits) > 0 {
		dfs("", 0)
	}
	return ans
}
