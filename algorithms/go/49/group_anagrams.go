package main

import (
	"bytes"
	"fmt"
)

func encodeString(s string) string {
	chars := make([]int, 26)
	for i := 0; i < len(s); i++ {
		chars[s[i] - 'a']++
	}
	encoded := &bytes.Buffer{}
	for i := 0; i < len(chars); i++ {
		encoded.WriteString(fmt.Sprintf("%d_%d", i, chars[i]))

	}
	return encoded.String()
}

func groupAnagrams0(strs []string) [][]string {
	groupMap := make(map[string][]string)

	for _, s := range strs {
		se := encodeString(s)
		groupMap[se] = append(groupMap[se], s)
	}

	groups := make([][]string, 0)

	for _, v := range groupMap {
		groups = append(groups, v)
	}

	return groups
}

func main() {
	src := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	dst := groupAnagrams0(src)

	fmt.Println("dst: ", dst)
}
