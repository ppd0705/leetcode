package main

import (
	"fmt"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	all := make([]string, 0)
	for i := 2; i <= 2; i++ {
		ans := readBinaryWatch2(i)
		all = append(all, ans...)
	}
	sort.Strings(all)
	fmt.Printf("ans: %v\n", all)
}