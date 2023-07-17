package main

import "testing"

func TestHello(t *testing.T) {
	t.Run(("Saying hello to people"), func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

	})

	t.Run(("Saying hello world"), func(t *testing.T) {

		got := Hello("", "english")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)

	})

	t.Run(("In spanish"), func(t *testing.T) {
		got := Hello("Dora", "spanish")
		want := "Ola, Dora"
		assertCorrectMessage(t, got, want)
	})

	t.Run(("In french"), func(t *testing.T) {
		got := Hello("Jeanne", "french")
		want := "Bonjour, Jeanne"
		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
