package main


// timeout  233
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	if idxOf(endWord, wordList) == -1 {
		return 0
	}
	minStep := len(wordList) + 2
	used := make([]bool, len(wordList))
	var dfs func(curr string, level int)
	dfs = func(curr string, level int) {
		if curr == endWord {
			if level < minStep {
				minStep = level
			}
			return
		}

		for i, w := range wordList {
			if !used[i] && hasOneDiff(w, curr) {
				used[i] = true
				dfs(w, level+1)
				used[i] = false
			}
		}
	}
	dfs(beginWord, 1)
	if minStep <= len(wordList) + 1{
		return minStep
	}
	return 0
}
