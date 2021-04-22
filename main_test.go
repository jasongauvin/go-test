package main

import "testing"

func Test_main(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }

	t.Run("Saying Hello to people", func(t *testing.T) {
		got := Hello("Marcel", "")
		want := "Hello, Marcel"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hello world when not name specified", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hello in spanish", func(t *testing.T) {
		got := Hello("Marcel", "Spanish")
		want := "Hola, Marcel"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hello in french", func(t *testing.T) {
		got := Hello("Marcel", "French")
		want := "Bonjour, Marcel"
		assertCorrectMessage(t, got, want)
	})
}
