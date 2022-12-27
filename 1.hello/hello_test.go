package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Guilherme", "")
		want := "Hello, Guilherme"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Guilherme", "Portuguese")
		want := "Olá, Guilherme"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese hen an empty string is supplied", func(t *testing.T) {
		got := Hello("", "Portuguese")
		want := "Olá, Mundo"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}