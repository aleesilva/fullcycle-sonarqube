package main

import "testing"

func TestSum(t *testing.T) {
	got := sum(1, 2)
	want := 3

	if got != want {
		t.Errorf("Sum(1, 2) = %d; want %d", got, want)
	}
}
