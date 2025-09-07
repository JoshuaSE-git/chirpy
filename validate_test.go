package main

import (
	"testing"
)

func TestCensorChirp(t *testing.T) {
	t.Run("text with censored word", func(*testing.T) {
		censorMap := map[string]struct{}{
			"bad": {},
		}
		text := "bad word"

		got := censorChirp(text, censorMap)
		want := "**** word"

		assertString(t, got, want)
	})
	t.Run("censored word with different case", func(*testing.T) {
		censorMap := map[string]struct{}{
			"bad": {},
		}
		text := "Bad word"

		got := censorChirp(text, censorMap)
		want := "**** word"

		assertString(t, got, want)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
