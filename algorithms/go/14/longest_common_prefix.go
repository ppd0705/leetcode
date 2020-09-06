package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var s []byte

	out:
	for i := 0; ; i++ {
		if i >= len(strs[0]) {
			break
		}
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != c {
				break out
			}
		}
		s = append(s, c)
	}
	return string(s)
}
