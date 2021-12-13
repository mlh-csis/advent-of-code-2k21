package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input, _ := ReadInput("input_test")
	parsed := ParseInput(input)
	got := Navigate(parsed, can_visit)
	want := 10
	if got != want {
		t.Errorf("got %d, wanted %d for ", got, want)
	}
}

func TestPart2(t *testing.T) {
	input, _ := ReadInput("input_test")
	parsed := ParseInput(input)
	got := Navigate(parsed, can_visit2)
	want := 36
	if got != want {
		t.Errorf("got %d, wanted %d for ", got, want)
	}
}
