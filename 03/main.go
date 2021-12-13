package main

import (
	"fmt"
	"strconv"
)

func main() {
	lines, _ := readInput("input")
	var matrix [][]int

	for _, measurement := range lines {
		var line []int
		for i := 0; i <= 11; i++ {
			bit, _ := strconv.Atoi(string(measurement[i]))
			line = append(line, bit)
		}
		matrix = append(matrix, line)
	}

	map_oxygen := make(map[int][][]int)
	oxygen := matrix
	
	for i := 0; i <= 11; i++ {
		for _, line := range oxygen {
			map_oxygen[line[i]] = append(map_oxygen[line[i]], line)
		}
		if len(map_oxygen[1]) >= len(map_oxygen[0]) {
			oxygen = map_oxygen[1]
		} else {
			oxygen = map_oxygen[0]
		}
		map_oxygen = make(map[int][][]int)
	}

	map_scrubber := make(map[int][][]int)
	var scrubber = matrix
	for i := 0; i <= 11; i++ {
		for _, line := range scrubber {
			map_scrubber[line[i]] = append(map_scrubber[line[i]], line)
		}
		if len(map_scrubber[1]) >= len(map_scrubber[0]) {
			scrubber = map_scrubber[0]
		} else {
			scrubber = map_scrubber[1]
		}
		if len(scrubber) == 1 {
			break
		}
		map_scrubber = make(map[int][][]int)
	}
	var stroxygen string
	for _, n := range oxygen[0] {
		stroxygen += strconv.Itoa(n)
	}

	var strscubber string
	for _, n := range scrubber[0] {
		strscubber += strconv.Itoa(n)
	}

	ox, _ := strconv.ParseInt(stroxygen, 2, 64)
	sc, _ := strconv.ParseInt(strscubber, 2, 64)
	fmt.Println(ox * sc)
}
