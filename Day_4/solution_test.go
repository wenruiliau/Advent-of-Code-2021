package main

import "testing"

func TestAnswer1(t *testing.T) {
	got := Answer1("test.txt")
	want := 4512

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestAnswer2(t *testing.T) {
	got := Answer2("test.txt")
	want := 1924

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
