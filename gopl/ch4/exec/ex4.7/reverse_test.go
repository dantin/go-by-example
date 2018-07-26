package main

import "testing"

func TestRevUTF8(t *testing.T) {
	s := []byte("深入理解计算机系统")
	want := "统系机算计解理入深"
	got := string(revUTF8(s))
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
