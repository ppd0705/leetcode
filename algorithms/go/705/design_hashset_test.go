package main

import "testing"

func Test(t *testing.T) {
	m := Constructor2()
	m.Add(10)
	m.Add(10)
	m.Add(10)
	m.Add(10)
	m.Add(10)
	m.Add(10)
	if !m.Contains(10) {
		t.Fatalf("10 shuld be contained")
	}
	m.Remove(10)
	if m.Contains(10) {
		t.Fatalf("10 shuld not be contained")
	}
	m.Add(11)
	m.Add(1)
	m.Add(100001)
	if !m.Contains(100001) {
		t.Fatalf("100001 shuld be contained")
	}
}
