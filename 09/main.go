package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := GetInput("input_test")
	//part1 := Part1(input)
	//fmt.Printf("Part 1: %d\n", part1)
	part2 := Part2(input)
	fmt.Printf("Part 2: %d\n", part2)
}

func ReadInput(path string) ([]string, error) {
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

func GetInput(path string) [][]int {
	input, _ := ReadInput(path)
	var output [][]int
	for _, line := range input {
		var numbers []int
		for _, number := range strings.Split(line, "") {
			n, _ := strconv.Atoi(number)
			numbers = append(numbers, n)
		}
		output = append(output, numbers)
	}
	return output
}

func Part1(input [][]int) int {
	var lowPoints []int
	for y, line := range input {
		for x, point := range line {
			fmt.Printf("iterating %d, %d\n", y, x)
			if IsLowPoint(y, x, input) {
				lowPoints = append(lowPoints, point)
			}
		}
	}

	sum := 0
	for _, p := range lowPoints {
		sum += 1 + p
	}

	return sum
}

func Part2(input [][]int) int {
	var basins [][]int
	for y, line := range input {
		for x := range line {
			CheckLowPointRecursive(y, x, input, &basins, "")
		}
	}
	return 0
}

func ok(y int, x int, input [][]int) bool {
	ysize := len(input) - 1
	xsize := len(input[0]) - 1
	if x <= -1 || y <= -1 || x > xsize || y > ysize {
		return false
	}
	return true
}

func IsLowPoint(y int, x int, input [][]int) bool {
	//fmt.Printf("IsLowPoint %d, %d\n", y, x)

	up, down, left, right := 9, 9, 9, 9

	if ok(y-1, x, input) {
		//fmt.Printf("Ignore up %d, %d, %s\n", y-1, x, ignore)
		up = input[y-1][x]
	}
	if ok(y+1, x, input) {
		//	fmt.Printf("Ignore down %d, %d\n", y+1, x)
		down = input[y+1][x]
	}
	if ok(y, x-1, input) {
		//	fmt.Printf("Ignore left %d, %d\n", y, x-1)
		left = input[y][x-1]
	}
	if ok(y, x+1, input) {
		//fmt.Printf("Ignore right %d, %d\n", y, x+1)
		right = input[y][x+1]
	}
	//fmt.Printf("IsLowPoint Done %d, %d\n", y, x)
	point := input[y][x]
	return point < up && point < down && point < left && point < right
}

func CheckLowPointRecursive(y int, x int, input [][]int, basins *[][]int, ignore string) {
	//fmt.Printf("IsLowPoint %d, %d\n", y, x)

	if !ok(y, x, input) {
		return
	}
	up, down, left, right := 9, 9, 9, 9

	if ignore != "up" && ok(y-1, x, input) {
		//fmt.Printf("Ignore up %d, %d, %s\n", y-1, x, ignore)
		up = input[y-1][x]
	}
	if ignore != "down" && ok(y+1, x, input) {
		//	fmt.Printf("Ignore down %d, %d\n", y+1, x)
		down = input[y+1][x]
	}
	if ignore != "left" && ok(y, x-1, input) {
		//	fmt.Printf("Ignore left %d, %d\n", y, x-1)
		left = input[y][x-1]
	}
	if ignore != "right" && ok(y, x+1, input) {
		//fmt.Printf("Ignore right %d, %d\n", y, x+1)
		right = input[y][x+1]
	}
	point := input[y][x]
	if point < up && point < down && point < left && point < right {
		fmt.Printf("IsLowPoint %d, %d\n", y, x)
		*basins = append(*basins, []int{y, x})
		CheckLowPointRecursive(y-1, x, input, basins, "down")
		CheckLowPointRecursive(y+1, x, input, basins, "up")
		CheckLowPointRecursive(y, x-1, input, basins, "right")
		CheckLowPointRecursive(y, x+1, input, basins, "left")
	}
}
