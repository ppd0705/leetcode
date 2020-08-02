package main

import "testing"

func Test(t *testing.T) {
	lru := Constructor(2)
	if ans := lru.Get(1); ans != -1 {
		t.Fatalf("hit error, got %d, expect %d", ans, -1)
	}
	lru.Put(1, 1)
	if ans := lru.Get(1); ans != 1 {
		t.Fatalf("hit error, got %d, expect %d", ans, 1)
	}
	lru.Put(1, 2)
	if ans := lru.Get(1); ans != 2 {
		t.Fatalf("hit error, got %d, expect %d", ans, 2)
	}
	lru.Put(2, 2)
	lru.Put(3, 2)
	if ans := lru.Get(1); ans != -1 {
		t.Fatalf("hit error, got %d, expect %d", ans, -1)
	}
}
