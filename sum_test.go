package main

import "testing"

func TestSum(t *testing.T) {
	got := sum(1, 2)
	want := 3

	if got != want {
		t.Errorf("Sum(1, 2) = %d; want %d", got, want)
	}
}

func TestSub(t *testing.T) {
	got := sub(1, 2)
	want := -1

	if got != want {
		t.Errorf("Sub(1, 2) = %d; want %d", got, want)
	}
}

func TestTimes(t *testing.T) {
	got := times(1, 2)
	want := 2

	if got != want {
		t.Errorf("Times(1, 2) = %d; want %d", got, want)
	}
}

func TestSumX(t *testing.T) {
	got := sumX(1, 2)
	want := 4

	if got != want {
		t.Errorf("SumX(1, 2) = %d; want %d", got, want)
	}
}
