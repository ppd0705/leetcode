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

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if idxOf(endWord, wordList) == -1 {
		return 0
	}
	step := 0
	used := make([]bool, len(wordList))
	queue := []string{beginWord}
	for len(queue) > 0 {
		step++
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i] == endWord {
				return step
			}
			for j, w := range wordList {
				if !used[j] && hasOneDiff(queue[i], w) {
					queue = append(queue, w)
					used[j] = true
				}
			}
		}
		queue = queue[l:]
	}
	return 0
}
