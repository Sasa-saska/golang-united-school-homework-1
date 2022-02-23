package main

import (
	"testing"
)

func Test_hello(t *testing.T) {
	var got = hello("Chris")
	var want = "Hello, Chris"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
