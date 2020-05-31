package main

func findLadders3(beginWord string, endWord string, wordList []string) [][]string {
	ans := make([][]string, 0)
	if idxOf(endWord, wordList) == -1 {
		return ans
	}
	startQueue := [][]string{{beginWord}}
	endQueue := [][]string{{endWord}}

	found := false
	for len(startQueue) > 0 {
		l := len(startQueue)
		for i := 0; i < l; i++ {
			curr := startQueue[i]
			for _, w := range wordList {
				if hasOneDiff(w, curr[len(curr)-1]){
					for _, end := range endQueue {
						if end[len(end)-1] == w {
							newSeq := make([]string, 0)
							if curr[0] == beginWord {
								newSeq = append(newSeq, curr...)
								for j := len(end) - 1; j >= 0; j-- {
									newSeq = append(newSeq, end[j])
								}
							} else {
								newSeq = append(newSeq, end...)
								for j := len(curr) - 1; j >= 0; j-- {
									newSeq = append(newSeq, curr[j])
								}
							}
							ans = append(ans, newSeq)
							found = true
						}
					}
					if !found {
						startQueue = append(startQueue, append([]string{}, curr...))
						startQueue[len(startQueue)-1] = append(startQueue[len(startQueue)-1], w)
					}

				}
			}
		}
		if found {
			break
		}
		startQueue = startQueue[l:]
		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
		}
	}
	return ans
}
