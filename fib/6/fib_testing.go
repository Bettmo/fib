package fib

import "testing"

func TestFib(t *testing.T) {
	// Her går testen
	input := 1
	got := Fib(input)
	want := 1
	if got != want {
		t.Errorf("Fib(%d) = %d; want %d", input, got, want)
	}

}

