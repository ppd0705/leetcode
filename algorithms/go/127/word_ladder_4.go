package main

func ladderLength4(beginWord string, endWord string, wordList []string) int {
	step := 0
	startQueue := []string{beginWord}
	endQueue := []string{endWord}
	startUsed := make([]bool, len(wordList))
	endUsed := make([]bool, len(wordList))
	if i := idxOf(endWord, wordList); i == -1 {
		return 0
	} else {
		endUsed[i] = true
	}
	for len(startQueue) > 0 {
		step++
		l := len(startQueue)
		for i := 0; i < l; i++ {
			for j, w := range wordList {
				if !startUsed[j] && hasOneDiff(startQueue[i], w) {
					startQueue = append(startQueue, w)
					startUsed[j] = true
					if endUsed[j] {
						return step+1
					}
				}
			}
		}
		startQueue = startQueue[l:]
		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
			startUsed, endUsed = endUsed, startUsed
		}
	}
	return 0
}
