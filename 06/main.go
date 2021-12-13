package main

import (
	"fmt"
	"os"
)

func main() {
	fishes, err := readInput("input")

	if err != nil {
		fmt.Printf("Read error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1: %d\n", calculatePopulationAt(fishes, 80))
	fmt.Printf("Part 2: %d\n", calculatePopulationAt(fishes, 256))
}
