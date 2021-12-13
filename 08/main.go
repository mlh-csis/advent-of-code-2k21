package main

import "fmt"

func main() {
	input, _ := readInput("input")
	displays := ParseInput(input)
	fmt.Printf("Part 1: %d\n", Part1(displays))
	fmt.Printf("Part 2: %d\n", Part2(displays))
}
