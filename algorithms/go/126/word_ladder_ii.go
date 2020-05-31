package main

func idxOf(str string, bank []string) int {
	for i, s := range bank {
		if str == s {
			return i
		}
	}
	return -1
}

func hasOneDiff(x, y string) bool {
	count := 0
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	if count == 1 {
		return true
	}
	return false
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	ans := make([][]string, 0)
	if idxOf(endWord, wordList) == -1 {
		return ans
	}
	minStep := len(wordList) + 2
	used := make([]bool, len(wordList))
	var dfs func(curr []string)
	dfs = func(curr []string) {
		if len(curr) > minStep {
			return
		}
		if curr[len(curr)-1] == endWord {
			ans = append(ans, append([]string{}, curr...))
			if len(curr) < minStep {
				ans = ans[len(ans)-1:]
				minStep = len(curr)
			}
			return
		}

		for i, w := range wordList {
			if !used[i] && hasOneDiff(curr[len(curr)-1], w) {
				curr = append(curr, w)
				used[i] = true
				dfs(curr)
				used[i] = false
				curr = curr[:len(curr)-1]
			}
		}
	}
	dfs([]string{beginWord})
	return ans
}
