package main

import (
	"fmt"
	"testing"
)

func Test_Subsets(t *testing.T) {
	ans := subsets([]int{1,2,3})
	fmt.Println(ans)
}