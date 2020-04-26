package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	charArr := strings.Split(s, "")
	sort.Strings(charArr)
	return strings.Join(charArr, "")
}

func groupAnagrams3(strs []string) [][]string {

	sorted := make([]string, len(strs))
	for i, s:= range strs {
		sorted[i] = sortString(s)
	}
	groupMap := make(map[string][]string)

	for i, s := range sorted {
		groupMap[s] = append(groupMap[s], strs[i])
	}

	groups := make([][]string, 0)

	for _, v := range groupMap {
		groups = append(groups, v)
	}

	return groups
}

func main() {
	src := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	dst := groupAnagrams3(src)

	fmt.Println("dst: ", dst)
}
