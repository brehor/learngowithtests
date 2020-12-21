package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Brett")
		want := "Hello, Brett"

		assertCorrectMessage(t, got, want)
	})
	t.Run("default to 'Hello, world!' if empty string is passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})
}
