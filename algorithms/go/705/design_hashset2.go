package main

import "container/list"

const SIZE2 = 769

type MyHashSet2 struct {
	data []list.List
}

/** Initialize your data structure here. */
func Constructor2() MyHashSet2 {
	return MyHashSet2{data: make([]list.List, SIZE2)}
}

func (this *MyHashSet2) Add(key int) {
	i := key % SIZE2
	if !this.Contains(key) {
		this.data[i].PushBack(key)
	}
}


func (this *MyHashSet2) Remove(key int) {
	i := key % SIZE2
	for e := this.data[i].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			this.data[i].Remove(e)
			return
		}
	}
}

/** Returns true if this data contains the specified element */
func (this *MyHashSet2) Contains(key int) bool {
	i := key % SIZE2
	for e := this.data[i].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet2 object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
