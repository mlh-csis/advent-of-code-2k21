package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input, _ := ReadInput("input_test")
	octopi := ParseInput(input)
	got, _ := Simulate(octopi, 100)
	want := 1656
	if got != want {
		t.Errorf("got %d, wanted %d for ", got, want)
	}
}

func TestPart2(t *testing.T) {
	input, _ := ReadInput("input_test")
	octopi := ParseInput(input)
	_, got := Simulate(octopi, 500)
	want := 195
	if got != want {
		t.Errorf("got %d, wanted %d for ", got, want)
	}
}
