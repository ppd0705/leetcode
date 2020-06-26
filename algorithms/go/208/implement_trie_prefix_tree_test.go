package main

import "testing"

func Test(t *testing.T) {
	w1 := "apple"
	w2 := "app"

	s := Constructor2()
	s.Insert(w1)

	if s.Search(w2) {
		t.Fatalf("error search, exclude %s actually\n", w2)
	}

	if !s.Search(w1) {
		t.Fatalf("error search, not found %s\n", w1)
	}

	if !s.StartsWith(w1) || !s.StartsWith(w2) {
		t.Fatalf("error prefix search, not found\n")
	}

	s.Insert(w2)
	if !s.Search(w2) {
		t.Fatalf("error search, not found %s\n", w2)
	}

}
