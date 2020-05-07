package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("should greet the user", func(t *testing.T) {
		got := Hello("anna", "English")
		want := "Hello, anna"

		assertCorrectMessage(t, got, want)
	})

	t.Run("should say 'Hello, World' if name is not defined", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Anton", "Spanish")
		want := "Hola, Anton"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("mary", "French")
		want := "Bonjour, mary"

		assertCorrectMessage(t, got, want)
	})
}

// func BenchmarkHello(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		fmt.Sprintf("hello")
// 	}
// }
