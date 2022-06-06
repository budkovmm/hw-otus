package main

import "testing"

func TestReverse(t *testing.T) {
	want := "!SUTO ,olleH"
	got := ReverseString("Hello, OTUS!")

	if got != want {
		t.Errorf("Expected: %q, but actual: %q", want, got)
	}
}
