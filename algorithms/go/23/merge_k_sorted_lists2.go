package main

type loserTree struct {
	tree   []int
	leaves []*ListNode
}

var  dummyVal = 100000
var  dummyListNode = ListNode{Val: dummyVal}

func New(leaves []*ListNode) *loserTree {
	k := len(leaves)
	if k & 1 == 1  {
		leaves = append(leaves, &dummyListNode)
		k++
	}
	lt := &loserTree{
		tree:   make([]int, k),
		leaves: leaves,
	}
	if k > 0 {
		lt.init()
	}
	return lt
}

func (lt *loserTree) init() {
	lt.tree[0] = lt.getWinner(0)
}

func (lt *loserTree) getWinner(idx int) int {
	if idx == 0 {
		return lt.getWinner(1)
	}
	if idx >= len(lt.tree) {
		return idx - len(lt.tree)
	}
	left := lt.getWinner(idx*2)
	right := lt.getWinner(idx*2+1)
	if lt.leaves[left] == nil {
		lt.leaves[left] = &dummyListNode
	}
	if lt.leaves[right] == nil {
		lt.leaves[right] = &dummyListNode
	}
	if lt.leaves[left].Val < lt.leaves[right].Val {
		left, right = right, left
	}
	lt.tree[idx] = left
	return right
}


func (lt *loserTree) Pop() *ListNode {
	if len(lt.tree) == 0 {
		return &dummyListNode
	}
	treeWinner := lt.tree[0]
	winner := lt.leaves[treeWinner]
	lt.leaves[treeWinner] = winner.Next
	if lt.leaves[treeWinner] == nil {
		lt.leaves[treeWinner] = &dummyListNode
	}
	treeIdx := (treeWinner + len(lt.tree)) / 2
	for treeIdx != 0 {
		treeLoser := lt.tree[treeIdx]
		if lt.leaves[treeLoser].Val  < lt.leaves[treeWinner].Val  {
			treeLoser, treeWinner  = treeWinner, treeLoser
		}
		lt.tree[treeIdx] = treeLoser
		treeIdx /= 2
	}
	lt.tree[0] = treeWinner
	return winner
}

func mergeKLists2(lists []*ListNode) *ListNode {
	dummy := new(ListNode)
	pre := dummy
	lt := New(lists)
	for {
		node := lt.Pop()
		if node == &dummyListNode {
			break
		}
		pre.Next = node
		pre = node
	}
	return dummy.Next
}
