package main

import "sort"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	merged := make([][]int, 0)

	for _, v := range intervals {
		if len(merged) == 0 || merged[len(merged)-1][1] < v[0] {
			merged = append(merged, v)
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], v[1])
		}
	}
	return merged
}
