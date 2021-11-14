package main

import (
	"sort"
	"strings"
)

type Item struct {
	key   string
	value int
}

type Items []Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	if items[i].key < items[j].key {
		return true
	}
	return false
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

type MapSum struct {
	items Items
}

func Constructor() MapSum {
	return MapSum{Items{}}
}

func (this *MapSum) Insert(key string, val int) {
	idx := sort.Search(len(this.items), func(i int) bool {
		return this.items[i].key >= key
	})
	if idx == len(this.items) {
		this.items = append(this.items, Item{key, val})
		return
	}
	if this.items[idx].key == key {
		this.items[idx].value = val
		return
	}

	this.items = append(this.items, Item{key, val})
	sort.Sort(this.items)

}

func (this *MapSum) Sum(prefix string) int {
	sum := 0
	idx := sort.Search(len(this.items), func(i int) bool {
		return this.items[i].key >= prefix
	})
	for i := idx; i < len(this.items); i++ {
		if !strings.HasPrefix(this.items[i].key, prefix) {
			break
		}
		sum += this.items[i].value
	}
	return sum
}
