package main

import (
	"fmt"
	"strconv"
)

type location struct {
	depth    int
	position int
	aim      int
}

func (l *location) forward(amount int) {
	l.position += amount
}

func (l *location) up(amount int) {
	l.depth -= amount
}

func (l *location) down(amount int) {
	l.depth += amount
}

func (l *location) forward2(amount int) {
	l.position += amount
	l.depth += l.aim * amount
}

func (l *location) up2(amount int) {
	l.aim -= amount
}

func (l *location) down2(amount int) {
	l.aim += amount
}


func main() {
	lines, _ := readInput("input")
	part1 := Part1(lines)
	fmt.Printf("Part 1: %d\n", part1)
	part2 := Part2(lines)
	fmt.Printf("Part 2: %d\n", part2)
}

func Part2(lines [][]string) int {
	var l location
	ops := map[string]func(amount int){
		"forward": l.forward2,
		"down":    l.down2,
		"up":      l.up2,
	}
	for _, directions := range lines {
		var amount, _ = strconv.Atoi(directions[1])
		ops[directions[0]](amount)
	}
	return l.depth * l.position
}

func Part1(lines [][]string) int {
	var l location
	ops := map[string]func(amount int){
		"forward": l.forward,
		"down":    l.down,
		"up":      l.up,
	}
	for _, directions := range lines {
		var amount, _ = strconv.Atoi(directions[1])
		ops[directions[0]](amount)
	}
	return l.depth * l.position
}