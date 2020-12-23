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
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("specifying Spanish returns Spanish greeting", func(t *testing.T) {
		got := Hello("Macy", "Spanish")
		want := "Hola, Macy"
		assertCorrectMessage(t, got, want)
	})

	t.Run("specifying French returns French greeting", func(t *testing.T) {
		got := Hello("Mindy", "French")
		want := "Bonjour, Mindy"
		assertCorrectMessage(t, got, want)
	})

	t.Run("specifying Italian returns Italian greeting", func(t *testing.T) {
		got := Hello("Mario", "Italian")
		want := "Ciao, Mario"
		assertCorrectMessage(t, got, want)
	})
}
