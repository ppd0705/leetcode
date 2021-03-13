package main

const (
	NULL    = -1
	DELETED = -2
	SIZE    = 13333 // 10000 / 0.75

)

type MyHashSet struct {
	data []int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	set := make([]int, SIZE)
	for i := 0; i < SIZE; i++ {
		set[i] = NULL
	}
	return MyHashSet{data: set}
}

func (this *MyHashSet) Add(key int) {
	i := this.hash(key)
	if this.data[i] == NULL {
		this.data[i] = key
	}
}

func (this *MyHashSet) hash(key int) int {
	i := key % SIZE
	for this.data[i] != key && this.data[i] != NULL {
		i = (i + 1) % SIZE
	}
	return i
}

func (this *MyHashSet) Remove(key int) {
	i := this.hash(key)
	if this.data[i] == key {
		this.data[i] = DELETED
	}
}

/** Returns true if this data contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.data[this.hash(key)] == key
}

/**
 * Your MyHashSet2 object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
