package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input, _ := ReadInput("input_test")
	got, _ := BothParts(input)
	want := 26397
	if got != want {
		t.Errorf("got %d, wanted %d for ", got, want)
	}
}

func TestPart1Line(t *testing.T) {
	got, _ := checkLine("{([(<{}[<>[]}>{[]{[(<()>")
	want := 1197
	if got != want {
		t.Errorf("got %d, wanted %d for ", got, want)
	}
}

func TestPart1Line2(t *testing.T) {
	got, _ := checkLine("[[<[([]))<([[{}[[()]]]")
	want := 3
	if got != want {
		t.Errorf("got %d, wanted %d for ", got, want)
	}
}

func TestPart1Line3(t *testing.T) {
	got, _ := checkLine("[{[{({}]{}}([{[{{{}}([]")
	want := 57
	if got != want {
		t.Errorf("got %d, wanted %d for ", got, want)
	}
}

func TestPart1Line4(t *testing.T) {
	got, _ := checkLine("[<(<(<(<{}))><([]([]()")
	want := 3
	if got != want {
		t.Errorf("got %d, wanted %d for ", got, want)
	}
}

func TestPart2Line1(t *testing.T) {
	_, got := checkLine("<{([{{}}[<[[[<>{}]]]>[]]")
	want := 294
	if got != want {
		t.Errorf("got %d, wanted %d for ", got, want)
	}
}

func TestPart2Line2(t *testing.T) {
	_, got := checkLine("[({(<(())[]>[[{[]{<()<>>")
	want := 288957
	if got != want {
		t.Errorf("got %d, wanted %d for ", got, want)
	}
}
