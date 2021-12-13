package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := readInput("input_test")
	if err != nil {
		t.Error(err)
	}

	displays := ParseInput(input)
	got := Part1(displays)
	want := 26

	if got != want {
		t.Errorf("got %d, wanted %d for ", got, want)
	}
}

func TestPart2(t *testing.T) {
	display := DisplayFromString("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	got := Part2([]Display{display})
	want := 26

	if got != want {
		t.Errorf("got %d, wanted %d for ", got, want)
	}
}
