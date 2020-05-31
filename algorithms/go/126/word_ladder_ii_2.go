package main

func findLadders2(beginWord string, endWord string, wordList []string) [][]string {
	ans := make([][]string, 0)
	if idxOf(endWord, wordList) == -1 {
		return ans
	}
	minStep := len(wordList) + 2
	startQueue := [][]string{{beginWord}}

	for len(startQueue) > 0 {
		l := len(startQueue)
		for i := 0; i < l; i++ {
			curr := startQueue[i]
			if len(curr) >= minStep {
				continue
			}
			for _, w := range wordList {
				if hasOneDiff(w, curr[len(curr)-1]) && idxOf(w, curr) == -1 {
					newSeq := append([]string{}, curr...)
					newSeq = append(newSeq, w)
					if w == endWord {
						ans = append(ans, newSeq)
						if len(newSeq) < minStep {
							minStep = len(newSeq)
							ans = ans[len(ans)-1:]
						}
					} else {
						startQueue = append(startQueue, newSeq)
					}
				}
			}
		}
		startQueue = startQueue[l:]
	}
	return ans
}
