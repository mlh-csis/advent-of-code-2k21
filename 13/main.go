package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input, _ := ReadInput("input")
	Part1(input)
	Part2(input)
}

func Part1(input []string) {
	m := make(map[[2]int]struct{}, 10)

	for _, line := range input {
		if line == "" {
			// skip empty line between points and directions
			continue
		}

		var c, r int
		_, err := fmt.Sscanf(line, "%d,%d", &c, &r)
		if err == nil {
			// set point
			m[[2]int{r, c}] = struct{}{}
		} else {
			// fold
			fields := strings.Fields(line)
			vals := strings.Split(fields[2], "=")
			direction := vals[0]
			from_str := vals[1]
			from, _ := strconv.Atoi(from_str)

			if direction == "y" {
				for k := range m {
					row, col := k[0], k[1]
					if row < from {
						continue
					} 
					new_row := from - (row - from)
					m[[2]int{new_row, col}] = struct{}{}
					delete(m, [2]int{row, col})
				}
			} else {
				for k := range m {
					row, col := k[0], k[1]
					if col < from {
						continue
					} 
					new_col := from - (col - from)
					m[[2]int{row, new_col}] = struct{}{}
					delete(m, [2]int{row, new_col})
				}
			}	
			break
		}
	}
	fmt.Printf("Part 1: %d\n", len(m))
}

func Part2(input []string) {
	m := make(map[[2]int]struct{}, 10)

	for _, line := range input {
		if line == "" {
			// skip empty line between points and directions
			continue
		}

		var c, r int
		_, err := fmt.Sscanf(line, "%d,%d", &c, &r)
		if err == nil {
			// set point
			m[[2]int{c, r}] = struct{}{}
		} else {
			// fold
			fields := strings.Fields(line)
			vals := strings.Split(fields[2], "=")
			direction := vals[0]
			from_str := vals[1]
			from, _ := strconv.Atoi(from_str)

			if direction == "y" {
				for k := range m {
					col, row := k[0], k[1]
					if row < from {
						continue
					} 
					new_row := from - (row - from)
					m[[2]int{col, new_row}] = struct{}{}
					delete(m, [2]int{col, row})
				}
			} else {
				for k := range m {
					col, row := k[0], k[1]
					if col < from {
						continue
					} 
					new_col := from - (col - from)
					m[[2]int{new_col, row}] = struct{}{}
					delete(m, [2]int{col, row})
				}
			}	
		}
	}
	fmt.Println("Part 2:")

	max_col := 0
	for k := range m {
		if k[0] > max_col {
			max_col = k[0]
		}
	}

	max_row := 0
	for k := range m {
		if k[1] > max_row {
			max_row = k[1]
		}
	}
	//fmt.Println(max_row, max_col)
	for r:= 0; r <= max_row; r++ {
		for c:= 0; c <= max_col; c++ {
			//fmt.Println(i, j)
			if _, ok := m[[2]int{c, r}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}