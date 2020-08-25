package main

func firstUniqChar0(s string) int {
	if len(s) == 0 {
		return -1
	}
	charIdxMap := make(map[rune]int, 0)
	cntArr := make([]int, len(s))
	for i, r := range s {
		if _, ok := charIdxMap[r]; !ok {
			charIdxMap[r] = i
		}
		cntArr[charIdxMap[r]]++
	}
	for i, v := range cntArr {
		if v == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar1(s string) int {
	if len(s) == 0 {
		return -1
	}
	charCountMap := make(map[rune]int, 0)
	for _, r := range s {
		charCountMap[r]++
	}
	for i, r := range s {
		if charCountMap[r] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}

	var arr [26]int

	for i, r := range s {
		arr[r-97] = i
	}

	for i, r := range  s {
		if arr[r-97] == i {
			return i
		} else {
			arr[r-97] = -1
		}
	}
	return -1
}
