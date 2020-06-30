package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
	{'e', 't', 'a', 'e'},
	{'i', 'h', 'k', 'r'},
	{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain", "ea"}
	expect := []string{"oath", "ea", "eat"}

	if ans := findWords(board, words); !reflect.DeepEqual(expect, ans) {
		t.Fatalf("expect %v, got %v\n", expect, ans)
	}
}