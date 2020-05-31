package main

func ladderLength3(beginWord string, endWord string, wordList []string) int {
	if idxOf(endWord, wordList) == -1 {
		return 0
	}
	step := 0
	startQueue := []string{beginWord}
	endQueue := []string{endWord}
	used := make([]bool, len(wordList))

	for len(startQueue) > 0 {
		step++
		l := len(startQueue)
		for i := 0; i < l; i++ {
			for k := 0; k < len(endQueue); k++ {
				if hasOneDiff(startQueue[i], endQueue[k]) {
					return step + 1
				}
			}
			for j, w := range wordList {
				if !used[j] && hasOneDiff(startQueue[i], w) {
					startQueue = append(startQueue, w)
					used[j] = true
				}
			}
		}
		startQueue = startQueue[l:]
		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
		}
	}
	return 0
}
