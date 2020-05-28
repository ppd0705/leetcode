package main

var mutationMap = map[uint8][3]string{
	'A': [...]string{"T", "G", "C"},
	'C': [...]string{"T", "G", "A"},
	'T': [...]string{"A", "G", "C"},
	'G': [...]string{"T", "A", "C"},
}

func idxOf(str string, bank []string) int {
	for i, s := range bank {
		if s == str {
			return i
		}
	}
	return -1
}

func minMutation(start string, end string, bank []string) int {
	if idxOf(end, bank) == -1 {
		return -1
	}
	isUsed := make([]bool, len(bank))

	queue := []string{start}
	count := 0
	for len(queue) > 0 {
		l := len(queue)
		for i:= 0; i < l; i++ {
			curr := queue[i]
			if curr == end {
				return count
			}
			for j:= 0;j< len(curr);j++ {
				for _, s := range mutationMap[curr[j]] {
					if idx := idxOf(curr[:j] + s + curr[j+1:], bank); idx != -1 && !isUsed[idx] {
						queue = append(queue, bank[idx])
						isUsed[idx] = true
					}
				}

			}

		}
		count++
		queue = queue[l:]
	}
	return -1
}
