package cmd

import (
	"testing"
)

func TestDisplayUserManual(t *testing.T) {
	got := DisplayGreeting()
	want := UserManual

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDisplayGreeting(t *testing.T) {

}

