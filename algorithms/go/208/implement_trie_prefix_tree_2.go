package main

type Trie2 struct {
	children [26]*Trie2
	isEnd    bool
}

/** Initialize your data structure here. */
func Constructor2() Trie2 {
	return Trie2{}
}

/** Inserts a word into the Trie2. */
func (this *Trie2) Insert(word string) {
	for _, w := range word {
		idx := w - 'a'
		if this.children[idx] == nil {
			this.children[idx] = &Trie2{}
		}
		this = this.children[idx]
	}
	this.isEnd = true
}

/** Returns if the word is in the Trie2. */
func (this *Trie2) Search(word string) bool {
	for _, w := range word {
		idx := w - 'a'
		if this.children[idx] == nil {
			return false
		}
		this = this.children[idx]
	}
	return this.isEnd
}

/** Returns if there is any word in the Trie2 that starts with the given prefix. */
func (this *Trie2) StartsWith(prefix string) bool {
	for _, w := range prefix {
		idx := w - 'a'
		if this.children[idx] == nil {
			return false
		}
		this = this.children[idx]
	}
	return true
}
