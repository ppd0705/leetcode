package main

type DNode struct {
	key   int
	value int
	prev  *DNode
	next  *DNode
}

type LRUCache struct {
	cap   int
	cache map[int]*DNode
	root  *DNode
}

func Constructor(capacity int) LRUCache {
	root := &DNode{}
	root.prev = root
	root.next = root
	return LRUCache{cap: capacity, root: root, cache: make(map[int]*DNode)}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; !ok {
		return -1
	} else {
		this.moveToEnd(node)
		return node.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToEnd(node)
	} else {
		if len(this.cache) < this.cap {
			node := &DNode{
				key:   key,
				value: value,
				next:  this.root,
				prev:  this.root.prev,
			}
			this.root.prev.next = node
			this.root.prev = node

			this.cache[key] = node
		} else {
			this.root.key = key
			this.root.value = value
			this.cache[key] = this.root
			this.root = this.root.next
			delete(this.cache, this.root.key)

		}
	}

}

func (this *LRUCache) moveToEnd(node *DNode) {
	node.next.prev, node.prev.next = node.prev, node.next
	node.prev, node.next = this.root.prev, this.root
	this.root.prev.next, this.root.prev = node, node
}
