package main

import (
	"fmt"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	groupMap := make(map[string][]string)
	groups := make([][]string, 0)

	for _, s := range strs {
		sByte := sortRunes(s)
		sort.Sort(sByte)
		group := groupMap[string(sByte)]
		groupMap[string(sByte)] = append(group, s)
	}
	for _, v := range groupMap {
		groups = append(groups, v)
	}
	return groups
}

func main() {


	src := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	dst := groupAnagrams(src)

	fmt.Println("dst: ", dst)
}
