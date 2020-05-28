package main

func minMutation2(start string, end string, bank []string) int {
	if idxOf(end, bank) == -1 {
		return -1
	}
	minCount := len(bank) + 1
	isUsed := make([]bool, len(bank))

	var dfs func(curr string, count int)
	dfs = func(curr string, count int) {
		if count >= minCount {
			return
		}
		if curr == end {
			if count < minCount {
				minCount = count
			}
			return
		}

		for i := 0; i < len(curr); i++ {
			for _, c := range mutationMap[curr[i]] {
				gene := curr[:i] + c + curr[i+1:]
				if idx := idxOf(gene, bank); idx != -1 && !isUsed[idx] {
					isUsed[idx] = true
					dfs(gene, count+1)
					isUsed[idx] = false
				}
			}
		}
	}
	dfs(start, 0)
	if minCount <= len(bank) {
		return minCount
	}
	return -1
}
