package main

import "testing"

func TestHello(t *testing.T) {
	assertEquals := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertEquals(t, got, want)
	})

	t.Run("empty string defaults saying to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertEquals(t, got, want)
	})

	t.Run("saying hello in Korean", func(t *testing.T) {
		got := Hello("John", "korean")
		want := "Annyeonghaseyo, John"
		assertEquals(t, got, want)
	})

	t.Run("saying hello in Japanese", func(t *testing.T) {
		got := Hello("Marin", "japanese")
		want := "Konnichiwa, Marin"
		assertEquals(t, got, want)
	})
}
