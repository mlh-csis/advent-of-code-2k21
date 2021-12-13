package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readInput(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	str := strings.Split(line, ",")
	var input []int
	for _, f := range str {
		f_n, err := strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
		input = append(input, f_n)
	}
	return input, scanner.Err()
}

func GetMinimumFuelPart1(input []int) int {
	min, max := MinMax(input)
	fuelSpentAtposition := make(map[int]int, max-min)
	for _, position := range input {
		for i := min; i <= max; i++ {
			fuelSpentAtposition[i] += Abs(position - i)
		}
	}

	min_fuel := 4294967295

	for fuel := range fuelSpentAtposition {
		if fuel < min {
			min_fuel = fuel
		}
	}
	return min_fuel
}

func GetMinimumFuelPart2(input []int) int {
	min, max := MinMax(input)
	stepCosts := make(map[int]int, max-min)
	stepCosts[0] = 0
	for i := 1; i <= max-min; i++ {
		stepCosts[i] = stepCosts[i-1] + i
	}

	fuelSpentAtposition := make(map[int]int, max-min)
	for _, position := range input {
		for i := min; i <= max; i++ {
			fuelSpentAtposition[i] += stepCosts[Abs(position-i)]
		}
	}

	min_fuel := 4294967295

	for _, fuel := range fuelSpentAtposition {
		if fuel < min {
			min = fuel
		}
	}
	return min_fuel
}


