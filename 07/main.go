package main

import (
	"fmt"
	"time"
)

func main() {
	input, _ := readInput("input")
	start := time.Now()
	fuel := GetMinimumFuelPart1(input)
	elapsed := time.Since(start)
	fmt.Printf("Part 1: %d - took %s\n", fuel, elapsed)

	start = time.Now()
	fuel2 := GetMinimumFuelPart2(input)
	elapsed = time.Since(start)
	fmt.Printf("Part 12 %d - took %s\n", fuel2, elapsed)
}
