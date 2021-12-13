package main

import (
	"testing"
)

func TestPart1TestInput(t *testing.T) {
	fishes := []int{3, 4, 3, 1, 2}

	parameters := [][]int{
		{1, 5},
		{2, 6},
		{3, 7},
		{4, 9},
		{5, 10},
		{6, 10},
		{7, 10},
		{8, 10},
		{9, 11},
		{10, 12},
		{11, 15},
		{12, 17},
		{13, 19},
		{14, 20},
		{15, 20},
		{16, 21},
		{17, 22},
		{18, 26},
	}

	for _, i := range parameters {
		got := calculatePopulationAt(fishes, i[0])
		want := i[1]
		if got != int64(want) {
			t.Errorf("got %d, wanted %d for days %d", got, want, i[0])
		}
	}
}
