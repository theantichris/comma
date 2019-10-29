package main

import "testing"

func TestComma(t *testing.T) {
	t.Run("returns digit number", func(t *testing.T) {
		got := comma("123")
		want := "123"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("returns number with 4 digit number", func(t *testing.T) {
		got := comma("1234")
		want := "1,234"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("returns number with 5 digit number", func(t *testing.T) {
		got := comma("12345")
		want := "12,345"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("returns number with 6 digit number", func(t *testing.T) {
		got := comma("123456")
		want := "123,456"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
