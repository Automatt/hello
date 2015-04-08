package main

import "testing"

func TestGetHello(t *testing.T) {
	want := "hello, world!"
	got := getHello()
	if want != got {
		t.Errorf("Expected: %q, Got: %q", want, got)
	}
}