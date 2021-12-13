package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func readInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

type Display struct {
	Input  []string
	Output []string
}

func DisplayFromString(line string) Display {
	parts := strings.Split(line, " | ")
	input, output := parts[0], parts[1]
	return Display{
		Input:  strings.Fields(input),
		Output: strings.Fields(output),
	}
}

func ParseInput(input []string) []Display {
	var displays []Display
	for _, s := range input {
		displays = append(displays, DisplayFromString(s))
	}
	return displays
}

func Part1(displays []Display) int {
	count := 0
	for _, display := range displays {
		output := display.Output
		for _, o := range output {
			length := len(o)
			if (length >= 2 && length <= 4) || length == 7 {
				count++
			}
		}
	}
	return count
}

func Part2(displays []Display) int {
	count := 0
	for _, display := range displays {
		input := append(display.Input)
		d := make([]string, 10)

		for _, i:= range input {
			str := SortString(i)
			switch len(i) {
			case 2:
				d[1] = str
			case 3:
				d[7] = str
			case 4:
				d[4] = str
			case 7:
				d[8] = str
			}
		}

		for _, i := range input {
			str := SortString(i)
			switch len(i) {
			case 5:
				if strings.Contains(str, d[1]) {
					d[3] = str
				} else if 

				d[2] = append(d[2], str)
				d[3] = append(d[3], str)
				d[5] = append(d[5], str)
			case 6:
				d[0] = append(d[0], str)
				d[6] = append(d[6], str)
				d[9] = append(d[9], str)
			case 7:
				d[8] = append(d[0], str)
			}
		}

		vals := make(map[string]int)

		for i, val := range d {
			if union(val) == val (

			)
			switch i {
			case 1, 7, 4, 8:
				if len(val) > 0 {
					vals[val[0]] = i
				}
			}
			if len(val) == 1 {
				continue
			}

			if i == 5 {

			}

			if i == 5 {

			}

		}

		fmt.Printf("%v", d)
	}
	return count
}

