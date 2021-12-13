package main

import (
	"fmt"
	"strconv"
)

func main() {
	lines := readIntInput("input")

	higher := getNumberOfIncreases(lines)
	fmt.Printf("Part 1: %d\n", higher)

	higherSums := getNumberOfIncreases2(lines)
	fmt.Printf("Part 2: %d\n", higherSums)
}

func readIntInput(input string) []int {
	lines, _ := ReadInput(input)
	var output []int
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		output = append(output, n)
	}
	return output
}

func getNumberOfIncreases(iter []int) int {
	higher := 0
	for i := 1; i < len(iter); i++ {
		if iter[i]-iter[i-1] > 0 {
			higher++
		}
	}
	return higher
}
func getNumberOfIncreases2(iter []int) int {
	higher := 0
	for i := 4; i < len(iter); i++ {
		a := iter[i] + iter[i-1] + iter[i-2]
		b := iter[i-1] + iter[i-2] + iter[i-3]
		
		if a > b {
			higher++
		}
	}
	return higher
}
