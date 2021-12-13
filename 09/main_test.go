package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := GetInput("input_test")
	got := Part1(input)
	want := 15
	if got != want {
		t.Errorf("got %d, wanted %d for ", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := GetInput("input_test")
	got := Part2(input)
	want := 1134
	if got != want {
		t.Errorf("got %d, wanted %d for ", got, want)
	}
}
