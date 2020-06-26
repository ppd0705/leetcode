package main

import "fmt"

type Trie struct {
	children map[rune]*Trie
	isEnd    bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	fmt.Println(233)
	return Trie{map[rune]*Trie{}, false}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, w := range word {
		if _, ok := this.children[w]; !ok {
			this.children[w] = &Trie{map[rune]*Trie{}, false}
		}
		this = this.children[w]
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, w := range word {
		if _, ok := this.children[w]; !ok {
			return false
		}
		this = this.children[w]
	}
	return this.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, w := range prefix {
		if _, ok := this.children[w]; !ok {
			return false
		}
		this = this.children[w]
	}
	return true
}
