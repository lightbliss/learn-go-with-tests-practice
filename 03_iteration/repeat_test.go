package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q, got %q", expected, repeated)
	}
}

func TestRepeat2(t *testing.T) {
	repeated := Repeat2("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q, got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}

func BenchmarkRepeat2(b *testing.B) {
	for b.Loop() {
		Repeat2("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat2("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}
