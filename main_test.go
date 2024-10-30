package main

import (
	"testing"
)

func TestGoDoSaysHello(t *testing.T) {
	got := OpeningMessage()
	want := "The first Go Code!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
