package main

import (
	"fmt"
	"strconv"
)

func main() {
	input, _ := ReadInput("input")
	octopi := ParseInput(input)

	part1a, part1b  := Simulate(octopi, 100)
	fmt.Printf("Part 1: %d, %d\n", part1a, part1b)
	part2a, part2b := Simulate(octopi, 1000)
	fmt.Printf("Part 2: %d, %d\n", part2a, part2b)
}

func ParseInput(lines []string) [][]int {
	var output [][]int
	for _, line := range lines {
		var numbers []int
		for _, c := range line {
			n, _ := strconv.Atoi(string(c))
			numbers = append(numbers, n)
		}
		output = append(output, numbers)
	}
	return output
}

func Simulate(input [][]int, steps int) (int, int) {
	flashes := 0

	var flashed func(y int, x int)
	flashed = func (y int, x int) {
		flashes++
		input[y][x] = -1
		for dy:=-1; dy <=1; dy++ {
			for dx:=-1; dx <=1; dx++ {
				new_y, new_x := y + dy, x + dx
				if new_y >= 0 && new_y < len(input) && new_x >= 0 && new_x < len(input[0]) && input[new_y][new_x] != -1 {
					input[new_y][new_x]++
					if input[new_y][new_x] >= 10{
						flashed(new_y, new_x)
					}
				}
			}
		}
	}

	i:= 0
	for {
		i++
		for y, row := range input {
			for x := range row {
				input[y][x]++
			}
		}

		for y, row := range input {
			for x := range row {
				if input[y][x] == 10 {
					flashed(y, x)
				}
			}
		}

		all_flashed := true

		for y, row := range input {
			for x := range row {
				if input[y][x] == -1 {
					input[y][x] = 0			
				} else {
					all_flashed = false
				}
			}
		}
		if i == steps {
			return flashes, 0
		}
		if all_flashed {
			return flashes, i
		}
	}
}
