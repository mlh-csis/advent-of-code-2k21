package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := GetInput("input")
	Solve(input, true)
	Solve(input, false)
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

func Solve(input [][]int, part_a bool) {
	var basins []int
	for y, line := range input {
		for x := range line {
			basin_size := GetBasinSize(y, x, input)
			if basin_size > 0 {
				if !part_a {
					basins = append(basins, basin_size)
				} else {
					basins = append(basins, 1)
				}
			}
		}
	}
	if !part_a {
		sort.Ints(basins)
		fmt.Printf("Part 2: %d\n", Product(basins[len(basins)-3:]))
	} else {
		fmt.Printf("Part 1: %d\n", sum(basins))
	}
}

type stack [][2]int

func (s *stack) Push(v [2]int) {
    *s = append(*s, v)
}

func (s *stack) Pop() ([2]int, error) {
	 if len(*s)-1 < 0 {
		return [2]int{0,0}, fmt.Errorf("Pop from an empty stack, len: %d", len(*s))
	 }
    res:=(*s)[len(*s)-1]
    *s=(*s)[:len(*s)-1]
    return res, nil
}

func GetBasinSize(y int, x int, input [][]int) int {
	dir := [][2]int {{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	visited := make(map[[2]int]bool)
	basins := 0

	var s stack
	s.Push([2]int{y, x})

	for {
		point, err := s.Pop()
		// no more points to visit, return the number of visited
		if err != nil {
			return basins
		}

		y, x := point[0], point[1]
		// if the point we are visiting in the grid does not exist
		if y < 0 || y >= len(input) || x < 0 || x >= len(input[0])  || input[y][x] == 9 {
			visited[[2]int{y,x}] = true
			continue
		}

		// skip if we have already visited this point
		if _, ok := visited[[2]int{y,x}]; ok {
			continue
		}

		smallest:= true
		for _, d := range dir {
			dy, dx := d[0], d[1]
			_, seen := visited[[2]int{y+dy,x+dx}]
			if y+dy >= 0 && y+dy < len(input) && x+dx >= 0 && x+dx < len(input[0]) && !seen {
				if input[y][x] > input[y+dy][x+dx] {
					smallest = false
					break
				}
			}
		}
		if !smallest {
			continue
		}
		visited[[2]int{y,x}] = true
		// if we reach here, we are smaller than our neighbors, so increase number
		basins++
		for _, d := range dir {
			dy, dx := d[0], d[1]
			s.Push([2]int{y+dy, x+dx})
		}
	}
}