package fib

import "testing"

func TestFib(t *testing.T) {
	// Her går testen
	input := 0
	got := fib(input)
	want := 0
	if got != want {
		t.Errorf("Fib(%d) = %d; want %d", input, got, want)
	}

}
