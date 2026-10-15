package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("greet", func(t *testing.T) {
		buffer := bytes.Buffer{}
		_, err := Greet(&buffer, "Chris")
		if err != nil {
			t.Fatal(err)
		}

		got := buffer.String()
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
