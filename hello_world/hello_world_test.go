package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := hello_world("Nishil", "English")
		want := "Hello, Nishil"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello, World when an empty string is supplied", func(t *testing.T) {
		got := hello_world("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)

	})

	t.Run("in spanish", func(t *testing.T) {
		got := hello_world("Nishil", "Spanish")
		want := "Hola, Nishil"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := hello_world("Nishil", "French")
		want := "Bonjour, Nishil"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
