package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := [26]int{}
	for i, j := range s {
		counter[j-'a']++
		counter[t[i]-'a']--
	}
	for _, j := range counter {
		if j != 0 {
			return false
		}
	}
	return true
}

func main() {
	a := "hello"
	b := "llohe"
	r := isAnagram(a, b)
	fmt.Println("ret: ", r)
}
