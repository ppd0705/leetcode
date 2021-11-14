package main

import "testing"

func TestMapSum(t *testing.T) {
	m := Constructor()
	if got := m.Sum("apple1"); got != 0 {
		t.Fatalf("expect 0, got %d\n", got)
	}
	m.Insert("apple", 1)
	if got := m.Sum("apple1"); got != 0 {
		t.Fatalf("expect 0, got %d\n", got)
	}
	if got := m.Sum("ap"); got != 1 {
		t.Fatalf("expect 1, got %d\n", got)
	}
	m.Insert("apple", 10)
	if got := m.Sum("ap"); got != 10 {
		t.Fatalf("expect 10, got %d\n", got)
	}
	m.Insert("bapple", 10)
	m.Insert("appq", 30)
	if got := m.Sum("ap"); got != 40 {
		t.Fatalf("expect 40, got %d\n", got)
	}
}
