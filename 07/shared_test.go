package main

import (
	"testing"
)

func TestPart1TestInput(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	fuel := GetMinimumFuelPart1(input)
	want := 37
	if fuel != want {
		t.Errorf("got %d, wanted %d for ", fuel, want)
	}
}
